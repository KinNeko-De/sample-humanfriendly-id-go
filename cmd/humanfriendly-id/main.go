package main

import (
	"github.com/kinneko-de/sample-humanfriendly-id-go/internal/app/humanfriendly"
)

func main() {
	humanFriendlyId, err := humanfriendly.NewHumanFriendlyId(16)
	if err != nil {
		panic(err)
	}
	println("Id: " + humanFriendlyId.Id)
	println("Display as: " + humanFriendlyId.DisplayId)
}
