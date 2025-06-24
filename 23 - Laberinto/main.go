package main

import "fmt"

func resolverLaberinto(laberinto [][]int, x, y, finX, finY int, visitado [][]bool) bool {
	filas, columnas := len(laberinto), len(laberinto[0])
	// Verificar limites
	if x < 0 || x >= filas || y < 0 || y >= columnas {
		return false
	}
	// Verificar si es una pared o ya fue visitada
	if laberinto[x][y] == 1 || visitado[x][y] {
		return false
	}
	// Marcar como visitado
	visitado[x][y] = true
	// Verificar si llegamos al destino
	if x == finX && y == finY {
		return true
	}
	// Explorar direcciones
	direcciones := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, dir := range direcciones {
		nuevoX, nuevoY := x+dir[0], y+dir[1]
		if resolverLaberinto(laberinto, nuevoX, nuevoY, finX, finY, visitado) {
			return true
		}
	}
	// Backtrack: desmarcar casilla actual
	visitado[x][y] = false
	return false
}

func buscarCamino(laberinto [][]int, inicioX, inicioY, finX, finY int) [][]int {
	filas, columnas := len(laberinto), len(laberinto[0])
	visitado := make([][]bool, filas)
	camino := make([][]int, filas)

	for i := range visitado {
		visitado[i] = make([]bool, columnas)
		camino[i] = make([]int, columnas)
	}

	if resolverLaberintoCamino(laberinto, inicioX, inicioY, finX, finY, visitado, camino) {
		return camino
	}
	return nil
}

func resolverLaberintoCamino(laberinto [][]int, x, y, finX, finY int, visitado [][]bool, camino [][]int) bool {
	rows, cols := len(laberinto), len(laberinto[0])
	if x < 0 || x >= rows || y < 0 || y >= cols {
		return false
	}
	if laberinto[x][y] == 1 || visitado[x][y] {
		return false
	}
	visitado[x][y] = true
	camino[x][y] = 1 // Marcar como parte del camino
	if x == finX && y == finY {
		return true
	}
	direcciones := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, dir := range direcciones {
		nuevoX, nuevoY := x+dir[0], y+dir[1]
		if resolverLaberintoCamino(laberinto, nuevoX, nuevoY, finX, finY, visitado, camino) {
			return true
		}
	}
	camino[x][y] = 0 // Quitar del camino (backtrack)
	return false
}

func imprimirLaberinto(laberinto [][]int) {
	for _, fila := range laberinto {
		for _, celda := range fila {
			if celda == 1 {
				fmt.Print("█ ") // Pared
			} else {
				fmt.Print("  ") // Camino libre
			}
		}
		fmt.Println()
	}
}

func imprimirSolucion(laberinto [][]int, camino [][]int) {
	for i, fila := range laberinto {
		for j, celda := range fila {
			if celda == 1 {
				fmt.Print("█ ") // Pared
			} else if camino[i][j] == 1 {
				fmt.Print("* ") // Camino solución
			} else {
				fmt.Print("  ") // Espacio libre
			}
		}
		fmt.Println()
	}
}

func main() {
	laberinto := [][]int{
		{0, 1, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 1, 1, 1, 1, 1, 0},
	}
	inicioX, inicioY := 0, 0  // Posición inicial (naranja)
	finX, endY := 4, 7     // Posición final (naranja)
	fmt.Println("Laberinto original:")
	imprimirLaberinto(laberinto)
	fmt.Println()
	// Resolver el laberinto
	camino := buscarCamino(laberinto, inicioX, inicioY, finX, endY)
	if camino != nil {
		fmt.Println("¡Solución encontrada!")
		fmt.Println("Camino marcado con *:")
		imprimirSolucion(laberinto, camino)
	} else {
		fmt.Println("No se encontró solución")
	}
}