package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	c, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Printf("Ошибка: некорректное число: %s\n", os.Args[1])
		return
	}
	fmt.Printf("%.1f\n", c*9/5+32)

}