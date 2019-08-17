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
const Lenght = 40

//InitCells represents the number of cells at the beginning
const InitCells = Height + Lenght

//Speed represents the number of generations that evolve in 1 sec
const Speed = 1

var field_ori = [Height][Lenght]string{}

var field_tmp = [Height][Lenght]string{}

func main() {
	counter := 0
	print("\033[H\033[2J")
	fmt.Printf("generation : %d\n", counter)
	fillField()
	//createRandomCells()
	createOscillator()
	printField()

	//life loop
	for {
		time.Sleep(1000 / Speed * time.Millisecond)
		counter = counter + 1

		//clear the terminal
		print("\033[H\033[2J")

		fmt.Printf("generation : %d\n", counter)
		copyOriToTmp()
		birth()
		survival()
		copyTmpToOri()
		printField()
	}
}

func fillField() {
	for r := 0; r < Height; r++ {
		for c := 0; c < Lenght; c++ {
			field_ori[r][c] = " "
		}
	}
}

func createRandomCells() {
	for cellsNumber := 0; cellsNumber < InitCells; cellsNumber++ {
		source1 := rand.NewSource(time.Now().UnixNano())
		source2 := rand.NewSource(time.Now().UnixNano())

		randomRow := rand.New(source1)
		randomCol := rand.New(source2)

		field_ori[randomRow.Intn(Height)][randomCol.Intn(Lenght)] = "#"

	}
}

func printField() {
	for r := 0; r < Height; r++ {
		for c := 0; c < Lenght; c++ {
			fmt.Printf("%s", field_ori[r][c])
		}
		fmt.Printf("\n")
	}
}

func survival() {
	//if a cell has 2 or 3 neighbors only

	for r := 0; r < Height; r++ {
		for c := 0; c < Lenght; c++ {

			neighbors := countNeighbors(r, c)

			if field_ori[r][c] == "#" {
				if neighbors == 2 || neighbors == 3 {
					//the cell survives : nothing changes
				} else {
					field_tmp[r][c] = " " //the cells dies
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

			if field_ori[r][c] == " " {
				if neighbors == 3 {
					field_tmp[r][c] = "#" //a cell is born
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
							if field_ori[rowToTest][colToTest] == "#" {
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
							if field_ori[rowToTest][colToTest] == "#" {
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
							if field_ori[rowToTest][colToTest] == "#" {
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
							if field_ori[rowToTest][colToTest] == "#" {
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
							if field_ori[rowToTest][colToTest] == "#" {
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
							if field_ori[rowToTest][colToTest] == "#" {
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
							if field_ori[rowToTest][colToTest] == "#" {
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
							if field_ori[rowToTest][colToTest] == "#" {
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
					if field_ori[rowToTest][colToTest] == "#" {
						nbNeighbors++
					}
				}

			}
		}

	}
	return nbNeighbors
}

func createOscillator() {
	for col := 5; col < 8; col++ {
		field_ori[10][col] = "#"
	}
	for col := 4; col < 7; col++ {
		field_ori[11][col] = "#"
	}
}

func copyOriToTmp() {
	for row := 0; row < Height; row++ {
		for col := 0; col < Lenght; col++ {
			field_tmp[row][col] = field_ori[row][col]
		}

	}
}

func copyTmpToOri() {
	for row := 0; row < Height; row++ {
		for col := 0; col < Lenght; col++ {
			field_ori[row][col] = field_tmp[row][col]
		}

	}
}
