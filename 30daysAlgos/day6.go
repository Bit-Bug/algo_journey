package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*********** Day  6**********/
func main() {
	var T int
	var s string
	var entries []string
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var even, noeven string

		fmt.Scan(&s)
		for j, v := range s {
			if j%2 == 0 {
				even = even + string(v)
			} else {
				noeven = noeven + string(v)
			}
		}
		entries = append(entries, even+" "+noeven)

	}
	for _, val := range entries {
		fmt.Println(val)
	}

}
