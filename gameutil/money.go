package gameutil

import (
	"bufio"
	"github.com/gookit/color"
	"strconv"
	"strings"
)

func StartingMoney(reader *bufio.Reader) int64 {
	result := 500
	for {
		color.Warn.Println("\nHow much money you would like to play with? (Must be greater than 0 and less than 1000001):")

		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		var err error
		result, err = strconv.Atoi(text)
		if err == nil && result > 0 && result < 1000001 {
			break
		} else {
			color.Error.Println("Value was invalid!")
		}
	}
	return int64(result)
}

func StartingBet(reader *bufio.Reader, totalMoney int64) int64 {
	result := int64(-1)
	for {
		color.Warn.Println("\nHow much money you would like to bet on this round? (Must be greater than 0 and less than or equal to", totalMoney, "Dollars):")

		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		temp, err := strconv.Atoi(text)
		result = int64(temp)
		if err != nil {
			color.Error.Println("Value was invalid!")
		} else if result <= 0 {
			color.Error.Println("Cannot bet less than or equal to 0 dollars!")
		} else if result > totalMoney {
			color.Error.Println("Cannot bet more than the money you have!")
		} else {
			break
		}
	}
	return result
}
