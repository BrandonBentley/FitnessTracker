package main

import (
	"fmt"
	"bufio"
	"os"
	"sync"
)

func prompt(wg sync.WaitGroup) {
	fmt.Println("Press 'Enter' to Abort...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	os.Exit(0)
	wg.Done()
}
