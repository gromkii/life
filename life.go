package main

import (
  "fmt"
)

/*
  We will need a grid to display the game on, comprised of
  a 2d array of booleans, a width, and a height. True = alive,
  false = dead.
*/
type Grid struct {
  cell [][]bool
  width, height int
}

/*
  We need a method to initialize a new grid, as well as set each
  cell to false by default. Returns the Grid type we just created.
  The go documentation says to point and reference Grid when
  creating a NewStruct
*/
func NewGrid(width, height int) *Grid {
  // Define the 2d array to being looping through.
  cell := make([][]bool, height)

  // Use range to sum the number of elements in the array
  for i := range cell {
    cell[i] = make([]bool, width)
    for j:= range cell[i] {
      cell[i][j] = false
    }
  }

  return &Grid{cell:cell, width:width, height:height}
}

/*
  This method will print the current board of our game state.
*/

func PrintGrid(grid *Grid) string {
  printGrid := ""

  for i := 0; i < grid.height; i++ {
    line := ""

    for j := 0; j < grid.height; j++ {
      if grid.cell[i][j] == true {
        line += "🐰"
      } else {
        line += " "
      }
    }

    printGrid += line + "\n"
  }

  return printGrid
}

func main(){
  fmt.Println("Hello from Go!")
  newGrid := NewGrid(15, 15)
  newGrid.cell[5][5] = true
  gridded := PrintGrid(newGrid)
  fmt.Println(gridded)
}
