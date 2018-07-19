package main

import (
	"log"

	"github.com/harveylowndes/wercker-test/pkg/name"
)

func main() {
	log.Printf("Starting...")
	t := &name.Test{}
	name.CreateName(t)
	log.Printf("Finished. Name is now '%s'", t.GetName())
}
