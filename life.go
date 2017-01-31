package main

import (
  "fmt"
  "math/rand"
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
        line += "🐰 "
      } else {
        line += "- "
      }
    }

    printGrid += line + "\n"
  }

  return printGrid
}

/*
  We need a method to seed the initial grid. Takes a grid
  as an argument, and populates each based on a random value
  of 0 or 1. 1 initalizes it as alive, 0 is ded.
*/
func InitGrid(grid *Grid) {
  for i:=0; i < grid.height; i++ {
    for j:=0; j < grid.width; j++ {
      coin := rand.Intn(2)

      if coin == 1 {
        grid.cell[i][j] = true
      }
    }
  }
}

/*
  Get cell value at a certain x y coordinate.
*/
func GetCell(grid *Grid, x, y int) bool {
  return grid.cell[y][x]
}

/*
  This method will return if the index for the next available
  cell is valid, returning false if not, true if it is. Kinda
  swapped this logic. Wraps around if index is nonexistent.
*/
func Next(grid *Grid, x, y int) bool {
  x += grid.width
	x %= grid.width
	y += grid.height
	y %= grid.height
	return grid.cell[y][x]
}

/*
  Count the number of nearby cells, return a sum of alive neighbors, need
  to figure how to handle when the index is non-existant.
*/
func GetAliveNeighbors(grid *Grid, x, y int) int {
  count := 0

  for i:=-1; i <= 1; i++ {
    for j:=-1; j <= 1; j++ {
      if (j != 0 || i != 0) && Next(grid, x+i, y+j) {
        count++
      }
    }
  }

  return count
}

func NextGeneration(grid *Grid) {
  for i:=0; i < grid.height; i++ {
    for j:=0; j < grid.width; j++ {
      var current = grid.cell[i][j]
      count := GetAliveNeighbors(grid, i , j)

      if current && (count == 3 || count == 2) && Next(grid, i, j) {
        grid.cell[i][j] = true
      } else if !current && count == 3 && Next(grid, i, j){
        grid.cell[i][j] = true
      } else {
        grid.cell[i][j] = false
      }

    }
  }
}



func main(){
  fmt.Println("Hello from Go!")
  newGrid := NewGrid(30, 30)
  InitGrid(newGrid)
  fmt.Println(PrintGrid(newGrid))
  for i:=0; i < 10; i++ {
    NextGeneration(newGrid)
    fmt.Println(PrintGrid(newGrid))
  }
}
