package main

import (
	"log"

	"github.com/boltdb/bolt"
	"github.com/spf13/viper"
)

var (
	db *bolt.DB
)

func main() {
	// Open database.
	db, err = bolt.Open(viper.GetString("wd")+"/"+viper.GetString("db")+".db", 0600, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Create buckets.
	createBuckets(sha2id)
	createBuckets(id2url)

	r.GET("ping", pong)

	log.Fatalln(r.Run(viper.GetString("addr")))
}
