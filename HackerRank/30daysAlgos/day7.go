package main

import (
	"bufio"
	"fmt"
	"io"
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

func formatInt32(n int32) string {
	return strconv.FormatInt(int64(n), 10)
}
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := int(n) - 1; i >= 0; i-- {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}
	valuesText := []string{}
	for _, v := range arr {
		valuesText = append(valuesText, formatInt32(v))
	}

	// Join our string slice.
	result := strings.Join(valuesText, " ")
	fmt.Println(result)
	// fmt.Println(strings.Join(arr, " "))

}
