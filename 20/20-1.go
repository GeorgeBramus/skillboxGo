package main

import "fmt"

func determinant(m [3][3]int) float32 {
	return float32(m[0][0]*m[1][1]*m[2][2] - m[0][0]*m[1][2]*m[2][1] - m[0][1]*m[1][0]*m[2][2] + m[0][1]*m[1][2]*m[2][0] + m[0][2]*m[1][0]*m[2][1] - m[0][2]*m[1][1]*m[2][0])
}

func main() {
	var matrix = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("Определитель матрицы:")
	fmt.Printf("%.2f\n", determinant(matrix))
}
