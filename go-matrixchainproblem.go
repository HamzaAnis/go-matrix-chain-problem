package main

import (
	"fmt"
	"log"

	"github.com/apcera/termtables"
	"github.com/fatih/color"
)

//Matrix store the values of the matrix
type Matrix struct {
	row    int
	column int
}

func main() {
	color.Red("Matrix Chain Problem")
	limit := 0
	color.Cyan("Enter number of the Matrices:")
	fmt.Scanf("%d", &limit)
	// matrix := make([]matrix, limit)
	var matrix []Matrix
	p := make([]int, limit)
	var product = [10][10]int{}
	var split = [10][10]int{}

	for i := 0; i < limit; i++ {
		var input int
		localTemp := Matrix{0, 0}
		if i == 0 {
			color.Cyan("Enter rows for matrix %v : ", i+1)

			if _, err := fmt.Scan(&input); err != nil {
				log.Print("  Scan for i failed, due to ", err)
				return
			}

			localTemp.row = input
			color.Cyan("Enter column for matrix %v : \n", i+1)
			if _, err := fmt.Scan(&input); err != nil {
				log.Print("  Scan for i failed, due to ", err)
				return
			}
			localTemp.column = input
			p[i] = localTemp.row

			matrix = append(matrix, localTemp)

			color.Yellow("Rows : %v  Columns : %v", matrix[i].row, matrix[i].column)
		} else {
			color.Cyan("Enter number of column for the matrix %v : ", i+1)
			if _, err := fmt.Scan(&input); err != nil {
				log.Print("  Scan for i failed, due to ", err)
				return
			}
			localTemp.column = input
			localTemp.row = matrix[i-1].column
			matrix = append(matrix, localTemp)
			p[i] = matrix[i].row
			if i == limit-1 {
				p = append(p, matrix[i].column)
			}
			color.Yellow("Rows : %v  Columns : %v", matrix[i].row, matrix[i].column)

		}
	}

	displayTable(matrix)

	for l := 0; l < limit; l++ {
		color.Red("In loop")
		for i := 0; i < limit-l; i++ {
			color.Red("In 2nd loop")
			j := i + l
			color.Red("J is", j)
			product[i][j] = -1
			for k := i; k < j; k++ {
				color.Red("In 3rd loop")
				cost := product[i][k] + product[k+1][j] + (p[i] * p[k+1] * p[j+1])
				if product[i][j] == -1 {
					product[i][j] = cost
				} else if cost < product[i][j] {
					product[i][j] = cost
					split[i][j] = k
				}
			}
		}
	}

	color.Yellow("%v", product[0][limit-1])

	for i := 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			if split[i][j] > 0 && split[i][j] < limit {
				color.Blue("For i: %v for j: %v split at: %v ", i, j, split[i][j])
			}
		}
	}

}
func displayTable(matrix []Matrix) {
	table := termtables.CreateTable()
	table.AddHeaders("Number", "Rows", "Number")
	for i := 0; i < len(matrix); i++ {
		table.AddRow(i+1, matrix[i].row, matrix[i].column)
	}
	color.Yellow(table.Render())
}
