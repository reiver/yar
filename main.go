package main

import (
	"fmt"

	"yar/srv/log"
)

func main() {
	log := logsrv.Begin()
	defer log.End()

	fmt.Println("yar ⚡")

	shout()
}
