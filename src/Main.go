package main

import (
	"github.com/fajardm/inventories/src/boot"
)

func main() {
	router := boot.SetupMain()

	router.Run(":3000")
}
