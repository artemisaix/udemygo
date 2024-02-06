package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	str := "123"
	num, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(num)
	fmt.Println(result)

	file, err := os.Create("hola.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	_, err = file.Write([]byte("Hola Andre"))
	if err != nil {
		fmt.Println("ErrorD: ", err)
		return
	}
	defer file.Close()

}

func divide(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("ERROR: No se puede dividir por 0")
	}
	return dividendo / divisor, nil
}
