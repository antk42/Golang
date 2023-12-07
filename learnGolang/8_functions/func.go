package main

import "fmt"

func main() {
	myString := addPrefix("myString")
	fmt.Println(myString)

	myString, err := addPrefixWithErr("MyStringWithErr")
	fmt.Println(myString, err)

	myString, length := addPrefixWithLen("MyStringWithLen")
	fmt.Println(myString, length)
}

func addPrefix(origin string) string {
	return "prefix_" + origin
}

func addPrefixWithErr(origin string) (string, error) {
	return "prefix_" + origin, nil

}

func addPrefixWithLen(origin string) (res string, length int) {
	res = "prefix_" + origin
	length = len(res)
	return res, length

}
