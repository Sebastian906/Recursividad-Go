package main

import "fmt"

func contarRepeticiones(vec []int, num, index int) int {
	if index >= len(vec) {
		return 0
	}
	if vec[index] == num {
		return 1 + contarRepeticiones(vec, num, index+1)
	}
	return contarRepeticiones(vec, num, index+1)
}

func main() {
	vector := []int{2, 3, 2, 5, 2, 7, 2}
	numero := 2
	fmt.Printf("El número %d aparece %d veces\n", numero, contarRepeticiones(vector, numero, 0))
}