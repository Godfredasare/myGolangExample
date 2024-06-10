package main

import (
	"fmt"

	"github.com/Godfredasare/myGolangExample.git/address_book/address"
)

func main() {
	addressBook, err := address.ReadSaveAddress()
	if err != nil {
		fmt.Print("Error: ", err)
		return
	}

	
	id := address.GetUserInput("Enter ID: ")
	name := address.GetUserInput("Enter Name: ")
	email := address.GetUserInput("Enter Email: ")
	addres := address.GetUserInput("Enter Address: ")
	
	add, err := address.New(id, name, email, addres)
	if err != nil {
		fmt.Print("Error: ", err)
		return
	}

	
	addressBook = append(addressBook, *add)

	err = address.Save(addressBook)
	if err != nil {
		fmt.Println("Error saving the file")
		return 
		}

	add.PrintOutput()

	//search for an item
	search := address.GetUserInput("Search Item: ")
	searches := address.Search(search, addressBook)
	fmt.Println(searches)

	//delete item by index
	index := 2
	addressBook =append(addressBook[:index], addressBook[index+1:]...)
	fmt.Println(addressBook)

}
