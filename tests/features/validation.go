package features

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/DATA-DOG/godog"
	"github.com/kowala-tech/kcoin/cluster"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/log"
)

type ValidationContext struct {
	globalCtx       *Context
	accountPassword string
	nodeRunning     bool
}

func NewValidationContext(parentCtx *Context) *ValidationContext {
	return &ValidationContext{
		globalCtx:       parentCtx,
		accountPassword: "test",
		nodeRunning:     false,
	}
}

func (ctx *ValidationContext) nodeID() cluster.NodeID {
	return cluster.NodeID("validator-under-test-" + ctx.globalCtx.nodeSuffix)
}

func (ctx *ValidationContext) IStopValidation() error {
	return godog.ErrPending
}

func (ctx *ValidationContext) IWaitForTheUnbondingPeriodToBeOver() error {
	return godog.ErrPending
}

func (ctx *ValidationContext) IStartTheValidator(kcoin int64) error {
	_, err := ctx.Do(setDeposit(kcoin))
	if err != nil {
		return err
	}

	_, err = ctx.Do(validatorStartCommand())
	if err != nil {
		return err
	}

	return nil
}

func (ctx *ValidationContext) IWaitForMyNodeToBeSynced() error {
	return common.WaitFor("timeout waiting for node sync", time.Second, time.Second*5, func() bool {
		return ctx.MyNodeIsAlreadySynchronised() == nil
	})
}

func (ctx *ValidationContext) IShouldBeAValidator() error {
	res, err := ctx.globalCtx.nodeRunner.Exec(ctx.nodeID(), isRunningCommand())
	if err != nil {
		log.Debug(res.StdOut)
		return err
	}
	if strings.TrimSpace(res.StdOut) != "true" {
		log.Debug(res.StdOut)
		return errors.New("validator is not running")
	}
	return nil
}

func (ctx *ValidationContext) IHaveMyNodeRunning(account string) error {
	if ctx.nodeRunning {
		return nil
	}
	ctx.nodeRunning = true

	spec := cluster.NewKcoinNodeBuilder().
		WithBootnode(ctx.globalCtx.bootnode).
		WithLogLevel(3).
		WithID(ctx.nodeID().String()).
		WithSyncMode("full").
		WithNetworkId(ctx.globalCtx.chainID.String()).
		WithGenesis(ctx.globalCtx.genesis).
		WithAccount(ctx.globalCtx.AccountsStorage, ctx.globalCtx.accounts[account]).
		NodeSpec()

	if err := ctx.globalCtx.nodeRunner.Run(spec, ctx.globalCtx.scenarioNumber); err != nil {
		return err
	}

	return nil
}

func (ctx *ValidationContext) IWithdrawMyNodeFromValidation() error {
	_, err := ctx.Do(stopValidatingCommand())
	if err != nil {
		return err
	}

	return nil
}

func (ctx *ValidationContext) ThereShouldBeTokensAvailableToMeAfterDays(expectedKcoins, days int) error {
	res, err := ctx.globalCtx.nodeRunner.Exec(ctx.nodeID(), getDepositsCommand())
	if err != nil {
		log.Debug(res.StdOut)
		return err
	}

	deposit, err := parseDepositResponse(res.StdOut)
	if err != nil {
		log.Debug(res.StdOut)
		return err
	}

	if expectedKcoins != *deposit.Value {
		return errors.New(fmt.Sprintf("kcoins don't match expected %d kcoins got %d", expectedKcoins, *deposit.Value))
	}

	daysExpected := time.Hour * 24 * time.Duration(days)
	expectedDate := time.Now().Add(daysExpected)
	if isSameDay(expectedDate, deposit.AvailableAt.Time()) {
		return errors.New(fmt.Sprintf("deposit available not within %d days, available at %s", daysExpected, deposit.AvailableAt.Time().String()))
	}

	return nil
}

func isSameDay(date1, date2 time.Time) bool {
	expectedYear, expectedMonth, expectedDay := date1.Date()
	availableYear, availableMonth, availableDay := date2.Date()
	return expectedYear != availableYear ||
		expectedMonth != availableMonth ||
		expectedDay != availableDay
}

func (ctx *ValidationContext) MyNodeShouldBeNotBeAValidator() error {
	res, err := ctx.globalCtx.nodeRunner.Exec(ctx.nodeID(), isRunningCommand())
	if err != nil {
		log.Debug(res.StdOut)
		return err
	}
	if strings.TrimSpace(res.StdOut) != "false" {
		log.Debug(res.StdOut)
		return errors.New("validator running")
	}
	return nil
}

func (ctx *ValidationContext) Reset() {
	ctx.nodeRunning = false
	ctx.globalCtx.nodeRunner.Stop(ctx.nodeID())
}

func (ctx *ValidationContext) MyNodeIsAlreadySynchronised() error {
	res, err := ctx.globalCtx.nodeRunner.Exec(ctx.nodeID(), isSyncedCommand())
	if err != nil {
		log.Debug(res.StdOut)
		return err
	}
	if strings.TrimSpace(res.StdOut) != "true" {
		log.Debug(res.StdOut)
		return errors.New("node is not synced")
	}
	return nil
}

// Do executes the command on the node and waits 1 block then
func (ctx *ValidationContext) Do(command []string) (*cluster.ExecResponse, error) {
	currentBlock, err := ctx.currentBlock()
	if err != nil {
		return nil, err
	}

	res, err := ctx.globalCtx.nodeRunner.Exec(ctx.nodeID(), command)
	if err != nil {
		log.Debug(res.StdOut)
		return nil, err
	}

	err = ctx.waitBlocksFrom(currentBlock,1)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (ctx *ValidationContext) waitBlocksFrom(block, n int) error {
	t := time.NewTicker(200*time.Millisecond)
	timeout := time.NewTimer(20*time.Second)
	defer t.Stop()

	var (
		err error
		newBlock int
	)

waitLoop:
	for {
		select {
		case <-timeout.C:
			return fmt.Errorf("timeout. started with block %d, finished with %d", block, newBlock)
		case <-t.C:
			newBlock, err = ctx.currentBlock()
			if err != nil {
				return err
			}

			blocks := newBlock - block

			if blocks >= n {
				break waitLoop
			}
		}
	}


	return nil
}

func (ctx *ValidationContext) waitBlocks(n int) error {
	currentBlock, err := ctx.currentBlock()
	if err != nil {
		return err
	}

	return ctx.waitBlocksFrom(currentBlock, n)
}

func (ctx *ValidationContext) currentBlock() (int, error) {
	res, err := ctx.globalCtx.nodeRunner.Exec(ctx.nodeID(), blockNumberCommand())
	if err != nil {
		log.Debug(res.StdOut)
		return 0, err
	}

	return strconv.Atoi(strings.TrimSpace(res.StdOut))
}

func blockNumberCommand() []string {
	return cluster.KcoinExecCommand("eth.blockNumber")
}

func isSyncedCommand() []string {
	return cluster.KcoinExecCommand("eth.blockNumber > 0 && eth.syncing == false")
}

func setDeposit(kcoin int64) []string {
	return cluster.KcoinExecCommand(fmt.Sprintf("validator.setDeposit(%d)", toWei(kcoin)))
}

func validatorStartCommand() []string {
	return cluster.KcoinExecCommand("validator.start()")
}

func stopValidatingCommand() []string {
	return cluster.KcoinExecCommand("validator.stop()")
}

func isRunningCommand() []string {
	return cluster.KcoinExecCommand("validator.isRunning()")
}

func getDepositsCommand() []string {
	return cluster.KcoinExecCommand("validator.getDeposits()")
}
