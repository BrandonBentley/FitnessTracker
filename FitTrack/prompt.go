package main

import (
	"fmt"
	"bufio"
	"os"
	"sync"
	"strings"
)

func prompt(wg sync.WaitGroup) {
	fmt.Println("Press 'Enter' to Abort...")
	for true {
		test, _ := bufio.NewReader(os.Stdin).ReadBytes('\n')
		testString := string(test)
		if strings.Contains(testString, "exit") {
			os.Exit(0)
		} else if strings.Contains(testString, "cookies"){
			for _, c := range cookieJar {
				fmt.Println(c)
			}
			fmt.Println()
		}
	}
	wg.Done()
}
