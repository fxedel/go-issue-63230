package main

import (
	"fmt"
	"os"
	"plugin"

	// _ "golang.org/x/sys/unix"
	_ "golang.org/x/xerrors"
)

func main() {
	fmt.Println("main")
	fmt.Println(os.Getwd())

	fmt.Println("load plugin-a:")
	fmt.Println(plugin.Open("./build/plugin-a.so"))

	fmt.Println("load plugin-b:")
	fmt.Println(plugin.Open("./build/plugin-b.so"))
}
