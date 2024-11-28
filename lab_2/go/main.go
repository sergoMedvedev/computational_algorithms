package main

import (
	"fmt"
)

/*
В данной лабораторной работе необходимо реализовать алгоритм интерполяции.
Использовать будем полином Ньютона.
*/

var mass [][]float64 = [][]float64{
	{0.0,  1.0,       -1},
	{0.15, 0.838771,  -1,14944},
	{0.3,  0.655336,  -1.29552},
	{0.45, 0.450447,  -1.43497},
	{0.6,  0.225336,  -1.56464},
	{0.75, -0.018310, -1.68164},
	{0.9,  -0.278390, -1.78333},
	{1.05, -0.552430, -1.86742},
}

var mass1 [][]float64 = [][]float64{
	{1, 2 ,0.5},
	{2, 3, 1},
	{3, 5, 1.5},
}

var x float64 = 0.3
var n int = 5

func main() {
	resultN := InterpolationNewton(x, mass, n)
	resultH := HermiteInterpolation(mass, x) 
	fmt.Println("Полином Ньютона: ", resultN)
	fmt.Println("Полином Эрмита: ", resultH)
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
			// raz1 := math.Round((desiredArray[i][y] - desiredArray[i+1][y])*1000)/1000
			// fmt.Println("Первая разность ", raz1)
			// raz2 := math.Round((desiredArray[i][0] - desiredArray[i+1+j][0])*1000)/1000
			// fmt.Println("Вторая разность ", raz2)
			// if math.IsNaN(raz1) {
			// 	fmt.Println(raz1)
			// 	return 0
			// }
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

// HermiteInterpolation выполняет интерполяцию Эрмита
// data - массив точек, где каждая строка содержит [x, y, y']
// targetX - точка, для которой нужно вычислить интерполяцию
func HermiteInterpolation(data [][]float64, targetX float64) float64 {
	// Количество узлов
	n := len(data)
	if n == 0 {
	 panic("data cannot be empty")
	}
   
	// Построение вспомогательных массивов z и Q
	m := 2 * n
	z := make([]float64, m)
	Q := make([][]float64, m)
	for i := range Q {
	 Q[i] = make([]float64, m)
	}
   
	// Заполнение значений z и начальных Q
	for i := 0; i < n; i++ {
	 z[2*i] = data[i][0]       // x_i
	 z[2*i+1] = data[i][0]     // x_i (дублируем)
	 Q[2*i][0] = data[i][1]    // y_i
	 Q[2*i+1][0] = data[i][1]  // y_i (дублируем)
	 Q[2*i+1][1] = data[i][2]  // y'_i (первая производная)
	 if i > 0 {
	  Q[2*i][1] = (Q[2*i][0] - Q[2*i-1][0]) / (z[2*i] - z[2*i-1])
	 }
	}
   
	// Вычисление разделённых разностей
	for j := 2; j < m; j++ {
	 for i := 0; i < m-j; i++ {
	  Q[i][j] = (Q[i+1][j-1] - Q[i][j-1]) / (z[i+j] - z[i])
	 }
	}
   
	// Вычисление многочлена Эрмита
	result := Q[0][0]
	product := 1.0
	for i := 1; i < m; i++ {
	 product *= (targetX - z[i-1])
	 result += Q[0][i] * product
	}
   
	return result
   }