package main

import (
	"fmt"
	"sort"
	"strings"
)

// KnightTour resuelve el problema del salto del caballo
type KnightTour struct {
	tablero     [][]int
	tamaño      int
	movimientos [][]int
	visitado    int
}

// NewKnightTour crea una nueva instancia del problema
func NewKnightTour(tamaño int) *KnightTour {
	tablero := make([][]int, tamaño)
	for i := range tablero {
		tablero[i] = make([]int, tamaño)
	}
	// Movimientos posibles del caballo
	movimientos := [][]int{
		{2, 1}, {2, -1}, {-2, 1}, {-2, -1},
		{1, 2}, {1, -2}, {-1, 2}, {-1, -2},
	}
	return &KnightTour{
		tablero:     tablero,
		tamaño:      tamaño,
		movimientos: movimientos,
	}
}

// verificar si un movimiento es válido
func (kt *KnightTour) movimientoValido(x, y int) bool {
	return x >= 0 && x < kt.tamaño && y >= 0 && y < kt.tamaño && kt.tablero[x][y] == 0
}
// resolver el problema usando backtracking básico
func (kt *KnightTour) resolver(x, y, moverCuenta int) bool {
	// Marcar la posición actual
	kt.tablero[x][y] = moverCuenta
	kt.visitado++
	// Si hemos visitado todas las casillas
	if moverCuenta == kt.tamaño*kt.tamaño {
		return true
	}
	// Probar todos los movimientos posibles
	for _, movimiento := range kt.movimientos {
		siguienteX, siguienteY := x+movimiento[0], y+movimiento[1]
		if kt.movimientoValido(siguienteX, siguienteY) {
			if kt.resolver(siguienteX, siguienteY, moverCuenta+1) {
				return true
			}
		}
	}
	// Backtrack
	kt.tablero[x][y] = 0
	kt.visitado--
	return false
}

// calcular el grado de una posición (número de movimientos válidos)
func (kt *KnightTour) getGrado(x, y int) int {
	conteo := 0
	for _, movimiento := range kt.movimientos {
		siguienteX, siguienteY := x+movimiento[0], y+movimiento[1]
		if kt.movimientoValido(siguienteX, siguienteY) {
			conteo++
		}
	}
	return conteo
}
// Move representa un movimiento con su grado
type Movimiento struct {
	x, y, grado int
}

// resolverWarnsdorff resuelve usando la heurística de Warnsdorff
// (siempre elegir el movimiento que lleve a la casilla menos accesible)
func (kt *KnightTour) resolverWarnsdorff(x, y, moverCuenta int) bool {
	kt.tablero[x][y] = moverCuenta
	kt.visitado++

	if moverCuenta == kt.tamaño*kt.tamaño {
		return true
	}

	// Obtener todos los movimientos válidos con su grado
	var movimientoValido []Movimiento
	for _, movimiento := range kt.movimientos {
		siguienteX, siguienteY := x+movimiento[0], y+movimiento[1]
		if kt.movimientoValido(siguienteX, siguienteY) {
			grado := kt.getGrado(siguienteX, siguienteY)
			movimientoValido = append(movimientoValido, Movimiento{siguienteX, siguienteY, grado})
		}
	}
	// Ordenar por grado (menor grado primero - heurística de Warnsdorff)
	sort.Slice(movimientoValido, func(i, j int) bool {
		return movimientoValido[i].grado < movimientoValido[j].grado
	})
	// Probar movimientos en orden de menor grado
	for _, movimiento := range movimientoValido {
		if kt.resolverWarnsdorff(movimiento.x, movimiento.y, moverCuenta+1) {
			return true
		}
	}
	// Backtrack
	kt.tablero[x][y] = 0
	kt.visitado--
	return false
}

// imprimir el tablero con los números de movimiento
func (kt *KnightTour) imprimirTablero() {
	fmt.Println("Tablero del recorrido del caballo:")
	for _, fila := range kt.tablero {
		for _, celda := range fila {
			if celda == 0 {
				fmt.Printf("%3s", ".")
			} else {
				fmt.Printf("%3d", celda)
			}
		}
		fmt.Println()
	}
	fmt.Printf("Casillas visitadas: %d/%d\n", kt.visitado, kt.tamaño*kt.tamaño)
}

