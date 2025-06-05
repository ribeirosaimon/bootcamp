package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Client struct {
	Id      uint
	File    string
	Name    string
	Phone   string
	Address string
}
type ClientsFile struct {
	Clients map[string]*Client
}

func (cf ClientsFile) Save(c Client) (*Client, error) {
	if _, ok := cf.Clients[strconv.Itoa(int(c.Id))]; ok {
		return nil, errors.New("this client already exists")
	}
	newClient := &Client{
		Id: c.Id,
	}
	cf.Clients[strconv.Itoa(int(c.Id))] = newClient
	return newClient, nil
}

var (
	fileName    = "aula07/ex03/customers.txt"
	clientsFile = ClientsFile{
		Clients: make(map[string]*Client),
	}
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("File not found")
	}
	defer func() {
		fmt.Println("closing file")
		file.Close()
	}()

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		var c Client
		if err = json.Unmarshal([]byte(line), &c); err != nil {
			panic(err)
		}
		clientsFile.Clients[strconv.Itoa(int(c.Id))] = &c
	}

	if err = fileScanner.Err(); err != nil {
		panic(err)
	}

	for {
		var (
			idString string
			id       int
		)
		fmt.Println("Digite o id")
		if _, err = fmt.Scanln(&idString); err != nil {
			panic(err)
		}
		if id, err = strconv.Atoi(idString); err != nil {
			panic(err)
		}

		if _, err = clientsFile.Save(Client{
			Id: uint(id),
		}); err != nil {
			panic(err)
		}
	}

}
