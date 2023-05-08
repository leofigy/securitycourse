package main

import (
	"fmt"
	"roles/model"
)

func main() {
	persistence := model.NewPersistence("localdb", false, "")
	handler, err := persistence.GetDB()

	if err != nil {
		panic(err)
	}

	fmt.Println(handler)
}
