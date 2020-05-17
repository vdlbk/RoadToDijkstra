package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strings"
)

const (
	Rock        = "ROCK"
	Paper       = "PAPER"
	Scissors    = "SCISSORS"
	Wall        = "#"
	EmptyCell   = " "
	Pellet      = "*"
	SuperPellet = "@"
	Player      = "P"
	Enemy       = "E"
)

type Pac struct {
	ID                int
	X                 int
	Y                 int
	Type              string
	Dead              bool
	ActiveRoundsCount int
	ImmobileSince     int
	SpeedTurnsLeft    int
	AbilityCooldown   int
	Memory            []string
	Sight             Cells
}

type PacsManager struct {
	Round          int
	Width          int
	Height         int
	Grid           [][]string
	PacIDs         []int
	Pacs           map[int]*Pac
	Enemies        map[int]*Pac
	PelletsInSight map[int]Cells
	EnemiesInSight map[int]map[*Pac]*Cell // map[pacID][enemy]*Cell
	PacDestination Cells
}

func NewPacsManager(width, height int) PacsManager {
	return PacsManager{
		Round:          -1,
		Width:          width,
		Height:         height,
		Grid:           make([][]string, height),
		Pacs:           make(map[int]*Pac),
		Enemies:        make(map[int]*Pac),
		PelletsInSight: make(map[int]Cells),
		EnemiesInSight: make(map[int]map[*Pac]*Cell),
		PacDestination: []*Cell{},
	}
}

func (p *PacsManager) String() string {
	return fmt.Sprintf("(PM %v)", p.Pacs)
}

func (p *PacsManager) NewRound() {
	p.Round++
}

func (p *PacsManager) ManageGrid(i int, line string) {
	p.Grid[i] = strings.Split(line, "")
}

func (p *PacsManager) IsWall(x, y int) bool {
	if y >= 0 && y < len(p.Grid) && x >= 0 && x < len(p.Grid[y]) {
		return p.Grid[y][x] == Wall
	}
	return true
}

func (p *PacsManager) PrintGrid() string {
	var grid []string
	for _, row := range p.Grid {
		grid = append(grid, strings.Join(row, ""))
	}
	return strings.Join(grid, "\n")
}

func (p *PacsManager) ManagePac(mine bool, pacID, x, y int, pacType string, speedTurnsLeft, abilityCooldown int) {
	if mine {
		if pac, ok := p.Pacs[pacID]; ok {
			pac.Update(x, y, pacType, speedTurnsLeft, abilityCooldown, p.Round)
		} else {
			p.Pacs[pacID] = NewPac(pacID, x, y, pacType, speedTurnsLeft, abilityCooldown)
			p.PelletsInSight[pacID] = []*Cell{}
		}
		p.Pacs[pacID].InitSight(p)
		p.PacIDs = append(p.PacIDs, pacID)
	} else {
		if pac, ok := p.Enemies[pacID]; ok {
			pac.Update(x, y, pacType, speedTurnsLeft, abilityCooldown, p.Round)
		} else {
			p.Enemies[pacID] = NewPac(pacID, x, y, pacType, speedTurnsLeft, abilityCooldown)
			p.Enemies[pacID].ActiveRoundsCount = p.Round
		}
		p.Enemies[pacID].InitSight(p)
	}
}

func (p *PacsManager) RegisterPellet(i, j, value int) {
	for pacID := range p.Pacs {
		cells := p.PelletsInSight[pacID]
		cells.Upsert(&Cell{X: i, Y: j, Value: value, CurrentlyInSight: true})
		p.PelletsInSight[pacID] = cells
	}
	pellet := Pellet
	if value == 10 {
		pellet = SuperPellet
	}
	p.Grid[j][i] = pellet
}

func (p *PacsManager) PrintPellets() {
	for k, v := range p.PelletsInSight {
		debug(k, v)
	}
}