// printPath imprime el camino como coordenadas
func (kt *KnightTour) imprimirCamino() {
	// Crear mapa de posiciones
	posiciones := make(map[int][2]int)
	for i, fila := range kt.tablero {
		for j, movimiento := range fila {
			if movimiento > 0 {
				posiciones[movimiento] = [2]int{i, j}
			}
		}
	}
	fmt.Println("\nSecuencia de movimientos:")
	for i := 1; i <= kt.tamaño*kt.tamaño; i++ {
		pos := posiciones[i]
		fmt.Printf("Movimiento %2d: (%d, %d)\n", i, pos[0], pos[1])
	}
}

// reiniciar reinicia el tablero
func (kt *KnightTour) reiniciar() {
	for i := range kt.tablero {
		for j := range kt.tablero[i] {
			kt.tablero[i][j] = 0
		}
	}
	kt.visitado = 0
}

// mostrar el recorrido paso a paso
func (kt *KnightTour) verRecorrido() {
	// Crear representación visual con símbolos de ajedrez
	fmt.Println("\nVisualización del recorrido:")
	for i, fila := range kt.tablero {
		for j, celda := range fila {
			if celda == 1 {
				fmt.Print(" ♞ ") // Posición inicial del caballo
			} else if celda == kt.tamaño*kt.tamaño {
				fmt.Print(" ♜ ") // Posición final
			} else if celda > 0 {
				fmt.Print(" · ") // Casilla visitada
			} else {
				// Alternar colores como tablero de ajedrez
				if (i+j)%2 == 0 {
					fmt.Print(" □ ") // Casilla blanca
				} else {
					fmt.Print(" ■ ") // Casilla negra
				}
			}
		}
		fmt.Println()
	}
}

func main() {
	// Probar diferentes tamaños de tablero
	tamaños := []int{5, 6, 8}
	for _, tamaño := range tamaños {
		fmt.Printf("=== Tablero de %dx%d ===\n", tamaño, tamaño)
		kt := NewKnightTour(tamaño)
		// Probar desde la esquina (0,0)
		fmt.Printf("Intentando reresolverr desde posición (0,0)...\n")
		if tamaño <= 6 {
			// Para tableros pequeños, usar backtracking básico
			if kt.resolver(0, 0, 1) {
				fmt.Println("¡Solución encontrada con backtracking básico!")
				kt.imprimirTablero()
				if tamaño <= 5 {
					kt.imprimirCamino()
				}
				kt.verRecorrido()
			} else {
				fmt.Println("No se encontró solución con backtracking básico")
			}
		} else {
			// Para tableros más grandes, usar heurística de Warnsdorff
			if kt.resolverWarnsdorff(0, 0, 1) {
				fmt.Println("¡Solución encontrada con heurística de Warnsdorff!")
				kt.imprimirTablero()
				kt.verRecorrido()
			} else {
				fmt.Println("No se encontró solución con heurística de Warnsdorff")
			}
		}
		fmt.Println()
		// Probar desde el centro para tableros más grandes
		if tamaño >= 6 {
			kt.reiniciar()
			centroX, centroY := tamaño/2, tamaño/2
			fmt.Printf("Intentando desde el centro (%d,%d)...\n", centroX, centroY)
			if kt.resolverWarnsdorff(centroX, centroY, 1) {
				fmt.Println("¡Solución encontrada desde el centro!")
				kt.imprimirTablero()
				kt.verRecorrido()
			} else {
				fmt.Println("No se encontró solución desde el centro")
			}
		}
		fmt.Println("\n" + strings.Repeat("=", 50) + "\n")
	}
	// Información adicional sobre el problema
	fmt.Println("INFORMACIÓN SOBRE EL PROBLEMA DEL SALTO DEL CABALLO:")
	fmt.Println("- El caballo debe visitar cada casilla exactamente una vez")
	fmt.Println("- Se mueve en L: 2 casillas en una dirección y 1 en perpendicular")
	fmt.Println("- Tableros 5x5 y mayores siempre tienen solución")
	fmt.Println("- La heurística de Warnsdorff mejora significativamente el rendimiento")
	fmt.Println("- Para tableros grandes (8x8 o más), es casi instantáneo con Warnsdorff")
}
