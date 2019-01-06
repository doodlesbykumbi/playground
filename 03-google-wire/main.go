package main

import (
	"context"
	"fmt"
)

func main() {
	c := context.Background()
	db, err := initializeDB(c)
	if err != nil {
		panic(err)
	}

	var s struct {
		I int `db:"num"`
	}
	err = db.Get(&s, "select 1 as num")
	if err != nil {
		panic(err)
	}
	fmt.Printf("got %d\n", s.I)
}