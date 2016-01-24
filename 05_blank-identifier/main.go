package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, _ := http.Get("http://www.google.com")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}

/*
Usando el blank identifier _ hacemos asignación pero descartamos la variable error
para no tener declaradas variables que no nos interesan o que no usaremos... porque Go da error
si no usamos variables que han sido declaradas
*/
