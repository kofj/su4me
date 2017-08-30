package main

import (
	"log"

	"github.com/boltdb/bolt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	sha2id = []byte("sha2id")
	id2url = []byte("id2url")

	db  *bolt.DB
	err error
)

func init() {

	// Parse flags.
	pflag.Int("port", 0, "Listen port")
	pflag.Int("min", 0, "min length")
	pflag.String("salt", "", "hash salt")
	pflag.String("wd", "var", "work dir")
	pflag.Parse()

	// Load config and set default values.
	viper.BindPFlags(pflag.CommandLine)
	viper.SetConfigName("cfg")
	viper.AddConfigPath(".")
	viper.AddConfigPath("etc")
	viper.AddConfigPath("/etc/su4me")
	viper.SetDefault("port", 2345)
	viper.SetDefault("min", 3)
	viper.SetDefault("salt", "aaa")
	viper.SetDefault("wd", "var")
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {
		log.Fatalln(err)
	}

	// Open database.
	db, err = bolt.Open(viper.GetString("wd")+"/ddz.db", 0600, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Create buckets.
	createBuckets(sha2id)
	createBuckets(id2url)
}
