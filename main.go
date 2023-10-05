package main

import (
	"Users/Anyelo/Documentos/cursos/ECommerceGoNext/go/src/Ang3l023/goInitial/variables"
	"fmt"
)

func main() {
	estado, texto := variables.ConviertoATexto(1520)
	fmt.Println(texto)
	fmt.Println(estado)
}
