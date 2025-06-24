package main

import "fmt"

func sumaMatriz(matriz [][]int, fila, col int) int {
	if fila >= len(matriz) {
		return 0
	}
	if col >= len(matriz[fila]) {
		return sumaMatriz(matriz, fila+1, 0)
	}
	return matriz[fila][col] + sumaMatriz(matriz, fila, col+1)
}

func main() {
	matriz := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Suma de matriz: ", sumaMatriz(matriz, 0, 0))
}