func (p *PacsManager) NotifyPacs() {
	// Remove dead pac
	for pacID, pac := range p.Pacs {
		if pac.ActiveRoundsCount != p.Round {
			debug(pac.ID, "is dead", pac.ActiveRoundsCount, p.Round)
			delete(p.Pacs, pacID)
			delete(p.PelletsInSight, pacID)
			delete(p.EnemiesInSight, pacID)
		}
	}

	for _, pac := range p.Pacs {
		// Init map of enemy for each pac
		if _, ok := p.EnemiesInSight[pac.ID]; !ok {
			p.EnemiesInSight[pac.ID] = make(map[*Pac]*Cell)
		}

		cells := p.PelletsInSight[pac.ID]
		for _, pac := range p.Pacs {
			cells.Remove(pac.X, pac.Y)
		}

		for _, enemy := range p.Enemies {
			cell := pac.SeeEnemy(p.EnemiesInSight, enemy)
			p.EnemiesInSight[pac.ID][enemy] = cell
			cells.Remove(enemy.X, enemy.Y)
		}

		cells.ResetSight()
		p.PelletsInSight[pac.ID] = cells
	}
}

func (p *PacsManager) CleanPellet() {
	cellsToRemove := Cells{}
	for pacID, pac := range p.Pacs {
		for _, cellInSight := range pac.Sight {
			cells := p.PelletsInSight[pacID]
			if cell, _ := cells.Contains(cellInSight); cell != nil && cell.CurrentlyInSight == false {
				cellsToRemove = append(cellsToRemove, cell)
			}
		}
	}
	for pacID, _ := range p.Pacs {
		cells := p.PelletsInSight[pacID]
		for _, cellToRemove := range cellsToRemove {
			cells.Remove(cellToRemove.X, cellToRemove.Y)
		}
		p.PelletsInSight[pacID] = cells
	}
}

func (p *PacsManager) Decide() {
	p.CleanPellet()
	closestCells := make(map[int]*Cell)
	for pacID, pac := range p.Pacs {
		cells := p.PelletsInSight[pacID]
		cells.Calibrate(pac.X, pac.Y)
		sort.Sort(cells)
		p.PelletsInSight[pacID] = cells

		// debug(pacID, "closest are:", cells[:3])

		if cell := p.FindClosestCell(pacID, 0); cell != nil {
			closestCells[pacID] = cell
		}
	}
	closestCells = p.CheckSingleClosestCells(closestCells, 0)
	var action []string

	for pacID, pac := range p.Pacs {

		// Is there something to do with the enemies (go away, chase them, ... ?)
		if s, ok := pac.CheckEnemies(p); ok {
			delete(closestCells, pacID)
			action = append(action, s)
			continue
		} else {
			// Nothing interesting to do with enemies, focus on pellets

			// TODO: this is an experiment. Can be remove
			if !p.HasEnemyInSight(pac) {
				if s, ok := pac.AskForSpeed(); ok {
					delete(closestCells, pacID)
					action = append(action, s)
					continue
				}
			}

			// If none pellets are in sight, change strategy
			if len(p.PelletsInSight[pacID]) == 0 {
				debug("go to random position")
				// TODO
			}

			// if pac is stuck for some turn
			immobileLimit := 2
			if pac.SpeedTurnsLeft > 0 {
				immobileLimit = 1
			}
			if pac.ImmobileSince > immobileLimit {
				delete(closestCells, pacID)
				// TODO
				if cell := p.FindFarthestCell(pac.ID); cell != nil {
					action = append(action, pac.Move(cell.X, cell.Y))
				}
			}
		}
	}
	// Default strategy : go to closest pellets with highest value
	for id, cell := range closestCells {
		if cell != nil {
			action = append(action, p.Pacs[id].Move(cell.X, cell.Y))
		} else {
			debug("NO CLOSEST CELL")
			action = append(action, p.Pacs[id].Move(p.GenerateRandomCell()))
		}
	}

	fmt.Printf(strings.Join(action, "|") + "\n")
	p.PacDestination = []*Cell{}
	p.PacIDs = []int{}
}

func (p *PacsManager) CheckSingleClosestCells(closestCells map[int]*Cell, startIdx int) map[int]*Cell {
	for pacID, cell := range closestCells {
		cursor := startIdx
		otherPacIDs := p.GetOtherPacID(pacID)
		for i := 0; i < len(otherPacIDs); {
			otherPacID := otherPacIDs[i]
			otherCell := closestCells[otherPacID]
			if otherCell != nil && otherCell.Equals(cell) && otherCell.Dist <= cell.Dist {
				// otherPacID est plus proche que pacID pour cette cellule
				cursor++
				if c := p.FindClosestCell(pacID, cursor); c != nil {
					cell = c
					i = 0
				}

			} else {
				i++
			}
		}
		closestCells[pacID] = cell
	}

	for pacID, cell := range closestCells {
		for otherPacID, otherCell := range closestCells {
			if pacID != otherPacID && cell != nil && cell.Equals(otherCell) {
				if startIdx > 10 {
					closestCells[pacID] = nil
					break
				} else {
					return p.CheckSingleClosestCells(closestCells, startIdx+1)
				}

			}
		}
	}

	return closestCells
}

