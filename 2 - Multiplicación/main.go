package main

import "fmt"

// Función auxiliar de la tabla con un parámetro de número
func tablaMultiplicar(numero int) {
	fmt.Printf("Tabla de multiplicar del %d:\n", numero)

	// Recorrido de los primeros 10 múltiplos
	for i := 1; i < 11; i++ {
		// Variable de resultado que multiplica al parámetro con la posición
		resultado := numero * i
		fmt.Printf("%d x %d = %d\n", numero, i , resultado)
	}
}

// Función principal del programa
func main() {
	// Variable del número para hallar su tabla
	var numero int

	fmt.Println("Ingrese un número para generar su tabla de multiplicar: ")
	fmt.Scan(&numero)

	// Llamdo a la función recursiva
	tablaMultiplicar(numero)
}