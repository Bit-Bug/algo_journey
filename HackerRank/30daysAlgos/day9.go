package main

import (
	// "bufio"
	"fmt"
	// "io"
	// "os"
	// "strconv"
	// "strings"
)

// Complete the factorial function below.
func factorial(n int32) int32 {
	if n == 1 {
		return 1
	} else {
		return n * factorial(n-1)
	}

}

func main() {
	// reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	// stdout, _ := os.Create(os.Getenv("../day10.go"))
	// // checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 1024*1024)

	// nTemp, _ := strconv.ParseInt(readLine(reader), 10, 64)
	// // checkError(err)
	// n := int32(nTemp)

	// result := factorial(n)

	fmt.Println(factorial(4))

	// writer.Flush()
}

// func readLine(reader *bufio.Reader) string {
// 	str, _, err := reader.ReadLine()
// 	if err == io.EOF {
// 		return ""
// 	}

// 	return strings.TrimRight(string(str), "\r\n")
// }

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
