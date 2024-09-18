package main

import (
	"fmt"

	"github.com/olebedev/when"
)

/*
В данной лабораторной работе необходимо реализовать алгоритм интерполяции.
Использовать будем полином Ньютона.
*/

func main() {
	mass := [][]float64{
		{0.5, 0.707},
		{0.25, 0.924},
		{0, 1},
		{0.25, 0.924},
		{0.5, 0.707},
		{0.75, 0.383},
		{1, 0},
	}

	var x float64 = 0.6
	n := 4
	InterpolationNewton(x, mass, n)
}

func InterpolationNewton(x float64, mass [][]float64, n int) {

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
	print(fmt.Sprintln("Индексы двух крайних значений: %d %d", indexUpInterval, indexDownInterval))

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

	for indexСycle != 2 {
		j := 0 // индекс для внутреннегшо цикла
		y := 1 //индекс для смещения по таблице вправо
		for i := 0; i < indexСycle; i++ {
			desiredArray[i] = append(desiredArray, (desiredArray[i][y] - desiredArray[i+1][y])/(desiredArray[i][0]-desiredArray[i+1+j][0]))
			y+=1
			indexСycle-=1
			j+=1
		} 
	}

}
