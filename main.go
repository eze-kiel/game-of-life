package main

import (
	"fmt"
	"math/rand"
	"time"
)

//FieldSize represents the size of the board for cells development
const FieldSize = 40

//Height represents the height of the field
const Height = 40

//Lenght represents the lenght of the field
const Lenght = 140

//InitCells represents the number of cells at the beginning
const InitCells = Height + Lenght

//Speed represents the duration of one generation in ms
const Speed = 100

var field = [Height][Lenght]string{}

func main() {
	counter := 0
	print("\033[H\033[2J")
	fmt.Printf("gen : %d\n", counter)
	fillField()
	createRandomCells()
	printField()

	//life loop
	for {
		time.Sleep(Speed * time.Millisecond)
		counter = counter + 1

		//clear the terminal
		print("\033[H\033[2J")

		fmt.Printf("generation : %d\n", counter)
		birth()
		survival()
		printField()
	}
}

func fillField() {
	for r := 0; r < Height; r++ {
		for c := 0; c < Lenght; c++ {
			field[r][c] = " "
		}
	}
}

func createRandomCells() {
	for cellsNumber := 0; cellsNumber < InitCells; cellsNumber++ {
		source1 := rand.NewSource(time.Now().UnixNano())
		source2 := rand.NewSource(time.Now().UnixNano())

		randomRow := rand.New(source1)
		randomCol := rand.New(source2)

		field[randomRow.Intn(Height)][randomCol.Intn(Lenght)] = "#"

	}
}

func printField() {
	for r := 0; r < Height; r++ {
		for c := 0; c < Lenght; c++ {
			fmt.Printf("%s", field[r][c])
		}
		fmt.Printf("\n")
	}
}

func survival() {
	//if a cell has 2 or 3 neighbors only

	for r := 0; r < Height; r++ {
		for c := 0; c < Lenght; c++ {

			neighbors := countNeighbors(r, c)

			if field[r][c] == "#" {
				if neighbors == 2 || neighbors == 3 {
					//the cell survives : nothing changes
				} else {
					field[r][c] = " " //the cells dies
				}
			}

		}

	}
}

func birth() {
	//if a blank has 3 neighbors

	for r := 0; r < Height; r++ {
		for c := 0; c < Lenght; c++ {
			neighbors := countNeighbors(r, c)
			//fmt.Printf("neighbors : %d\n", neighbors)

			if field[r][c] == " " {
				if neighbors == 3 {
					field[r][c] = "#" //a cell is born
				}
			}
		}
	}
}

//returns the number of neighbors of a specific set of row/col
func countNeighbors(row int, col int) int {
	nbNeighbors := 0

	if row == 0 || col == 0 || row == Height-1 || col == Lenght-1 { //are we checking the edge of the field ?

		if row == 0 {
			switch col {
			case 0:
				for rowToTest := row; rowToTest < row+2; rowToTest++ {
					for colToTest := col; colToTest < col+2; colToTest++ {
						if rowToTest == row && colToTest == col {
						} else {
							if field[rowToTest][colToTest] == "#" {
								nbNeighbors++
							}
						}
					}
				}
			case Lenght - 1:
				for rowToTest := row; rowToTest < row+2; rowToTest++ {
					for colToTest := col - 1; colToTest < col+1; colToTest++ {
						if rowToTest == row && colToTest == col {

						} else {
							if field[rowToTest][colToTest] == "#" {
								nbNeighbors++
							}
						}

					}
				}
			default:
				for rowToTest := row; rowToTest < row+2; rowToTest++ {
					for colToTest := col - 1; colToTest < col+2; colToTest++ {
						if rowToTest == row && colToTest == col {

						} else {
							if field[rowToTest][colToTest] == "#" {
								nbNeighbors++
							}
						}

					}
				}
			}

		}

		if row == Height-1 {
			switch col {
			case 0:
				for rowToTest := row - 1; rowToTest < row+1; rowToTest++ {
					for colToTest := col; colToTest < col+2; colToTest++ {
						if rowToTest == row && colToTest == col {

						} else {
							if field[rowToTest][colToTest] == "#" {
								nbNeighbors++
							}
						}

					}
				}
			case Lenght - 1:
				for rowToTest := row - 1; rowToTest < row+1; rowToTest++ {
					for colToTest := col - 1; colToTest < col+1; colToTest++ {
						if rowToTest == row && colToTest == col {

						} else {
							if field[rowToTest][colToTest] == "#" {
								nbNeighbors++
							}
						}

					}
				}
			default:
				for rowToTest := row - 1; rowToTest < row+1; rowToTest++ {
					for colToTest := col - 1; colToTest < col+2; colToTest++ {
						if rowToTest == row && colToTest == col {

						} else {
							if field[rowToTest][colToTest] == "#" {
								nbNeighbors++
							}
						}

					}
				}
			}
		}

		if col == 0 {
			if row > 0 && row < Height-1 {
				for rowToTest := row - 1; rowToTest < row+2; rowToTest++ {
					for colToTest := col; colToTest < col+2; colToTest++ {
						if rowToTest == row && colToTest == col {

						} else {
							if field[rowToTest][colToTest] == "#" {
								nbNeighbors++
							}
						}

					}
				}
			}
		}

		if col == Lenght-1 {
			if row > 0 && row < Height-1 {
				for rowToTest := row - 1; rowToTest < row+2; rowToTest++ {
					for colToTest := col - 1; colToTest < col+1; colToTest++ {
						if rowToTest == row && colToTest == col {

						} else {
							if field[rowToTest][colToTest] == "#" {
								nbNeighbors++
							}
						}

					}
				}
			}
		}
	} else {
		for rowToTest := row - 1; rowToTest < row+2; rowToTest++ {
			for colToTest := col - 1; colToTest < col+2; colToTest++ {
				if rowToTest == row && colToTest == col {

				} else {
					if field[rowToTest][colToTest] == "#" {
						nbNeighbors++
					}
				}

			}
		}

	}
	return nbNeighbors
}
