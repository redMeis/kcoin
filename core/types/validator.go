package types

import (
	"github.com/kowala-tech/kUSD/common"
)

// validator represents a consensus validator
type Validator struct {
	address common.Address
	power   uint64 // voting power
	accum   uint64 // @TODO (rgeraldes) - overflow big.Int?
}

func NewValidator(address common.Address, power uint64) *Validator {
	return &Validator{
		address: address,
		power:   power,
		accum:   0,
	}
}

func (val *Validator) Hash() common.Hash {
	return rlpHash([]interface{}{val.address, val.power})
}
func (val *Validator) Address() common.Address { return val.address }
func (val *Validator) Power() uint64           { return val.power }

type ValidatorSet struct {
	validators []*Validator
	proposer   *Validator

	// cache
}

func NewValidatorSet(validators []*Validator) *ValidatorSet {
	// @TODO (rgeraldes) - size needs to be > 0
	return &ValidatorSet{
		validators: validators,
	}
}

func (set *ValidatorSet) AtIndex(i int) *Validator {
	return set.validators[i]
}

func (set *ValidatorSet) Size() int {
	return len(set.validators)
}

func (set *ValidatorSet) Proposer() common.Address {
	// @TODO (rgeraldes) complete - return the first validator for now
	return set.validators[0].Address()
}