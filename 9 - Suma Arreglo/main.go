package main

import "fmt"

// Función recursiva
func sumaArray(arr []int, index int) int {
	// Caso base: cuando se llega al final del arreglo
	if index >= len(arr) {
		return 0
	}

	// Paso recursivo: suma el elemento actual con la suma del resto del arreglo
	return arr[index] + sumaArray(arr, index+1)
}

func main() {
	numeros := []int{1, 2, 3, 4, 5}

	fmt.Println("Arreglo original: ", numeros)

	suma := sumaArray(numeros, 0)

	fmt.Println("Suma total: ", suma)

	arreglo := []int{10, 20, 30, 40}
	fmt.Println("\nSegundo arreglo: ", arreglo)
	fmt.Println("Suma total: ", sumaArray(arreglo, 0))
}