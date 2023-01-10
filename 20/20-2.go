package main

import "fmt"

func matrixMultiplication(mA [][]int, mB [][]int) [][]int {
	// Новая матрица, которую верну
	mC := make([][]int, len(mA))
	// Размеры новой матрицы
	numberOfLines := len(mA)
	numberOfColumns := len(mB[0])

	// Строка новой матрицы
	for i := 0; i < numberOfLines; i++ {
		// Столбец новой матрицы
		for j := 0; j < numberOfColumns; j++ {

			// Подсчитаю новый элемент
			newElem := 0
			m := 0
			for _, elem := range mA[i] {
				newElem += elem * mB[m][j]
				m++
			}

			mC[i] = append(mC[i], newElem)
		}
	}
	return mC
}

func main() {
	matrixA := [][]int{
		{1, 0, -3, 9, 3},
		{2, -7, 11, 5, 0},
		{-9, 4, 25, 84, 1},
	}
	matrixB := [][]int{
		{9, 13, -15, 2},
		{1, 11, 7, 6},
		{-3, 7, 4, 1},
		{6, 0, 7, 10},
		{8, 9, -2, 4},
	}
	var matrixC [][]int

	matrixC = matrixMultiplication(matrixA, matrixB)
	fmt.Println(matrixC)
}