func (p *PacsManager) FindClosestCell(pacID int, startIdx int) *Cell {
	for i, cell := range p.PelletsInSight[pacID] {
		if i == startIdx {
			p.PacDestination = append(p.PacDestination, cell)
			return cell
		}
	}
	return nil
}

// FIXME: Not a big deal for now but 2+ can go to the same cell
func (p *PacsManager) FindFarthestCell(pacID int) *Cell {
	for i := len(p.PelletsInSight[pacID]) - 1; i >= 0; i-- {
		cell := p.PelletsInSight[pacID][i]
		if existingCell, _ := p.PacDestination.Contains(cell); existingCell == nil {
			p.PacDestination = append(p.PacDestination, cell)
			return cell
		}
	}
	return nil
}

func (p *PacsManager) FindRandomCell(pacID int) *Cell {
	length := len(p.PelletsInSight[pacID])
	for i := 0; i < length; i++ {
		k := 0
		if length > 1 {
			k = rand.Intn(length - 1)
		}
		cell := p.PelletsInSight[pacID][k]
		if existingCell, _ := p.PacDestination.Contains(cell); existingCell != nil {
			p.PacDestination = append(p.PacDestination, cell)
			return cell
		}
	}
	return nil
}

func (p *PacsManager) GenerateRandomCell() (int, int) {
	x := rand.Intn(p.Width) - 1
	if x < 0 {
		x = 0
	}
	y := rand.Intn(p.Height) - 1
	if y < 0 {
		y = 0
	}
	return x, y
}

func (p *PacsManager) GenerateCellFarFrom(distMin float64, pac *Pac) (int, int) {
	x, y := p.GenerateRandomCell()
	dist := computeDistance(convert(pac.X, pac.Y, x, y))
	for dist < distMin {
		x, y = p.GenerateRandomCell()
		dist = computeDistance(convert(pac.X, pac.Y, x, y))
	}
	return x, y
}

func (p *PacsManager) GetOtherPacID(pacID int) []int {
	var otherPacIDs []int
	for _, id := range p.PacIDs {
		if id != pacID {
			otherPacIDs = append(otherPacIDs, id)
		}
	}
	return otherPacIDs
}

func (p *PacsManager) GetCellToAvoid(pac *Pac) *Cell {
	for _, pID := range p.GetOtherPacID(pac.ID) {
		otherPac := p.Pacs[pID]
		for _, cell := range pac.Sight {
			if cell.X == otherPac.X && cell.Y == otherPac.Y {
				return cell
			}
		}
	}

	if len(pac.Sight) == 0 {
		return nil
	}

	return pac.Sight[0]
}

func (p *PacsManager) HasEnemyInSight(pac *Pac) bool {
	for _, cell := range pac.Sight {
		for _, enemy := range p.Enemies {
			if cell.X == enemy.X && cell.Y == enemy.Y {
				return true
			}
		}
	}
	return false
}

func NewPac(ID, x, y int, pacType string, left, cooldown int) *Pac {
	p := Pac{
		ID:                ID,
		X:                 x,
		Y:                 y,
		Dead:              false,
		ActiveRoundsCount: 0,
		Type:              pacType,
		SpeedTurnsLeft:    left,
		AbilityCooldown:   cooldown,
	}
	return &p
}

func (p *Pac) Update(x, y int, pacType string, left, cooldown, round int) {
	if p.X == x && p.Y == y {
		p.ImmobileSince++
	} else {
		p.ImmobileSince = 0
	}
	p.X = x
	p.Y = y
	//p.Dead = false // Handle death later
	p.ActiveRoundsCount = round
	p.Type = pacType
	p.SpeedTurnsLeft = left
	p.AbilityCooldown = cooldown
}

func (p *Pac) SeePellet(i, j, value int) *Cell {
	dist := computeDistance(convert(p.X, p.Y, i, j))
	return &Cell{X: i, Y: j, Value: value, Dist: dist}
}

