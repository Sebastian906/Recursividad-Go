package main

import "fmt"

// Función recursiva para imprimir un ArrayList (slice)
func imprimirArrayList(arr []interface{}, indice int) {
	// Caso base: cuando el índice alcanza la longitud del slice
	if indice >= len(arr) {
		return
	}
	// Imprimir el elemento actual
	fmt.Printf("Elemento [%d]: %v\n", indice, arr[indice])
	// Llamada recursiva para el siguiente elemento
	imprimirArrayList(arr, indice+1)
}

func main() {
	// Crear un ArrayList con diferentes tipos de datos
	arrayList := []interface{}{10, "Hola", 3.14, true, 'A', struct{}{}}
	fmt.Println("Contenido del ArrayList: ")
	imprimirArrayList(arrayList, 0)
}