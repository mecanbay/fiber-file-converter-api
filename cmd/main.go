package main

import (
	"fiber-file-converter-api/pkg/config"
	"fmt"
)

func main() {
	// load config file
	cfg := config.Load()
	fmt.Println(cfg)
}