func (p *Pac) SeeEnemy(m map[int]map[*Pac]*Cell, enemy *Pac) *Cell {
	dist := computeDistance(convert(p.X, p.Y, enemy.X, enemy.Y))
	value := -100
	if enemy.Type == Rock {
		if p.Type == Paper {
			value = 100
		} else if p.Type == Rock {
			value = 0
		}
	} else if enemy.Type == Paper {
		if p.Type == Scissors {
			value = 100
		} else if p.Type == Paper {
			value = 0
		}
	} else {
		if p.Type == Rock {
			value = 100
		} else if p.Type == Scissors {
			value = 0
		}
	}

	if cell, ok := m[p.ID][enemy]; ok {
		cell.Update(enemy.X, enemy.Y, value, dist)
		return cell
	} else {
		return &Cell{X: enemy.X, Y: enemy.Y, Value: value, Dist: dist}
	}
}

func (p *Pac) Move(x, y int) string {
	return fmt.Sprintf("MOVE %v %v %v", p.ID, x, y)
}

func (p *Pac) AskForSpeed() (string, bool) {
	//debug("AskForSpeed", p.ID, p.AbilityCooldown)
	if p.AbilityCooldown == 0 {
		return fmt.Sprintf("SPEED %v", p.ID), true
	}
	return "", false
}

func (p *Pac) AskForMorph(newType string) (string, bool) {
	debug("AskForMorph", p.ID, newType)
	if p.AbilityCooldown == 0 && newType != p.Type {
		return fmt.Sprintf("SWITCH %v %s", p.ID, newType), true
	}
	return "", false
}

func (p *Pac) CheckEnemies(pacManager *PacsManager) (string, bool) {
	for enemy, cell := range pacManager.EnemiesInSight[p.ID] {
		var distTrigger float64 = 3
		//if enemy.SpeedTurnsLeft > 0 && p.SpeedTurnsLeft == 0{
		//	distTrigger = 1
		//}
		// Skip this enemy if its been a while since we saw him
		if _, idx := p.Sight.Contains(&Cell{X: enemy.X, Y: enemy.Y}); idx == -1 && (p.ActiveRoundsCount-enemy.ActiveRoundsCount) > 2 {
			debug(p.ID, "ignore check enemy", enemy, cell)
			continue
		} else if idx != -1 && p.ActiveRoundsCount != enemy.ActiveRoundsCount {
			debug(p.ID, "ignore check enemy is in sight but its not", p.Sight, idx, enemy, cell)
			continue
		}
		if cell.Dist > distTrigger {
			debug(p.ID, "ignore, too far", cell.Dist, distTrigger)
			continue
		}
		if cell.Trend < 0 && cell.Value < 0 {
			// Danger: an enemy is coming and can eat me. Try to morph into his weak spot
			debug(p.ID, "danger", enemy, cell)
			if s, ok := p.AskForMorph(weakSpot(enemy.Type)); ok {
				return s, ok
			} else {
				debug(p.ID, "failed to morph in danger situation", ok, enemy, cell)
				// TODO: an enemy is coming and we fail switch. Try to go add the opposite
				if s := p.AvoidCell(enemy.X, enemy.Y); len(s) > 0 {
					return s, true
				}
				//if cell := pacManager.FindFarthestCell(p.ID); cell != nil {
				//	return p.Move(cell.X, cell.Y), true
				//}
			}
		} else if cell.Trend < 0 && cell.Value > 0 {
			// Go for it. You can eat him
			debug(p.ID, "go for it", enemy, cell)
			if enemy.AbilityCooldown >= 3 && enemy.SpeedTurnsLeft == 0 {
				debug("chasing food", enemy.AbilityCooldown)
				return p.Move(enemy.X, enemy.Y), true
			} else if s := p.AvoidCell(enemy.X, enemy.Y); enemy.AbilityCooldown <= 1 && len(s) > 0 {
				debug("stop chasing food", enemy.AbilityCooldown)
				return s, true
			}
		} else if cell.Trend < 0 && cell.Value == 0 {
			// Cannot hurt me because we are in the same type
			debug(p.ID, "cannot hurt", enemy, cell)
			if s, ok := p.AskForMorph(weakSpot(enemy.Type)); ok && enemy.AbilityCooldown > 2 {
				return s, ok
			} else {
				debug(p.ID, "failed to morph in safe situation", ok, enemy, cell)
				if s := p.AvoidCell(enemy.X, enemy.Y); enemy.AbilityCooldown <= 1 && len(s) > 0 {
					return s, true
				}
			}
		} else if cell.Trend > 0 && cell.Value < 0 {
			// A danger is leaving
			debug(p.ID, "danger leaving", enemy, cell)
			if s := p.AvoidCell(enemy.X, enemy.Y); enemy.AbilityCooldown != 0 && len(s) > 0 {
				return s, true
			}
		} else if cell.Trend > 0 && cell.Value > 0 {
			// An opponent that I can eat is leaving. Maybe chase him
			debug(p.ID, "food leaving", enemy, cell)
			// FIXME
			//if s, ok := p.AskForSpeed(); ok && cell.Dist <= 2 {
			//	return s, ok
			//}
			//return p.Move(enemy.X, enemy.Y), true
		} else if cell.Trend > 0 && cell.Value == 0 {
			// An opponent that cannot hurt me if leaving. Ignore him.
			debug(p.ID, "enemy of same type is leaving", enemy, cell)
		} else if cell.Trend == 0 && cell.Value < 0 {
			debug(p.ID, "following enemy", enemy, cell)
			// I'm following an enemy that can eat me. Consider to go away from him
			if s, ok := p.AskForMorph(weakSpot(enemy.Type)); ok {
				return s, ok
			} else {
				debug(p.ID, "failed to morph in bad situation", ok, enemy, cell)
				// TODO: an enemy is coming and we fail switch. Try to go add the opposite
				if s := p.AvoidCell(enemy.X, enemy.Y); len(s) > 0 {
					return s, true
				}
				//if cell := pacManager.FindFarthestCell(p.ID); cell != nil {
				//	return p.Move(cell.X, cell.Y), true
				//}
			}
		} else if cell.Trend == 0 && cell.Value > 0 {
			// I'm following an enemy I can eat. Maybe chase him
			debug(p.ID, "following food", enemy, cell)
			if enemy.AbilityCooldown >= 3 {
				debug("chasing food", enemy.AbilityCooldown)
				return p.Move(enemy.X, enemy.Y), true
			} else if s := p.AvoidCell(enemy.X, enemy.Y); enemy.AbilityCooldown <= 1 && len(s) > 0 {
				debug("stop chasing food", enemy.AbilityCooldown)
				return s, true
			}
		} else if cell.Trend == 0 && cell.Value == 0 {
			// I'm following an enemy that does matter. Ignore him
			debug(p.ID, "following swiss", enemy, cell)
			if s, ok := p.AskForMorph(weakSpot(enemy.Type)); ok && enemy.AbilityCooldown > 2 {
				return s, ok
			} else {
				debug(p.ID, "failed to morph in neutral situation", ok, enemy, cell)
				if s := p.AvoidCell(enemy.X, enemy.Y); enemy.AbilityCooldown <= 1 && len(s) > 0 {
					return s, true
				}
			}
		} else {
			debug(p.ID, "is in an unkown situation face to", enemy, cell)
		}
	}
	return "", false
}

