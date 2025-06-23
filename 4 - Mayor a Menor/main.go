package main

import "fmt"

// Función quicksort descendente
func quicksort(arr []int) []int {
	// Caso base: arreglo vacío o con un solo elemento ya está ordenado
	if len(arr) <= 1 {
		return arr
	}
	// Seleccionar el pivote (primer elementos)
	pivot := arr[0]
	// Partir el arreglo en 3 partes
	var menores, mayores, iguales []int
	// Recorrido a los valores del arreglo
	for _, num := range arr {
        if num > pivot {
            mayores = append(mayores, num)
        } else if num < pivot {
            menores = append(menores, num)
        } else {
            iguales = append(iguales, num)
        }
    }
	// Ordenar recursivamente las partes y combinar (mayores primero)
    return append(quicksort(mayores), append(iguales, quicksort(menores)...)...)
}

// Función alterna con bubblesort
func bubblesort(arr []int, n int) {
	// Caso base
	if n == 1 {
		return
	}
	for i := 0; i < n-1; i++ {
        if arr[i] < arr[i+1] {
            arr[i], arr[i+1] = arr[i+1], arr[i]
        }
    }
	bubblesort(arr, n-1)
}

func main() {
	numeros := []int{5, 2, 9, 1, 5, 6, 3, 8, 4}
	fmt.Println("Arreglo original:", numeros)
	
	numerosOrdenados := quicksort(numeros)
	fmt.Println("Arreglo ordenado descendente:", numerosOrdenados)

	numerosBubblesort := make([]int, len(numeros))
	copy(numerosBubblesort, numeros)
	bubblesort(numerosBubblesort, len(numerosBubblesort))
	fmt.Println("Arreglo ordenado con bubblesort:", numerosBubblesort)
}