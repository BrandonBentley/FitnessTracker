package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/satori/go.uuid"
)

func main() {
	for i := 0; i < 10; i++ {
		id, err := uuid.NewV4()
		if err != nil {
			fmt.Print("This happened: ")
			fmt.Println(err)
		}
		fmt.Println(id)
	}
}
