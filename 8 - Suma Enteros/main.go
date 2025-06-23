package main

import "fmt"

func sumaEntera(N int) (int, []string) {
	var suma int
	var errores []string

	if N <= 0 {
		errores = append(errores, fmt.Sprintf("Error: %d no es un entero positivo", N))
		return suma, errores
	}

	if N%2 != 0 {
		errores = append(errores, fmt.Sprintf("Error: %d no es un número par", N))
		N-- // Reducir N en 1 para empezar desde el número par anterior
	}

	for i := N; i >= 2; i -= 2 {
		suma += i
	}

	return suma, errores
}

func main() {
	var numero int

	fmt.Printf("Ingrese un número entero positivo N: ")
	fmt.Scan(&numero)

	suma, errores := sumaEntera(numero)

	if len(errores) > 0 {
		fmt.Println("\nMensaje de error: ")
		for _, err := range errores {
			fmt.Println(err)
		}
	}

	fmt.Printf("\nLa suma de los números pares desde %d hasta 2 es: %d\n", numero, suma)
}