package main

import (
	"bufio"
	"fmt"
	"strconv"
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

/**************** Day 3 ******************/
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	N := int32(NTemp)
	if N < 1 || N > 100 || N%2 != 0 {
		fmt.Println("Weird")
	} else {

		if N >= 2 && N <= 5 || N > 20 {
			fmt.Println("Not Weird")
		}
		if N >= 6 && N <= 20 {
			fmt.Println("Weird")
		}
	}

}