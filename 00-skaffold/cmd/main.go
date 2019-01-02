package main

import (
	"example/pkg"
	"time"
)

func main() {
	person := pkg.NewPerson("Kumbi")
	for {
		person.Speak()

		time.Sleep(time.Second * 1)
	}
}
