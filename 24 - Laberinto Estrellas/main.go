package main

import "fmt"

type Posicion struct {
	x, y int
}

func coleccionarEstrellas(laberinto [][]int, x, y int, estrellas map[Posicion]bool, camino []Posicion, visitado [][]bool, todosCaminos *[][]Posicion) {
	filas, columnas := len(laberinto), len(laberinto[0])
	// Verificar limites
	if x < 0 || x >= filas || y < 0 || y >= columnas {
		return
	}
	// Verificar si es pared o ya fue visitado
	if laberinto[x][y] == 1 || visitado[x][y] {
		return
	}
	// Marcar como visitado
	visitado[x][y] = true
	// Agregar posición actual al camino
	caminoActual := append(camino, Posicion{x, y})
	// Si hay una estrella, recolectarla
	estrellaActual := make(map[Posicion]bool)
	for k, v := range estrellas {
		estrellaActual[k] = v
	}
	if laberinto[x][y] == 2 {
		estrellaActual[Posicion{x, y}] = true
	}
	// Verificar si hemos recolectado todas las estrellas
	totalEstrellas := contarEstrellas(laberinto)
	if len(estrellaActual) == totalEstrellas {
		// Guardar el camino exitoso
		copiaCamino := make([]Posicion, len(caminoActual))
		copy(copiaCamino, caminoActual)
		*todosCaminos = append(*todosCaminos, copiaCamino)
		visitado[x][y] = false
		return
	}
	// Explorar las 4 direcciones
	direcciones := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, dir := range direcciones {
		nuevoX, nuevoY := x+dir[0], y+dir[1]
		coleccionarEstrellas(laberinto, nuevoX, nuevoY, estrellaActual, caminoActual, visitado, todosCaminos)
	}
	// Backtrack
	visitado[x][y] = false
}

// contar el número total de estrellas en el laberinto
func contarEstrellas(laberinto [][]int) int {
	conteo := 0
	for i := range laberinto {
		for j := range laberinto[i] {
			if laberinto[i][j] == 2 {
				conteo++
			}
		}
	}
	return conteo
}

// Encontrar el camino más corto que recolecta todas las estrellas
func encontrarCaminoCorto(laberinto [][]int, inicioX, inicioY int) []Posicion {
	filas, columnas := len(laberinto), len(laberinto[0])
	visitado := make([][]bool, filas)
	for i := range visitado {
		visitado[i] = make([]bool, columnas)
	}
	var todosCaminos [][]Posicion
	estrellas := make(map[Posicion]bool)
	coleccionarEstrellas(laberinto, inicioX, inicioY, estrellas, []Posicion{}, visitado, &todosCaminos)
	// Encontrar el camino más corto
	if len(todosCaminos) == 0 {
		return nil
	}
	caminoMasCorto := todosCaminos[0]
	for _, camino := range todosCaminos {
		if len(camino) < len(caminoMasCorto) {
			caminoMasCorto = camino
		}
	}
	return caminoMasCorto
}

// versión optimizada usando DFS con poda
func resolverLaberinto(laberinto [][]int, inicioX, inicioY int) []Posicion {
	totalEstrellas := contarEstrellas(laberinto)
	if totalEstrellas == 0 {
		return []Posicion{{inicioX, inicioY}}
	}
	filas, columnas := len(laberinto), len(laberinto[0])
	visitado := make([][]bool, filas)
	for i := range visitado {
		visitado[i] = make([]bool, columnas)
	}
	var mejorCamino []Posicion
	masLargo := int(^uint(0) >> 1) // Infinito
	var dfs func(x, y int, estrellasColeccionadas int, camino []Posicion)
	dfs = func(x, y int, estrellasColeccionadas int, camino []Posicion) {
		// Poda: si el camino actual ya es más largo que el mejor, abandonar
		if len(camino) >= masLargo {
			return
		}
		// Verificar límites y obstáculos
		if x < 0 || x >= filas || y < 0 || y >= columnas || laberinto[x][y] == 1 || visitado[x][y] {
			return
		}
		visitado[x][y] = true
		caminoActual := append(camino, Posicion{x, y})
		// Contar estrellas recolectadas
		if laberinto[x][y] == 2 {
			estrellasColeccionadas++
		}
		// Si recolectamos todas las estrellas
		if estrellasColeccionadas == totalEstrellas {
			if len(caminoActual) < masLargo {
				masLargo = len(caminoActual)
				mejorCamino = make([]Posicion, len(caminoActual))
				copy(mejorCamino, caminoActual)
			}
			visitado[x][y] = false
			return
		}
		// Explorar direcciones
		direcciones := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for _, dir := range direcciones {
			nuevoX, nuevoY := x+dir[0], y+dir[1]
			dfs(nuevoX, nuevoY, estrellasColeccionadas, caminoActual)
		}
		visitado[x][y] = false
	}
	dfs(inicioX, inicioY, 0, []Posicion{})
	return mejorCamino
}

// Imprimir el laberinto con estrellas
func imprimirLaberinto(laberinto [][]int) {
	for _, fila := range laberinto {
		for _, celda := range fila {
			switch celda {
			case 0:
				fmt.Print("  ") // Libre
			case 1:
				fmt.Print("█ ") // Pared
			case 2:
				fmt.Print("★ ") // Estrella
			}
		}
		fmt.Println()
	}
}

// imprimr la solución con el camino y estrellas
func imprimirSolucion(laberinto [][]int, camino []Posicion) {
	// Crear mapa de posiciones del camino
	caminoMapa := make(map[Posicion]int)
	for i, pos := range camino {
		caminoMapa[pos] = i + 1
	}
	for i, fila := range laberinto {
		for j, celda := range fila {
			pos := Posicion{i, j}
			if paso, incamino := caminoMapa[pos]; incamino {
				if celda == 2 {
					fmt.Printf("★ ")
				} else if paso == 1 {
					fmt.Printf("S ") // inicio
				} else {
					fmt.Printf("%d ", paso%10)
				}
			} else {
				switch celda {
				case 0:
					fmt.Print("  ")
				case 1:
					fmt.Print("█ ")
				case 2:
					fmt.Print("☆ ") // Estrella no recolectada
				}
			}
		}
		fmt.Println()
	}
}

func main() {
	// Laberinto del ejercicio 2 con estrellas
	// 0 = libre, 1 = pared, 2 = estrella
	laberinto := [][]int{
		{0, 1, 1, 1, 1, 1, 1, 0},
		{0, 2, 0, 2, 0, 0, 0, 0},
		{1, 0, 1, 0, 1, 2, 1, 0},
		{0, 0, 0, 2, 0, 0, 0, 0},
		{0, 1, 1, 1, 1, 1, 1, 0},
	}
	inicioX, inicioY := 0, 0
	fmt.Println("Laberinto con estrellas:")
	imprimirLaberinto(laberinto)
	fmt.Printf("\nTotal de estrellas: %d\n\n", contarEstrellas(laberinto))
	// Encontrar el camino óptimo
	camino := resolverLaberinto(laberinto, inicioX, inicioY)
	if camino != nil {
		fmt.Printf("¡Solución encontrada! Camino de %d pasos:\n", len(camino))
		imprimirSolucion(laberinto, camino)
		fmt.Println("\nSecuencia de movimientos:")
		for i, pos := range camino {
			fmt.Printf("Paso %d: (%d, %d)", i+1, pos.x, pos.y)
			if laberinto[pos.x][pos.y] == 2 {
				fmt.Print(" ★ ¡Estrella recolectada!")
			}
			fmt.Println()
		}
	} else {
		fmt.Println("No se encontró solución para recolectar todas las estrellas")
	}
}