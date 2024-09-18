package main

/*
В данной лабораторной работе необходимо реализовать алгоритм интерполяции.
Использовать будем полином Ньютона.
*/

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
