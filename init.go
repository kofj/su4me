package main

import (
	"log"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const defaultAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_"

var (
	sha2id = []byte("sha2id")
	id2url = []byte("id2url")

	err error
)

func init() {

	// Parse flags.
	pflag.StringP("add", "a", "0:2345", "server <addr>:<port>")
	pflag.IntP("min", "m", 0, "min length")
	pflag.StringP("salt", "s", "", "hash salt")
	pflag.StringP("wd", "w", "var", "work dir")
	pflag.String("db", "su4me", "database name")
	pflag.Parse()

	// Load config and set default values.
	viper.BindPFlags(pflag.CommandLine)
	viper.SetConfigName("cfg")
	viper.AddConfigPath(".")
	viper.AddConfigPath("etc")
	viper.AddConfigPath("/etc/su4me")
	viper.SetDefault("addr", "0:2345")
	viper.SetDefault("min", 3)
	viper.SetDefault("salt", "aaa")
	viper.SetDefault("wd", "var")
	viper.SetDefault("alphabet", defaultAlphabet)
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {
		log.Fatalln(err)
	}

}