func (p *Pac) InitSight(manager *PacsManager) {
	p.Sight = Cells{}
	// North
	for i := p.Y - 1; ; i-- {
		if i < 0 {
			i = manager.Height - 1
		}
		if manager.IsWall(p.X, i) || i == p.Y {
			break
		}
		p.Sight.Upsert(&Cell{X: p.X, Y: i})
	}

	//South
	for i := p.Y + 1; ; i++ {
		if i >= manager.Height {
			i = 0
		}
		if manager.IsWall(p.X, i) || i == p.Y {
			break
		}
		p.Sight.Upsert(&Cell{X: p.X, Y: i})
	}

	// West
	for i := p.X - 1; ; i-- {
		if i < 0 {
			i = manager.Width - 1
		}
		if manager.IsWall(i, p.Y) || i == p.X {
			break
		}
		p.Sight.Upsert(&Cell{X: i, Y: p.Y})
	}

	// East
	for i := p.X + 1; ; i++ {
		if i >= manager.Width {
			i = 0
		}
		if manager.IsWall(i, p.Y) || i == p.X {
			break
		}
		p.Sight.Upsert(&Cell{X: i, Y: p.Y})
	}
	sort.Sort(p.Sight)
}

func (p *Pac) AvoidCell(enemyX, enemyY int) string {
	if p.X == enemyX {
		if p.Y < enemyY {
			// Je dois monter (baisser mon Y // tout sauf augmenter mon Y)
			for _, cell := range p.Sight {
				if p.Y < cell.Y {
					continue
				}
				return p.Move(cell.X, cell.Y)
			}
		} else {
			// Je dois descendre (augmenter mon Y // tauf sauf baisser mon Y)
			for _, cell := range p.Sight {
				if p.Y > cell.Y {
					continue
				}
				return p.Move(cell.X, cell.Y)
			}
		}
	} else if p.Y == enemyY {
		if p.X < enemyX {
			// Je dois aller à gauche (baisser mon X // tout sauf augmenter X)
			for _, cell := range p.Sight {
				if p.X < cell.X {
					continue
				}
				return p.Move(cell.X, cell.Y)
			}
		} else {
			// Je dois aller à droite (augmenter mon X // tout sauf baisser mon X)
			for _, cell := range p.Sight {
				if p.X > cell.X {
					continue
				}
				return p.Move(cell.X, cell.Y)
			}
		}
	}
	return ""
}

