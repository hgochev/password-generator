package main

import (
	"fmt"
	"os"

	"github.com/hgochev/password-generator/internal/flags"
	"github.com/hgochev/password-generator/internal/password"
)

func main() {
	opts := flags.ParseArguments()

	passwd, err := password.Generate(opts)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to generate password: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(passwd)
}
