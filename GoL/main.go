package main

import (
	"fmt"
	"math/rand"
)

func main() {
	/*Punto de entrada de la simulación*/
	fmt.Println("Iniciando simulación")
	var universo = inicializarUniverso(20, 20)
	//fmt.Println(universo)
	var generaciones = 10

	/*universo[0] = []int{0, 0, 0}
	universo[1] = []int{1, 1, 1}
	universo[2] = []int{0, 0, 0}*/
	var poblacion = crearGeneracionExpontanea(universo)
	imprimirUniverso(poblacion)

	for gen := 0; gen < generaciones; gen++ {
		poblacion = pasarGeneracion(poblacion)
		//time.Sleep(1 * time.Second)
		imprimirUniverso(poblacion)
	}
}

/*El universo es un arreglo de 0's y 1's en donde
un 0 es una célula muerta y un 1 e una célula viva*/
func inicializarUniverso(ancho, alto int) [][]int {
	//var universo = [5][alto]int{}
	var universo = make([][]int, alto)
	for i := range universo {
		universo[i] = make([]int, ancho)
	}
	return universo
}

/*Imprime en consola un estado del universo específico*/
func imprimirUniverso(universo [][]int) {
	var ancho = len(universo)
	var alto = len(universo[0])

	for x := 0; x < ancho; x++ {
		for y := 0; y < alto; y++ {
			fmt.Print(universo[x][y])
		}
		fmt.Println()
	}
	fmt.Println()
}

/*Agrega vida al universo de forma pseudo-aleatoria*/
func crearGeneracionExpontanea(universo [][]int) [][]int {
	var ancho = len(universo)
	var alto = len(universo[0])

	for x := 0; x < ancho; x++ {
		for y := 0; y < alto; y++ {
			universo[x][y] = rand.Intn(2)
		}
	}
	return universo
}

/*Juzga al individuo*/
func pruebaDeVida(universo [][]int, coordinadaX, coordenadaY int) int {
	var estadoSalud = 0
	var estado = universo[coordinadaX][coordenadaY] //1 para viva, 0 para muerta
	var vecinos = contarVecinos(universo, coordinadaX, coordenadaY)
	if estado == 1 {
		if vecinos < 2 {
			estadoSalud = 0 //muere por soledad
		}
		if vecinos == 2 || vecinos == 3 {
			estadoSalud = 1 //sobrevive por comunidad
		}
		if vecinos > 3 {
			estadoSalud = 0 //muere por sobrepoblación
		}
	} else {
		if vecinos == 3 {
			estadoSalud = 1
		}
	}
	return estadoSalud
}

/*Hace un senso de la vecindad del individuo y retorna la cantidad de vecinos vivos*/
func contarVecinos(universo [][]int, coordinadaX, coordenadaY int) int {
	var cantidad = 0
	var xFinalUniverso = len(universo)
	var yFinalUniverso = len(universo[0])
	//validar por puntos. son 8 en total
	if vecinoExiste(0, 0, xFinalUniverso, yFinalUniverso, coordinadaX-1, coordenadaY-1) {
		cantidad = cantidad + analizarVecino(universo, coordinadaX-1, coordenadaY-1)
	}
	if vecinoExiste(0, 0, xFinalUniverso, yFinalUniverso, coordinadaX-1, coordenadaY) {
		cantidad = cantidad + analizarVecino(universo, coordinadaX-1, coordenadaY)
	}
	if vecinoExiste(0, 0, xFinalUniverso, yFinalUniverso, coordinadaX-1, coordenadaY+1) {
		cantidad = cantidad + analizarVecino(universo, coordinadaX-1, coordenadaY+1)
	}
	if vecinoExiste(0, 0, xFinalUniverso, yFinalUniverso, coordinadaX, coordenadaY-1) {
		cantidad = cantidad + analizarVecino(universo, coordinadaX, coordenadaY-1)
	}
	if vecinoExiste(0, 0, xFinalUniverso, yFinalUniverso, coordinadaX, coordenadaY+1) {
		cantidad = cantidad + analizarVecino(universo, coordinadaX, coordenadaY+1)
	}
	if vecinoExiste(0, 0, xFinalUniverso, yFinalUniverso, coordinadaX+1, coordenadaY-1) {
		cantidad = cantidad + analizarVecino(universo, coordinadaX+1, coordenadaY-1)
	}
	if vecinoExiste(0, 0, xFinalUniverso, yFinalUniverso, coordinadaX+1, coordenadaY) {
		cantidad = cantidad + analizarVecino(universo, coordinadaX+1, coordenadaY)
	}
	if vecinoExiste(0, 0, xFinalUniverso, yFinalUniverso, coordinadaX+1, coordenadaY+1) {
		cantidad = cantidad + analizarVecino(universo, coordinadaX+1, coordenadaY+1)
	}
	return cantidad
}

/*Analizar vida del vecino*/
func analizarVecino(universo [][]int, x, y int) int {
	var sentencia = 0
	if universo[x][y] == 1 {
		sentencia = 1
	}
	return sentencia
}

/*Localiza al vecino*/
func vecinoExiste(xInicioUniverso, yInicioUniverso, xFinalUniverso, yFinalUniverso, x, y int) bool {
	var existencia = false
	if x >= xInicioUniverso && x < xFinalUniverso && y >= yInicioUniverso && y < yFinalUniverso {
		existencia = true
	}
	return existencia
}

/*Cumple las leyes de la vida*/
func iusVitae(generacionVieja [][]int) [][]int {
	var ancho = len(generacionVieja)
	var alto = len(generacionVieja[0])
	var generacionNueva = inicializarUniverso(alto, ancho)
	for x := 0; x < ancho; x++ {
		for y := 0; y < alto; y++ {
			generacionNueva[x][y] = pruebaDeVida(generacionVieja, x, y)
		}
	}
	return generacionNueva
}

/*Avanza la generación sustituye la original con una nueva*/
func pasarGeneracion(generacionAlfa [][]int) [][]int {
	return iusVitae(generacionAlfa)
}
