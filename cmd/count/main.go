package main

import (
	"context"
	"log"

	"github.com/goreleaser/func/count"
)

func main() {
	log.SetFlags(0)
	c, err := count.Count(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(c)
}
