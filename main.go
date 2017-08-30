package main

import (
	"log"

	"github.com/spf13/viper"
)

func main() {
	r.GET("/ping", pong)

	log.Fatalln(r.Run(viper.GetString("addr")))
}
