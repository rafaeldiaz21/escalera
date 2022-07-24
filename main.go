package main

import (
	"fmt"
	"sort"
)

func main() {

	var cases [][]int

	var s0 = []int{3, 4, 5, 6, 7}
	var s1 = []int{11, 10, 13, 12, 1}
	var s2 = []int{14, 2, 5, 3, 4}
	var s3 = []int{3, 4, 12, 6, 3}
	var s4 = []int{1, 10, 11, 12, 13, 14, 32}
	var s5 = []int{1, 10, 11, 12, 13, 14, 7, 3}
	var s6 = []int{7, 7, 12, 11, 3, 4, 14}
	var s7 = []int{1, 10, 11, 12, 13, 14, 3, 2, 3, 4}
	var s8 = []int{7, 3, 2}

	cases = append(cases, s0, s1, s2, s3, s4, s5, s6, s7, s8)

	for _, v := range cases {
		fmt.Println(fmt.Sprintln(v, "Is Straight =", isStraight(v)))
	}

}

func isStraight(cartas []int) bool {
	var cartas2 []int
	var verify bool

	// Validando Request
	if !validCards(cartas) {
		return false
	}
	//Verificando la escalera con las cartas dadas
	if verify = verifyStraight(cartas); verify {
		return true
	}
	//Verificando la escalera cambiando el 1 por el 14
	cartas2 = changeCard(1, cartas)
	if verify = verifyStraight(cartas2); verify {
		return true
	}
	//Verificando la escalera cambiando el 14 por el 1
	cartas2 = changeCard(14, cartas)
	if verify = verifyStraight(cartas2); verify {
		return true
	}

	return false
}

func verifyStraight(cartas []int) bool {
	var array []int
	cartas = unique(cartas)
	sort.Ints(cartas)
	array = generateStraight(cartas)
	return len(array) >= 5
}

//Cambia el AS de posición
func changeCard(card int, cartas []int) []int {
	cartas = unique(cartas)
	var count = len(cartas)
	var cambio int
	switch card {
	case 14:
		cambio = 1
	case 1:
		cambio = 14
	}
	for i := 0; i < count; i++ {
		if cartas[i] == card {
			cartas[i] = cambio
		}
	}
	return cartas
}

//Devuelve un arreglo de cartas que estén aumentando de 1 en 1
func generateStraight(cartas []int) []int {
	var array []int
	var count = len(cartas)

	for i := 0; i < count-1; i++ {
		if (cartas[i+1] - cartas[i]) == 1 {
			array = append(array, cartas[i])
			break
		}
	}
	for i := 0; i < count-1; i++ {
		if (cartas[i+1] - cartas[i]) == 1 {
			array = append(array, cartas[i+1])
		}

	}
	return array
}

//Elimina los elementos repetidos de un arreglo []int
func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}

	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

//Valida las cartas
func validCards(cartas []int) bool {
	var count = len(cartas)
	//Rechazando si tiene más de 7 cartas
	if count > 7 {
		fmt.Println("Número de cartas permitidas excedido")
		return false
	}
	//Rechazando si no tiene cartas suficientes para hacer escalera
	if count < 5 {
		fmt.Println("Número de cartas insuficientes")
		return false
	}
	//Rechazando si hay cartas inválidas
	for i := 0; i < len(cartas); i++ {
		if cartas[i] < 1 || cartas[i] > 14 {
			fmt.Printf("La carta %d no es válida \n", cartas[i])
			return false
		}
	}
	return true
}
