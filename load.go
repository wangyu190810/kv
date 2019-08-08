package main

import (
	"fmt"
	"kv/config"
	"kv/utils"
	"log"
)

var (
	logger   *log.Logger
	data     utils.Data
	filename = "kvdata.god"
)

func init() {
	data, _ = utils.Load(filename)
	config.Data = data
	fmt.Println(data)
}
