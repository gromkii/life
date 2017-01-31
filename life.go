package main

import (
  "fmt"
  "math/rand"
  "time"
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
      seed := rand.Intn(4)

      if seed == 1 {
        grid.cell[i][j] = true
      }
    }
  }
}

/*
  Check to see if the current index of grid.cell is valid, return
  true or false depending on if it is or not.
*/
func (grid *Grid) Next(x, y int) bool {
  if x < 0 || x >= grid.width {
    return false
  }

  if y < 0 || y >= grid.height {
    return false
  }

  return true
}

/*
  Count the number of nearby cells, return a sum of alive neighbors, need
  to figure how to handle when the index is non-existant.
*/
func (grid *Grid)GetAliveNeighbors(x, y int) int {
  count := 0

  for i := -1; i <= 1; i++ {
    for j := -1; j <= 1; j++ {
      if grid.Next(x+i, j+y) && grid.cell[i+x][j+y] == true && !(i == 0 && j == 0) {
        count++
      }
    }
  }

  return count
}

func NextGeneration(grid *Grid) {
  for x:=0; x < grid.height; x++ {
    for y:=0; y < grid.width; y++ {
      // Get # of alive neighbors.
      count := grid.GetAliveNeighbors(x, y)
      if grid.cell[x][y] && ( count == 3 || count == 2) {
        grid.cell[x][y] = true
      } else if grid.cell[x][y] == false && count == 3 {
        grid.cell[x][y] = true
      } else if grid.cell[x][y] && (count > 3 || count < 2) {
        grid.cell[x][y] = false
      } else {
        grid.cell[x][y] = false
      }
    }
  }
}



func main(){
  fmt.Println("Hello from Go!")
  newGrid := NewGrid(15, 15)
  InitGrid(newGrid)
  fmt.Println(PrintGrid(newGrid))
  for i:=0; i < 50; i++ {
    NextGeneration(newGrid)
    fmt.Println(PrintGrid(newGrid))
    time.Sleep(500 * time.Millisecond)
  }
}
