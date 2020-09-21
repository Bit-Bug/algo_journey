package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	var s string
	var entries []string

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)
	m := make(map[string]string)
	for i := 0; i < int(n); i++ {
		arrTemp := strings.Split(readLine(reader), " ")
		m[arrTemp[0]] = arrTemp[1]

	}
	for i := 0; i < int(n); i++ {
		fmt.Scan(&s)
		entries = append(entries, s)

	}
	for _, val := range entries {
		if _, ok := m[val]; ok {
			fmt.Println(val + "=" + m[val])
		} else {
			fmt.Println("Not found")

		}
	}

}
