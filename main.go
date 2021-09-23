package main

import (
	"fmt"
	"go-interfaces/consumer"
	"go-interfaces/good"
	"go-interfaces/logger"

	"github.com/alicebob/miniredis"
)

var log = logger.New("main")

func main() {
	animals := []consumer.Animal{
		consumer.GetAnimal("Dog"),
		consumer.GetAnimal("Cat"),
		consumer.GetAnimal("Snake"),
		consumer.GetAnimal("Bird"),
	}

	// BoltDB
	log.Infof("Initialising bolt database ...")
	db, err := good.GetDatabase("bolt.db", false)
	if err != nil {
		log.Error(err)
	}
	log.Infow("Bolt database successfully opened", "Database", db)
	consumer.ProcessAnimalsBolt(db, animals)

	// Redis
	log.Infof("Initialising redis database ...")
	mr, err := miniredis.Run()
	if err != nil {
		log.Error(err)
	}
	rdb, err := good.NewClient(fmt.Sprintf("redis://%s/0", mr.Addr()))
	if err != nil {
		log.Error(err)
	}
	log.Infow("Redis database successfully opened", "Database", rdb)
	consumer.ProcessAnimalsRedis(rdb, animals)
}
