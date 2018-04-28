package main

import (
	"fmt"
	"time"
)

func main() {
	year, month, day := time.Now().Date()
	todaysDate := fmt.Sprintf("%d-%02d-%d", year, int(month), day)
	fmt.Println(todaysDate)
}
