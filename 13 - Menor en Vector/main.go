package main

import (
	"fmt"
	"math"
)

func menorVector(vec []int, index int, actual int) int {
	if index >= len(vec) {
		return actual
	}
	if vec[index] < actual {
		return menorVector(vec, index+1, vec[index])
	}
	return menorVector(vec, index+1, actual)
}

func main() {
	vector := []int{5, 3, 8, 4, 9, 2}
	menor := menorVector(vector, 0, math.MaxInt32)
	fmt.Println("El número menor es: ", menor)
}