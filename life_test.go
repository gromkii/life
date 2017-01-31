package main

import (
  "testing"
  "reflect"
)

func TestTesting(t *testing.T) {
  t.Log("should establish tests")
}

// Initialize a new grid, all cells should be false
func TestNewGrid(t *testing.T) {
  grid := NewGrid(3, 3)
  for i:=0; i < grid.height; i++ {
    for j:=0; j < grid.width; j++ {
      if grid.cell[i][j] == true {
        t.Error("Except new cell to init to false, was true")
        t.Fail()
      }
    }
  }
}

func TestPrintGrid(t *testing.T) {
  grid := NewGrid(3, 3)
  print := PrintGrid(grid)

  if (reflect.TypeOf(print)) != reflect.TypeOf("string") {
    t.Error("Expect PringGrid to return a string")
  }
}

func TestNext(t *testing.T) {
  x, y := -1, -1
  a, b := 3, 3
  grid := NewGrid(3, 3)
  caseA := grid.Next(x, y)
  caseB := grid.Next(a, b)
  caseC := grid.Next(1, 1)

  if caseA != false {
    t.Error("Expect false if index is out of range.")
  }

  if caseB != false {
    t.Error("Expect false if index is out of range.")
  }

  if caseC != true {
    t.Error("Expect true if index is in range.")
  }
}

func TestNeighborCount(t *testing.T) {
  grid := NewGrid(3,3)
  grid.cell[0][1] = true
  grid.cell[1][0] = true

  count := grid.GetAliveNeighbors(0,0)

  if count != 2 {
    t.Error("Expect alive neightbors to be true")
  }
}

func TestNextGeneration(t *testing.T) {
  gridA := NewGrid(3,3)
  gridA.cell[0][0] = true
  gridA.cell[0][1] = true
  gridA.cell[1][0] = true

  gridB := NewGrid(3,3)
  gridB.cell[0][0] = true
  gridB.cell[0][1] = false
  gridB.cell[1][0] = true
  gridB.cell[1][1] = true

  gridC := NewGrid(3,3)
  gridC.cell[0][0] = true
  gridC.cell[0][1] = true
  gridC.cell[0][2] = true
  gridC.cell[1][0] = true
  gridC.cell[1][1] = true

  gridD := NewGrid(3,3)
  gridD.cell[0][0] = true
  gridD.cell[0][1] = true


  NextGeneration(gridA)
  NextGeneration(gridB)
  NextGeneration(gridC)
  NextGeneration(gridD)


  if !gridA.cell[0][0] {
    t.Error("Expect alive cell with 2 neighbors to be alive")
  }

  if !gridB.cell[0][1] {
    t.Error("Expect dead cell with exactly 3 alive neighbors to be alive")
  }

  if gridC.cell[0][1] {
    t.Error("Expect alive cell with > 3 neighbors to be dead")
  }

  if gridD.cell[0][0] {
    t.Error("Expect alive cells with < 2 neightbos to be dead")
  }


}
