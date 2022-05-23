package main

import (
	"fmt"
	"os/user"
)

func main() {
	v, err := HelloWorld()
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
}

func HelloWorld() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Hello %v", u.Name), nil
}
