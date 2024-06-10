package address

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Address struct {
	ID      string
	Name    string
	Email   string
	Address string
}

func (d Address) PrintOutput() {
	fmt.Printf("Your address name is %v, email is %v and address is %v\n\ngo run ./address_book", d.Name, d.Email, d.Address)
}

func New(id string, name string, email string, addres string) (*Address, error) {
	if id == "" || name == "" || email == "" || addres == "" {
		return &Address{}, errors.New("all fields are required")
	}
	address := &Address{
		ID:      id,
		Name:    name,
		Email:   email,
		Address: addres,
	}


	return address, nil
}

func Save(address []Address) error {
	file, err := os.Create("address.json")

	if err != nil {
		return err
	}

	defer file.Close()

	json, err := json.MarshalIndent(address, "", "")
	if err != nil {
		return err
	}

	_, err = file.Write(json)
	return err

}

func ReadSaveAddress() ([]Address, error) {
	data, err := os.ReadFile("address.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Address{}, nil
		}
		return nil, err
	}

	var addressBook []Address

	err = json.Unmarshal(data, &addressBook)
	if err != nil {
		return []Address{}, err
	}
	return addressBook, nil

}

func Search(searchItem string, addresses []Address) ([]Address) {
    var results []Address

	for _, address := range addresses {
		if strings.Contains(address.ID, searchItem) || strings.Contains(address.Name, searchItem) || strings.Contains(address.Email, searchItem){
			results = append(results, address)
		}
	}
	return results
}

func GetUserInput(prompt string) string {
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value)
	return value
}
