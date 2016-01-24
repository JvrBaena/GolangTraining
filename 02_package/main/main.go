package main

import (
	"fmt"

	"github.com/JvrBaena/udemyTraining/02_package/stringutil"
)

func main() {
	fmt.Println(stringutil.MyName) // MyName expuesto en el paquete, _myName no
	stringutil.PrintName("LOL")
}
