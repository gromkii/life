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
*/
func NewGrid(width, height) Grid {
  // Define the 2d array to being looping through.
  cell := make([][]bool, height)
  for i:=0 ; i < range cell ; i++ {
    cell[i] = make([]bool, width)
    for j:=0; j < range cell[i]; j++ {
      cell[i][j] = false
    }
    
  }

  return Grid{cell:cell, width:width, height:height}
}

func main(){
  fmt.Println("Hello from Go!")
}
