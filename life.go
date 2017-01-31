package main

import (
  "fmt"
  "math/rand"
  "time"
  "flag"
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

    for j := 0; j < grid.width; j++ {
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

func (grid *Grid) GetCell(x, y int) bool{
  return grid.cell[y][x]
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
      if !(i == 0 && j == 0) && grid.Next(x+i, y+j) && grid.GetCell(x+i, y+j) {
        count++
      }
    }
  }


  return count
}

/*
  This is where the magic happens. Grid is iterated over, and
  game logic is applied to each cell to return a new grid.
*/
func NextGeneration(grid *Grid) {
  for x:=0; x < grid.width; x++ {
    for y:=0; y < grid.height; y++ {
      // Get # of alive neighbors.
      count := grid.GetAliveNeighbors(x, y)
      if grid.GetCell(x,y) && ( count == 3 || count == 2) {
        grid.cell[y][x] = true
      } else if grid.GetCell(x,y) == false && count == 3 {
        grid.cell[y][x] = true
      } else if grid.GetCell(x,y) && (count > 3 || count < 2) {
        grid.cell[y][x] = false
      } else {
        grid.cell[y][x] = false
      }
    }
  }
}



func main(){
  // Allows users to set custom parameters for the game of life.
  var width, height, cycles, speed int
  flag.IntVar(&width, "width", 30, "grid width")
  flag.IntVar(&height, "height", 30, "grid height")
  flag.IntVar(&cycles, "cycles", 15, "number of cycles")
  flag.IntVar(&speed, "speed", 500, "how fast it be")
  flag.Parse()

  // Run the game based on user parameters/defaults.
  grid := NewGrid(width, height)
  InitGrid(grid)
  for i:=0; i < cycles; i++ {
    fmt.Println(PrintGrid(grid))
    NextGeneration(grid)
    time.Sleep(time.Duration(speed) * time.Millisecond)
  }
}
