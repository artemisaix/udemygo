package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	str := "123"
	num, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	result, err := divide(10, 5)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(num)
	fmt.Println(result)
}

func divide(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("ERROR: No se puede dividir por 0")
	}
	return dividendo / divisor, nil
}
