package main

import "fmt"

// Función recusiva, con los parámetros del vector (array) e índice del vector
func invertirVector(vector []int, indice int) []int {
	// Caso base: cuando el índice llega al final del vector
	if indice >= len(vector) {
		// Retorna un vector vacío
		return []int{}
	}
	// Llamda recursiva para invertir vector
	invertido := invertirVector(vector, indice+1)
	// Construye el nuevo vector invertido
	return append(invertido, vector[indice])
}

func main() {
	vectorOriginal := []int{1, 2, 3, 4, 5}
	fmt.Println("Vector original: ", vectorOriginal)

	vectorInvertido := invertirVector(vectorOriginal, 0)
	fmt.Println("Vector invertido: ", vectorInvertido)
}