package main

import (
	"fmt"

	"github.com/jisdisai/godesde0/variables"
)

func main() {
	//variables.MuestroEnteros()
	//variables.RestoVariables()

	estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado, texto)

}

// correr un programa en go es go run main.go
// correr un build con exe go build main.go
