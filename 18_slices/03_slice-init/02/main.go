package main

import "fmt"

func main() {
	conShorthand := []int{} // Con este patrón obtenemos un slice pero hemos inicializado el array subyacente a length CERO
	// conShorthand[0] = "PACO" -> Esto daría error. NO tenemos acceso por índice
	// conShorthand = append(conShorthand, "PACO") -> Habría que hacerlo con append porque append se encarga de hacer sitio si no hay
	fmt.Println(conShorthand)
	fmt.Println(conShorthand == nil) //Esto es falso

}
