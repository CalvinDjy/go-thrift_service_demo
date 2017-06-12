package main

import (
	"fmt"
	"boot"
)

func main()  {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	boot.NewServer().Start()
}
