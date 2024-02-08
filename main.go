package main

import (
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

	divide(10, 2)
	divide(10, 0)

	fmt.Println(num)

	//uso de defer
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

	//uso del panic

}

func divide(dividendo, divisor int) {
	validateZero(divisor)
	fmt.Println(dividendo / divisor)
}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("ERROR: No se puede dividir por 0")
	}
}
