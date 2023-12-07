package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	total := 0

	namesStr := "one,two,three,four,five,six,seven,eight,nine"
	nameArr := strings.Split(namesStr, ",")

	dat, err := os.ReadFile("example.dat")
	check(err)
	dataStr := string(dat)

	dataArr := strings.Split(dataStr, "\n")
	for i := 0; i < len(dataArr); i++ {
		fmt.Print(dataArr[i], "\n")

		for k := 0; k < len(nameArr); k++ {
			dataStr = strings.Replace(dataArr[i], nameArr[k], fmt.Sprint(k+1), -1)

			dataArr = strings.Index(dataArr[i], nameArr[i])
		}

		for j := 0; j < len(dataArr[i]); j++ {
			fmt.Print(string(dataArr[i][j]), "\n")
			if dataArr[i][j] < 48 || dataArr[i][j] > 57 {
				continue
			}
			fmt.Print("first: ", string(dataArr[i][j]), "\n")
			first, err := strconv.Atoi(string(dataArr[i][j]))
			check(err)
			total += first * 10
			break
		}

		for j := len(dataArr[i]) - 1; j >= 0; j-- {
			fmt.Print(string(dataArr[i][j]), "\n")
			if dataArr[i][j] < 48 || dataArr[i][j] > 57 {
				continue
			}
			fmt.Print("last: ", string(dataArr[i][j]), "\n")
			last, err := strconv.Atoi(string(dataArr[i][j]))
			check(err)
			total += last
			break
		}
	}
	fmt.Print("Total: ", total, "\n")
}
