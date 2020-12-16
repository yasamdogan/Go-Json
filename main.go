package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Date    int    `json:"date"`
	Social Social `json:"social"`
}

type Social struct {
	Twitter  string `json:"twitter"`
	Website  string `json:"website"`
}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("users.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users Users

	json.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Users); i++ {

		fmt.Println("Type: " + users.Users[i].Type)
		fmt.Println("Programming Name: " + users.Users[i].Name)
		fmt.Println("Programming Age: " + strconv.Itoa(users.Users[i].Date))
		fmt.Println("Twitter Url: " + users.Users[i].Social.Twitter)
		fmt.Println("Website Url: " + users.Users[i].Social.Website)
	}

}