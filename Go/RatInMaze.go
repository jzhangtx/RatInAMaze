package main

import (
	"fmt"
	"strconv"
)

func GetArrayFromInput(count int, prompt string) []int {
	if count == 0 {
		return []int{}
	}

	fmt.Print(prompt)

	v := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scan(&v[i])
	}

	return v
}

func GetCheese(maze [][]int, row, col int) bool {
	if row == len(maze)-1 && col == len(maze[0])-1 {
		return true
	}

	if row == len(maze) || col == len(maze[0]) {
		return false
	}

	if maze[row][col] == 0 {
		return false
	}

	return GetCheese(maze, row+1, col) || GetCheese(maze, row, col+1)
}

func CanGetCheese(maze [][]int) bool {
	return GetCheese(maze, 0, 0)
}

func main() {
	for {
		fmt.Print("The size of the maze (rows, cols): ")
		var rows, cols int
		fmt.Scan(&rows, &cols)
		if rows == 0 && cols == 0 {
			break
		}

		maze := make([][]int, rows)
		for i := 0; i < rows; i++ {
			prompt := "Please enter the " + strconv.Itoa(i) + "th row: "
			maze[i] = GetArrayFromInput(cols, prompt)
		}

		f := CanGetCheese(maze)
		fmt.Print("The rat can ")
		if !f {
			fmt.Print("not ")
		}
		fmt.Println("reach the cheese!")
	}
}
