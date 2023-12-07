package main

import "fmt"

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

func (c Client) HasAvatar() bool {
	if c.IMG.URL != "" {
		return true
	}
	return false
}

func (c *Client) UpdateAvatar() {
	c.IMG.URL = "new_url"
}

func main() {

	// some pic -> true
	client := Client{
		ID:   777,
		Name: "Andy777",
		Age:  100,
		IMG: Avatar{
			URL:  "some_url",
			Size: 10,
		},
	}
	fmt.Println(client.HasAvatar())
	fmt.Printf("%#v\n", client)
	client.UpdateAvatar()
	fmt.Printf("%#v\n", client)
}
