package main

import "fmt"

func imprimirVector(vector []int, indice int) {
	// Caso base: cuando llegamos al final del vector
	if indice >= len(vector) {
		return
	}
	// Imprimir el elemento actual
	fmt.Printf("vector[%d] = %d\n", indice, vector[indice])
	// Llamada recursiva para el siguiente elemento
	imprimirVector(vector, indice+1)
}

func main() {
	vector := []int{10, 20, 30, 40, 50}
	fmt.Println("Impresión recursiva del vector: ")
	imprimirVector(vector, 0)
}