package main

import "fmt"

// Función recursiva / Parámetro: veces --> Saludos restantes
func saludar(veces int) {
	// Caso base: Cuando llegamos a 0, se termina la recursión
	if veces <= 0 {
		return
	}
	// Se imprime el saludo
	fmt.Println("Hola mundo")
	// Llamada recursiva
	saludar(veces - 1)
}

func main() {
	saludar(10)
}