package main

import "fmt"

// Función recursiva
func invertirLista(original, invertida []int) []int {
	if len(original) == 0 {
		return invertida
	}

	return invertirLista(original[:len(original)-1], append(invertida, original[len(original)-1]))
}

func main() {
	lista := []int{1, 2, 3, 4, 5}
	fmt.Println("Original: ", lista)

	invertida := invertirLista(lista, []int{})
	fmt.Println("Invertida: ", invertida)
}