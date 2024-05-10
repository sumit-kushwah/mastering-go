package osmodule

import (
	"fmt"
	"os"
)

func checkGetgidHere() {
	fmt.Println("Inside checkGetgidHere")
	fmt.Println(os.Getgid())
	fmt.Println(os.Getpid())
	fmt.Println(os.Getppid())
	fmt.Println("Exiting checkGetgidHere")
}

func MainOsBasics() {
	// zeroth value in args is program name
	fmt.Println(os.Args)
	fmt.Println(os.ExpandEnv("my home path is $HOME"))
	fmt.Println(os.Getgid())  // group id
	fmt.Println(os.Geteuid()) // effective user id
	go checkGetgidHere()
	fmt.Println(os.Getpid())
	fmt.Println(os.Getppid())
	fmt.Println(os.Getgroups())
	fmt.Println(os.Hostname())
}
