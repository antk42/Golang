package main

import (
	"fmt"
)

type Avatar struct {
	URL  string
	Size int64
}

type Client struct {
	ID   int
	Name string
	Age  int
	IMG  Avatar
}

func main() {
	i := new(string)
	_ = i

	client := Client{}
	fmt.Printf("%#v\n", client)

	updateAvatar(&client)
	fmt.Printf("%#v\n", client)
	fmt.Printf("%#v\n", client.IMG)

	updateClient(&client)
	fmt.Printf("%#v\n", client)

}

func updateAvatar(client *Client) {
	client.IMG.URL = "updated_url"
}

func updateClient(client *Client) {
	client.Name = "Updated_name"
}
