package main

import "fmt"

// Estrategia 1 - Lista Enlazada
type Nodo struct {
	valor int
	siguiente *Nodo
}

// Función recursiva para imprimir la lista
func imprimirListaRecursiva(n *Nodo) {
	// Caso base: Llegar al final de la lista
	if n == nil {
		return
	}
	// Imprimir el valor actual
	fmt.Println(n.valor)
	imprimirListaRecursiva(n.siguiente)
}

// Función recursiva para imprimir un slice (representación de lista)
func imprimirListaSlice(lista []int, indice int) {
	// Caso base: Llegar al final del slice
	if indice >= len(lista) {
		return
	}
	// Imprimir elemento actual
	fmt.Printf("Elemento %d: %d\n", indice, lista[indice])
	// Llamda recursiva para el siguiente elemento
	imprimirListaSlice(lista, indice+1)
}

func main() {
	// Crear una lista enlazada
	lista := &Nodo{10, &Nodo{20, &Nodo{30, &Nodo{40, &Nodo{50, nil}}}}}
	fmt.Println("Elementos de la lista: ")
	imprimirListaRecursiva(lista)

	// Crear un slice
	slice := []int{10, 20, 30, 40, 50}
	fmt.Println("Elementos del slice: ")
	imprimirListaSlice(slice, 0)
}