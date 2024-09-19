package main

import (
	"cli_app/app"
	"fmt"
	"os"
)

func main() {
	application := app.Generate()
	erro := application.Run(os.Args)
	if erro != nil {
		fmt.Println(erro)
	}
}
