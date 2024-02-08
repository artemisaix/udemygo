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
	divide(10, 3)

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
	defer func() { //si se produce un panico lo recupera y se ejecuta hasta
		//el final de la funcion, y no interrumpe el flujo de la aplicacion
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	fmt.Println(dividendo / divisor)
}
