package main

import "fmt"

var mass [][]float64 = [][]float64{
	{0, 1, 4, 9, 16},
	{1, 2, 5, 10, 17},
	{4, 5, 8, 13, 20},
	{9, 10, 13, 18, 25},
	{16, 17, 20, 25, 32},
}
var nx int = 4
var ny int = 3
var x float64 = 2.5
var y float64 = 3


func main() {


	//для начало необходимо проинтерполировать про постоянных y
	fmt.Print(MultidimensionalInterpolation(x, y, nx, ny, mass))
}

func MultidimensionalInterpolation(x float64, y float64, nx int, ny int, mass [][]float64) float64 {
	var resultMass [][]float64
	for i := 0; i < len(mass); i++ {
		resultMass = append(resultMass, make([]float64, 0))
		var massBuff [][]float64 = make([][]float64, 0)
		for j := 0; j < len(mass[0]); j++ {
			var buf []float64
			buf = append(buf, float64(j))
			buf = append(buf, mass[i][j])
			massBuff = append(massBuff, buf)
			fmt.Print(massBuff, "\n")
		}
		resultMass[i] = append(resultMass[i], float64(i))
		result := InterpolationNewton(x, massBuff, nx)
		resultMass[i] = append(resultMass[i], result)
	}
	return InterpolationNewton(y, resultMass, ny)
}

func InterpolationNewton(x float64, mass [][]float64, n int) float64 {

	//проверка есть ли число x в массиве
	for index := 0; index < len(mass); index++ {
		if mass[index][0] == x {
			return mass[index][1]
		}
	}

	// поиск интервала
	var indexUpInterval int
	var indexDownInterval int

	for i := 0; i < len(mass)-1; i++ {
		var indexUp float64 = mass[i][0]
		var indexDown float64 = mass[i+1][0]
		if indexUp < x && indexDown > x {
			indexUpInterval = i
			indexDownInterval = i + 1
		}
	}

	var middle int = n / 2

	if indexUpInterval-middle < 0 {
		indexUpInterval = 0
		indexDownInterval = n
	} else {
		indexUpInterval -= middle
		indexDownInterval += (n + 1) - middle
	}

	if indexDownInterval+middle > len(mass)-1 {
		indexDownInterval = len(mass) - 1
		indexUpInterval = (len(mass) - 1) - n
	} else {
		indexUpInterval -= middle
		indexDownInterval += (n + 1) - middle
	}
	var desiredArray [][]float64

	for i := indexUpInterval; i <= indexDownInterval; i++ {
		bufArray := []float64{mass[i][0], mass[i][1]}
		desiredArray = append(desiredArray, bufArray)
	}
	//Реализации подсчета коэффициентов

	indexСycle := n
	j := 0 // индекс для внутреннегшо цикла
	y := 1 //индекс для смещения по таблице вправо
	for indexСycle != 0 {
		for i := 0; i < indexСycle; i++ {
			number := (desiredArray[i][y] - desiredArray[i+1][y]) / (desiredArray[i][0] - desiredArray[i+1+j][0])
			desiredArray[i] = append(desiredArray[i], number)
		}
		y += 1
		indexСycle -= 1
		j += 1
	}

	result := desiredArray[0][1]
	xi := 0
	var buff float64 = 1
	for k := 2; k < len(desiredArray[0]); k++ {
		buff *= x - desiredArray[xi][0]
		result = result + buff*desiredArray[0][k]
		xi += 1
	}
	return result
}