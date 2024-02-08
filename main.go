package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	log.SetPrefix("main():")
	str := "123"
	num, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	divide(10, 2)
	// divide(10, 0)
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

	file, err = os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.SetOutput(file) // lo escribe en un archivo de log
	log.Print("Oye,soy un log en el file")
	defer file.Close() //se cierra el fujo del archivo de log

	//uso de logs()

	//log.Panic("No se puede continuar") //detiene la ejecucion del programa
	log.Print("Mensaje de registros...")
	//log.Fatal("No se puede continuar") //detiene la ejecucion del programa
	log.Println("Este es otro mensaje de registros...")

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