func (p *Pac) String() string {
	return fmt.Sprintf("PAC(%v/%v)(%+v;%v)(%v %v)", p.Type, p.ActiveRoundsCount, p.X, p.Y, p.SpeedTurnsLeft, p.AbilityCooldown)
}

type Cell struct {
	X                int
	Y                int
	Value            int
	Dist             float64
	Trend            float64
	CurrentlyInSight bool
}

type Cells []*Cell

func (c *Cell) String() string {
	return fmt.Sprintf("C(%v %v %v %.3f %.2f)", c.X, c.Y, c.Value, c.Dist, c.Trend)
}

func (c *Cell) Equals(other *Cell) bool {
	return c.X == other.X && c.Y == other.Y
}

func (c *Cell) Update(x, y, value int, dist float64) {
	c.X = x
	c.Y = y
	c.Value = value
	c.UpdateDist(dist)
}

func (c *Cell) UpdateDist(dist float64) {
	c.Trend = dist - c.Dist
	c.Dist = dist
}

func (c *Cell) SanitizePosition() (int, int) {
	x := c.X
	y := c.Y
	if x--; x < 0 {
		x = 0
	}
	if y--; y < 0 {
		y = 0
	}
	return x, y
}

func (c *Cells) Upsert(value *Cell) {
	values := *c
	if cell, _ := c.Contains(value); cell == nil {
		values = append(values, value)
	} else {
		cell.CurrentlyInSight = value.CurrentlyInSight
	}
	*c = values
}

func (c *Cells) Contains(value *Cell) (*Cell, int) {
	if value == nil {
		return nil, -1
	}
	for i, cell := range *c {
		if cell.Equals(value) {
			return cell, i
		}
	}
	return nil, -1
}

func (c *Cells) Remove(x, y int) {
	if _, idx := c.Contains(&Cell{X: x, Y: y}); idx != -1 {
		values := *c
		values[len(values)-1], values[idx] = values[idx], values[len(values)-1]
		*c = values[:len(values)-1]
	}
}

func (c *Cells) Calibrate(x, y int) {
	for _, cell := range *c {
		dist := computeDistance(convert(x, y, cell.X, cell.Y))
		cell.UpdateDist(dist)
	}
}

func (c *Cells) ResetSight() {
	for _, cell := range *c {
		cell.CurrentlyInSight = false
	}
}

func (c Cells) Len() int      { return len(c) }
func (c Cells) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c Cells) Less(i, j int) bool {
	if c[i].Value > c[j].Value {
		return true
	} else if c[i].Value == c[j].Value {
		if c[i].CurrentlyInSight == c[j].CurrentlyInSight {
			return c[i].Dist < c[j].Dist
		} else if c[i].CurrentlyInSight {
			return c[i].Dist-c[j].Dist < 2
		} else {
			return false
		}
	}
	return false
}

func (c Cells) NormalLess(i, j int) bool {
	return c[i].Dist < c[j].Dist
}

func computeDistance(x, y, i, j float64) float64 {
	return math.Sqrt(math.Pow(i-x, 2) + math.Pow(y-j, 2))
}

func convert(x, y, i, j int) (float64, float64, float64, float64) {
	return float64(x), float64(y), float64(i), float64(j)
}

func weakSpot(pacType string) string {
	switch pacType {
	case Rock:
		return Paper
	case Paper:
		return Scissors
	case Scissors:
		return Rock
	default:
		return pacType
	}
}

func exists(ints []int, value int) bool {
	for _, x := range ints {
		if x == value {
			return true
		}
	}
	return false
}

func debug(values ...interface{}) {
	fmt.Fprintln(os.Stderr, values...)
}
