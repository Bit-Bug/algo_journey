package main

import (
	"fmt"
)

func main() {
	var T int
	var s int
	var entries []string
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		fmt.Scan(&s)
		entries = append(entries, strconv.Itoa(s))

	}
	fmt.Println(strings.Join(entries, " "))

}
