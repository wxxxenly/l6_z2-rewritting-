package main

import (
	"fmt"
)

func shiftLayer(matrix [][]int, k int, p int) {

	n := len(matrix)
	m := len(matrix[0])

	top := k - 1
	bottom := n - k
	left := k - 1
	right := m - k

	for i := left; i < right-p; i++ {
		matrix[top][i+p] = matrix[top][i]
	}
	for i := top + 1; i < bottom; i++ {
		matrix[i][right-1] = matrix[i-1][right-1]
	}
	for i := right - 2; i >= left+p; i-- {
		matrix[bottom-1][i-p] = matrix[bottom-1][i]
	}
	for i := bottom - 2; i > top; i-- {
		matrix[i][left] = matrix[i+1][left]
	}

	for i := left; i < right-p; i++ {
		matrix[top][i] = 0
	}
	for i := top + 1; i < bottom; i++ {
		matrix[i-1][right-1] = 0
	}
	for i := right - 2; i >= left+p; i-- {
		matrix[bottom-1][i] = 0
	}
	for i := bottom - 2; i > top; i-- {
		matrix[i+1][left] = 0
	}
}

func main() {

	var n, m, k, p int
	fmt.Print("Введите количество строк матрицы: ")
	fmt.Scan(&n)
	fmt.Print("Введите количество столбцов матрицы: ")
	fmt.Scan(&m)
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, m)
	}
	fmt.Println("Введите элементы матрицы:")
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}
	fmt.Print("Введите номер слоя: ")
	fmt.Scan(&k)
	fmt.Print("Введите количество шагов сдвига вправо: ")
	fmt.Scan(&p)

	shiftLayer(matrix, k, p)

	fmt.Println("Результат сдвига элементов слоя матрицы:")
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}
