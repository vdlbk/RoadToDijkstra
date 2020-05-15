package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"testing"
)

func TestPlayground(t *testing.T) {
	fmt.Println(rand.Intn(0))
}

func TestSortCells(t *testing.T) {
	Point := struct {
		X int
		Y int
	}{
		X: 2,
		Y: 3,
	}
	slice := Cells{
		{X: 0, Y: 1, Value: 1, Dist: 0, CurrentlyInSight: true},
		{X: 4, Y: 8, Value: 1, Dist: 0, CurrentlyInSight: true},
		{X: 2, Y: 1, Value: 1, Dist: 0, CurrentlyInSight: true},
		{X: 1, Y: 0, Value: 1, Dist: 0, CurrentlyInSight: true},
		{X: 3, Y: 3, Value: 1, Dist: 0, CurrentlyInSight: true},
		{X: 4, Y: 4, Value: 10, Dist: 0, CurrentlyInSight: true},
		{X: 6, Y: 2, Value: 1, Dist: 0, CurrentlyInSight: true},
		{X: 5, Y: 9, Value: 10, Dist: 0, CurrentlyInSight: true},
	}

	slice.Calibrate(Point.X, Point.Y)

	fmt.Println(slice)

	sort.Sort(slice)
	fmt.Println(slice)

	expectedSlice := Cells{
		{X: 4, Y: 4, Value: 10, Dist: 0, CurrentlyInSight: true},
		{X: 5, Y: 9, Value: 10, Dist: 0, CurrentlyInSight: true},
		{X: 3, Y: 3, Value: 1, Dist: 0, CurrentlyInSight: true},
		{X: 2, Y: 1, Value: 1, Dist: 0, CurrentlyInSight: true},
		{X: 0, Y: 1, Value: 1, Dist: 0, CurrentlyInSight: true},
		{X: 1, Y: 0, Value: 1, Dist: 0, CurrentlyInSight: true},
		{X: 6, Y: 2, Value: 1, Dist: 0, CurrentlyInSight: true},
		{X: 4, Y: 8, Value: 1, Dist: 0, CurrentlyInSight: true},
	}
	expectedSlice.Calibrate(Point.X, Point.Y)
	fmt.Println(expectedSlice)
}

func TestSortCells_WithSight(t *testing.T) {
	Point := struct {
		X int
		Y int
	}{
		X: 2,
		Y: 3,
	}
	slice := Cells{
		{X: 0, Y: 1, Value: 1, Dist: 0, CurrentlyInSight: false},
		{X: 4, Y: 8, Value: 1, Dist: 0, CurrentlyInSight: true},
		{X: 2, Y: 1, Value: 1, Dist: 0, CurrentlyInSight: true},
		{X: 1, Y: 0, Value: 1, Dist: 0, CurrentlyInSight: true},
		{X: 3, Y: 3, Value: 1, Dist: 0, CurrentlyInSight: true},
		{X: 4, Y: 4, Value: 10, Dist: 0, CurrentlyInSight: true},
		{X: 6, Y: 2, Value: 1, Dist: 0, CurrentlyInSight: true},
		{X: 5, Y: 9, Value: 10, Dist: 0, CurrentlyInSight: true},
	}

	slice.Calibrate(Point.X, Point.Y)

	fmt.Println(slice)

	sort.Sort(slice)
	fmt.Println(slice)

	expectedSlice := Cells{
		{X: 4, Y: 4, Value: 10, Dist: 0, CurrentlyInSight: true},
		{X: 5, Y: 9, Value: 10, Dist: 0, CurrentlyInSight: true},
		{X: 3, Y: 3, Value: 1, Dist: 0, CurrentlyInSight: true},
		{X: 2, Y: 1, Value: 1, Dist: 0, CurrentlyInSight: true},
		{X: 1, Y: 0, Value: 1, Dist: 0, CurrentlyInSight: true},
		{X: 0, Y: 1, Value: 1, Dist: 0, CurrentlyInSight: false},
		{X: 6, Y: 2, Value: 1, Dist: 0, CurrentlyInSight: true},
		{X: 4, Y: 8, Value: 1, Dist: 0, CurrentlyInSight: true},
	}
	expectedSlice.Calibrate(Point.X, Point.Y)
	fmt.Println(expectedSlice)
}

