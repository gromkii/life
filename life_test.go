package test

import (

)

type TestGrid struct {
  cell [][]bool
  width, height int
}

func MakeTestGrid(w, h int) *TestGrid {
  cell := make([][]bool, h)
  for i := range cell {
    cell[i] = make([]bool, w)
    for j := range cell[i] {
      cell[i][j] = false
    }
  }

  return &TestGrid{cell:cell, width: w, height:h}
}


func NeighborTest() *TestGrid{
  grid := MakeTestGrid(3, 3)

  // Make a control grid, use for all tests.
  grid.cell[0][0] = false
  grid.cell[0][1] = true
  grid.cell[0][2] = true
  grid.cell[1][0] = false
  grid.cell[1][1] = false
  grid.cell[1][2] = false
  grid.cell[2][0] = true
  grid.cell[2][1] = true
  grid.cell[2][2] = false

  return grid

}
