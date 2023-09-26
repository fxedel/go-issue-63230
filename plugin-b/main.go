package main

import (
	"fmt"

	// _ "golang.org/x/sys/unix"
	_ "golang.org/x/xerrors"
)

func init() {
	fmt.Println("plugin-b init!")
}