func TestCells_Upsert(t *testing.T) {
	slice := Cells{}
	slice.Upsert(&Cell{X: 0, Y: 1})
	if len(slice) != 1 {
		t.Error("Unexpected length. Want 1")
	}
	slice.Upsert(&Cell{X: 1, Y: 3})
	if len(slice) != 2 {
		t.Error("Unexpected length. Want 2")
	}
	slice.Upsert(&Cell{X: 1, Y: 3})
	if len(slice) != 2 {
		t.Error("Unexpected length. Want 2")
	}
}

func TestCells_Calibrate(t *testing.T) {
	slice := Cells{
		{X: 0, Y: 1},
		{X: 2, Y: 2},
		{X: 3, Y: 1},
		{X: 5, Y: 9},
	}

	slice.Calibrate(4, 4)
	for _, cell := range slice {
		if cell.Dist == 0 {
			t.Error("Unexpected dist. Want > 0")
		}
	}
}

func TestCells_Remove(t *testing.T) {
	slice := Cells{
		{X: 0, Y: 1},
		{X: 2, Y: 2},
		{X: 3, Y: 1},
		{X: 5, Y: 9},
	}

	slice.Remove(2, 2)
	if len(slice) != 3 {
		t.Error("Unexpected length. Want 3")
	}
}

func TestManageGrid(t *testing.T) {
	pacManager := NewPacsManager(33, 13)
	pacManager.ManageGrid(0, "#################################")
	pacManager.ManageGrid(1, "#     # #   #       #   # #     #")
	pacManager.ManageGrid(2, "# # ### ### # # # # # ### ### # #")
	pacManager.ManageGrid(3, "# #   #       #   #       #   # #")
	pacManager.ManageGrid(4, "# ### # ##### # # # ##### # ### #")
	pacManager.ManageGrid(5, "#         #   #   #   #         #")
	pacManager.ManageGrid(6, "# # ##### # # ##### # # ##### # #")
	pacManager.ManageGrid(7, "# #   #     #       #     #   # #")
	pacManager.ManageGrid(8, "### # # # # ######### # # # # ###")
	pacManager.ManageGrid(9, "# #     # #           # #     # #")
	pacManager.ManageGrid(10, "# # ##### ##### # ##### ##### # #")
	pacManager.ManageGrid(11, "#               #               #")
	pacManager.ManageGrid(12, "#################################")
	fmt.Print(pacManager.PrintGrid())
	pacManager.RegisterPellet(1, 1, 10)
	pacManager.RegisterPellet(1, 4, 1)
	fmt.Print(pacManager.PrintGrid())
}

