package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// Guardar contactos en un file jSon
func SaveContacts(contacts []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file) //convierte structs en JSON
	err = encoder.Encode(contacts)   //codificar el slice de contactos en formato json y lo escribe en file
	if err != nil {
		return err
	}
	return nil //no falló nada y todo se guardó ok
}

func GetContacts(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()

	//recuperar el Contenido dle file JSON
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}
	return nil

}

func main() {
	var contacts []Contact
	err := GetContacts(&contacts)
	if err != nil {
		log.Println("Failed to get contacts")
	}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("==GESTOR DE CONTACTOS==\n",
			"1. Agregar\n",
			"2. Mostrar\n",
			"3. Eliminar\n",
			"4. Salir\n",
			"Ingrese una opcion: ")
		//Leer opcion del usuario
		var option int
		_, err = fmt.Scanln(&option)
		if err != nil {
			log.Println("Error al leer la opción", err)
			return
		}
		//Manejar la opcion del usuario
		switch option {

		case 1:
			//Ingresar y crear contacto
			var c Contact
			fmt.Print("Nombre: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Phone: ")
			c.Phone, _ = reader.ReadString('\n')
			//se agrega el contacto nuevo al slice
			contacts = append(contacts, c)
			if err := SaveContacts(contacts); err != nil {
				log.Println("Failed to save new contact", err)
			}

		case 2:
			//Mostrar contactos
			fmt.Print("============================================================\n")
			for index, c := range contacts {
				fmt.Printf("\n%d. Name: %s Email: %s Phone: %s\n",
					index+1, c.Name, c.Email, c.Phone)
			}
			fmt.Print("============================================================\n")
		case 3:
			//Eliminar contacto
			fmt.Print("Ingrese el numero del contacto que desea eliminar: ")
			var index int
			_, err = fmt.Scanln(&index)
			if err != nil {
				log.Println("Error al leer la opción", err)
				return
			}
			if index > len(contacts) {
				fmt.Println("Contacto no encontrado")
				continue
			}
			contacts = append(contacts[:index], contacts[index+1:]...)
			if err := SaveContacts(contacts); err != nil {
				log.Println("Failed to save new contact", err)
			}
		case 4:
			//Salir
			fmt.Println("Bye")
			return
		default:
			fmt.Println("¡¡¡¡¡  Invalid option   !!!!!!")
		}
	}
}
