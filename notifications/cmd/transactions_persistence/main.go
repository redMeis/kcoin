package main

import (
	"os"
		"github.com/sirupsen/logrus"

	"github.com/go-redis/redis"
	"github.com/yourheropaul/inj"

		"github.com/kowala-tech/kcoin/notifications/core"
	"github.com/kowala-tech/kcoin/notifications/environment"
		"github.com/kowala-tech/kcoin/notifications/persistence"
	"github.com/kowala-tech/kcoin/notifications/pubsub"
)

func main() {
	envReader := environment.NewReaderOs()
	redisAddr := envReader.Read("REDIS_ADDR")
	nsqAddr := envReader.Read("NSQ_ADDR")
	logLevelRaw := envReader.Read("LOG_LEVEL")
	if logLevelRaw == "" {
		logLevelRaw = "info"
	}

	logLevel, err := logrus.ParseLevel(logLevelRaw)
	if err != nil {
		panic(err)
	}

	logger := logrus.New()
	logger.SetLevel(logLevel)
	logger.Out = os.Stdout

	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = redisClient.Ping().Result()
	if err != nil {
		panic(err)
	}

	worker := core.NewTransactionsPersistanceWorker(logrus.NewEntry(logger))

	g := inj.NewGraph()
	g.Provide(
		worker,
		persistence.NewRedisPersistence(redisClient),
		pubsub.NewNSQSubscriber("notifications", "db-persistance", nsqAddr, logrus.NewEntry(logger)),
	)

	if valid, errors := g.Assert(); !valid {
		panic(errors)
	}

	err = worker.Start()
	if err != nil {
		panic(err)
	}
}