func TestSameDist(t *testing.T) {
	width, height := 34, 15
	pacsManager := NewPacsManager(width, height)
	for i := 0; i < height; i++ {
		pacsManager.ManageGrid(i, strings.Repeat(" ", width))
	}

	pacsManager.NewRound()
	pacsManager.ManagePac(true, 0, 21, 10, Rock, 0, 0)
	pacsManager.ManagePac(true, 1, 3, 4, Paper, 0, 0)
	pacsManager.ManagePac(true, 2, 14, 11, Scissors, 0, 0)
	pacsManager.ManagePac(true, 3, 10, 3, Rock, 0, 0)
	fmt.Println(pacsManager.String())
	pacsManager.NotifyPacs()

	initTestPellet(&pacsManager, 0)

	fmt.Println(pacsManager.String())

	pacsManager.Decide()

	// New turn
	pacsManager.NewRound()
	pacsManager.ManagePac(true, 0, 21, 11, Rock, 5, 10)
	pacsManager.ManagePac(true, 1, 3, 5, Paper, 5, 10)
	pacsManager.ManagePac(true, 2, 15, 11, Scissors, 5, 10)
	pacsManager.ManagePac(true, 3, 9, 3, Rock, 5, 10)
	pacsManager.NotifyPacs()

	for i := 0; i < 4; i++ {
		cells := pacsManager.PelletsInSight[i]
		if _, idx := cells.Contains(&Cell{X: 21, Y: 11}); idx != -1 {
			t.Error("cells was not deleted", i, idx)
		}
		if _, idx := cells.Contains(&Cell{X: 3, Y: 5}); idx != -1 {
			t.Error("cells was not deleted", i, idx)
		}
		if _, idx := cells.Contains(&Cell{X: 15, Y: 11}); idx != -1 {
			t.Error("cells was not deleted", i, idx)
		}
		if _, idx := cells.Contains(&Cell{X: 9, Y: 3}); idx != -1 {
			t.Error("cells was not deleted", i, idx)
		}
	}

	initTestPellet(&pacsManager, 1)
	pacsManager.Decide()

	// New turn
	pacsManager.NewRound()
	pacsManager.ManagePac(true, 0, 20, 11, Rock, 4, 9)
	pacsManager.ManagePac(true, 1, 3, 6, Paper, 4, 9)
	pacsManager.ManagePac(true, 2, 16, 11, Scissors, 4, 9)
	pacsManager.ManagePac(true, 3, 8, 3, Rock, 4, 9)
	pacsManager.NotifyPacs()

	for i := 0; i < 4; i++ {
		cells := pacsManager.PelletsInSight[i]
		if _, idx := cells.Contains(&Cell{X: 20, Y: 11}); idx != -1 {
			t.Error("cells was not deleted", i, idx)
		}
		if _, idx := cells.Contains(&Cell{X: 3, Y: 6}); idx != -1 {
			t.Error("cells was not deleted", i, idx)
		}
		if _, idx := cells.Contains(&Cell{X: 16, Y: 11}); idx != -1 {
			t.Error("cells was not deleted", i, idx)
		}
		if _, idx := cells.Contains(&Cell{X: 8, Y: 3}); idx != -1 {
			t.Error("cells was not deleted", i, idx)
		}
	}

	initTestPellet(&pacsManager, 2)
	pacsManager.Decide()

	// New turn
	pacsManager.NewRound()
	pacsManager.ManagePac(true, 0, 19, 11, Rock, 3, 8)
	pacsManager.ManagePac(true, 1, 3, 7, Paper, 3, 8)
	pacsManager.ManagePac(true, 2, 17, 11, Scissors, 3, 8)
	pacsManager.ManagePac(true, 3, 7, 3, Rock, 3, 8)
	pacsManager.NotifyPacs()

	for i := 0; i < 4; i++ {
		cells := pacsManager.PelletsInSight[i]
		if _, idx := cells.Contains(&Cell{X: 19, Y: 11}); idx != -1 {
			t.Error("cells was not deleted", i, idx)
		}
		if _, idx := cells.Contains(&Cell{X: 3, Y: 7}); idx != -1 {
			t.Error("cells was not deleted", i, idx)
		}
		if _, idx := cells.Contains(&Cell{X: 17, Y: 11}); idx != -1 {
			t.Error("cells was not deleted", i, idx)
		}
		if _, idx := cells.Contains(&Cell{X: 7, Y: 3}); idx != -1 {
			t.Error("cells was not deleted", i, idx)
		}
	}

	initTestPellet(&pacsManager, 3)
	pacsManager.Decide()
}

func initTestPellet(pacsManager *PacsManager, super int) {
	if super == 0 {
		pacsManager.RegisterPellet(15, 11, 1)
		pacsManager.RegisterPellet(3, 5, 1)
		pacsManager.RegisterPellet(21, 11, 1)
		pacsManager.RegisterPellet(9, 3, 1)
	}

	if super <= 1 {
		pacsManager.RegisterPellet(20, 11, 1)
		pacsManager.RegisterPellet(3, 6, 1)
		pacsManager.RegisterPellet(16, 11, 10)
		pacsManager.RegisterPellet(8, 3, 1)
	}

	if super <= 2 {
		pacsManager.RegisterPellet(19, 11, 1)
		pacsManager.RegisterPellet(3, 7, 1)
		pacsManager.RegisterPellet(17, 11, 1)
		pacsManager.RegisterPellet(7, 3, 1)
	}

	if super <= 3 {

	}

	pacsManager.RegisterPellet(18, 11, 10)
	pacsManager.RegisterPellet(20, 11, 1)
	pacsManager.RegisterPellet(21, 9, 1)
	pacsManager.RegisterPellet(21, 9, 1)
	pacsManager.RegisterPellet(22, 9, 1)
	pacsManager.RegisterPellet(23, 9, 1)
	pacsManager.RegisterPellet(23, 10, 1)
	pacsManager.RegisterPellet(23, 11, 1)
	pacsManager.RegisterPellet(24, 11, 1)
	pacsManager.RegisterPellet(25, 11, 1)
	pacsManager.RegisterPellet(25, 12, 1)
	pacsManager.RegisterPellet(23, 12, 1)
	pacsManager.RegisterPellet(23, 13, 1)
	pacsManager.RegisterPellet(24, 13, 1)
	pacsManager.RegisterPellet(25, 13, 1)
	pacsManager.RegisterPellet(26, 13, 1)
	pacsManager.RegisterPellet(27, 13, 1)
	pacsManager.RegisterPellet(27, 12, 1)
	pacsManager.RegisterPellet(27, 11, 10)
	pacsManager.RegisterPellet(17, 10, 1)
	pacsManager.RegisterPellet(17, 9, 1)
	pacsManager.RegisterPellet(13, 11, 1)
	pacsManager.RegisterPellet(13, 10, 1)
	pacsManager.RegisterPellet(13, 12, 1)

	pacsManager.RegisterPellet(3, 8, 1)
	pacsManager.RegisterPellet(3, 9, 1)
	pacsManager.RegisterPellet(3, 10, 1)
	pacsManager.RegisterPellet(4, 9, 1)
	pacsManager.RegisterPellet(5, 9, 1)
	pacsManager.RegisterPellet(5, 10, 1)
	pacsManager.RegisterPellet(5, 11, 1)
	pacsManager.RegisterPellet(6, 11, 1)
	pacsManager.RegisterPellet(7, 11, 10)
	pacsManager.RegisterPellet(10, 3, 1)
	pacsManager.RegisterPellet(11, 4, 1)
	pacsManager.RegisterPellet(11, 7, 1)
}

func TestPac_InitSight(t *testing.T) {
	pacManager := NewPacsManager(33, 13)
	pacManager.ManageGrid(0, "############################## ##")
	pacManager.ManageGrid(1, "#     # #   #       #   # #     #")
	pacManager.ManageGrid(2, "# # ### ### # # # # # ### ###   #")
	pacManager.ManageGrid(3, "# #   #       #   #       #     #")
	pacManager.ManageGrid(4, "# ### # ##### # # # ##### # ##  #")
	pacManager.ManageGrid(5, "              #   #              ")
	pacManager.ManageGrid(6, "# # ##### # # ##### # # #####   #")
	pacManager.ManageGrid(7, "# #   #     #       #     #     #")
	pacManager.ManageGrid(8, "### # # # # ######### # # # #  ##")
	pacManager.ManageGrid(9, "        # #           # #     *  ")
	pacManager.ManageGrid(10, "# # ##### ##### # ##### ######  #")
	pacManager.ManageGrid(11, "#               #               #")
	pacManager.ManageGrid(12, "############################## ##")

	pacManager.ManagePac(true, 0, 1, 1, Rock, 0, 0)
	pacManager.ManagePac(true, 1, 13, 5, Rock, 0, 0)
	pacManager.ManagePac(true, 2, 30, 9, Rock, 0, 0)
	pacManager.Pacs[0].InitSight(&pacManager)
	pacManager.Pacs[1].InitSight(&pacManager)
	pacManager.Pacs[2].InitSight(&pacManager)

	if len(pacManager.Pacs[0].Sight) != 10 {
		t.Errorf("Unexpected length. Got: %v; Want: %v", len(pacManager.Pacs[0].Sight), 10)
	}

	if len(pacManager.Pacs[1].Sight) != 33 {
		t.Errorf("Unexpected length. Got: %v; Want: %v", len(pacManager.Pacs[1].Sight), 33)
	}

	// 15 + 12
	if len(pacManager.Pacs[2].Sight) != 27 {
		t.Errorf("Unexpected length. Got: %v; Want: %v", len(pacManager.Pacs[2].Sight), 27)
	}
}
