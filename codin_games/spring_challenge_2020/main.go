package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strings"
)

// [x] Optimiser le fait que c'est au Pac le plus proche d'aller chercher la cell avec la plus de value
// [ ] Il nous faudrait une info pour savoir si l'info et la position connu de l'ennemy est fraiche et si on peut s'y fier
// [ ] Garder une trace des pellets meme si on ne les voit plus, avoir un système de "est-ce à jour ?"
// [ ] Prévoir un buffer d'action pour un PAC ? Genre il faut speed + aller là, ça éviterai de devoir recalculer un truc qu'on sait déjà et se tromper

const (
	Rock     = "ROCK"
	Paper    = "PAPER"
	Scissors = "SCISSORS"
)

type Pac struct {
	ID              int
	X               int
	Y               int
	Type            string
	ImmobileSince   int
	SpeedTurnsLeft  int
	AbilityCooldown int
}

type PacsManager struct {
	Width          int
	Height         int
	PacIDs         []int
	Pacs           map[int]*Pac
	Enemies        map[int]*Pac
	PelletsInSight map[int]Cells
	EnemiesInSight map[int]map[*Pac]*Cell // map[pacID][enemy]*Cell
	PacDestination Cells
}

func NewPacsManager(width, height int) PacsManager {
	return PacsManager{
		Width:          width,
		Height:         height,
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

func (p *PacsManager) ManagePac(mine bool, pacID, x, y int, pacType string, speedTurnsLeft, abilityCooldown int) {
	if mine {
		if pac, ok := p.Pacs[pacID]; ok {
			pac.Update(x, y, pacType, speedTurnsLeft, abilityCooldown)
		} else {
			p.Pacs[pacID] = NewPac(pacID, x, y, pacType, speedTurnsLeft, abilityCooldown)
		}
		p.PacIDs = append(p.PacIDs)
		p.PelletsInSight[pacID] = []*Cell{}
	} else {
		if pac, ok := p.Enemies[pacID]; ok {
			pac.Update(x, y, pacType, speedTurnsLeft, abilityCooldown)
		} else {
			p.Enemies[pacID] = NewPac(pacID, x, y, pacType, speedTurnsLeft, abilityCooldown)
		}
	}
}

func (p *PacsManager) RegisterPellet(i, j, value int) {
	for pacID, v := range p.Pacs {
		// TODO: Share the pellet to all pac
		p.PelletsInSight[pacID] = append(p.PelletsInSight[pacID], v.SeePellet(i, j, value))
	}
}

func (p *PacsManager) PrintPellets() {
	for k, v := range p.PelletsInSight {
		debug(k, v)
	}
}

func (p *PacsManager) NotifyEnemies() {
	for _, pac := range p.Pacs {
		if _, ok := p.EnemiesInSight[pac.ID]; !ok {
			p.EnemiesInSight[pac.ID] = make(map[*Pac]*Cell)
		}

		// FIXME: If an enemy disappear from sight, this can cause misleading info
		for _, enemy := range p.Enemies {
			p.EnemiesInSight[pac.ID][enemy] = pac.SeeEnemy(p.EnemiesInSight, enemy)
		}
	}
}

func (p *PacsManager) Decide() {
	for _, cells := range p.PelletsInSight {
		sort.Sort(cells)
	}
	var action []string
	var closestCells []*Cell

	for pacID := range p.Pacs {
		if cell := p.FindClosestCell(pacID, nil); cell != nil {
			closestCells = append(closestCells, cell)
		}
	}

	for pacID, pac := range p.Pacs {
		if rand.Float64() > 0.66 {
			if s, ok := pac.AskForSpeed(); ok {
				action = append(action, s)
				continue
			}
		}
		// Is there something to do with the enemies (go away, chase them, ... ?)
		if s, ok := pac.CheckEnemies(p); ok {
			action = append(action, s)
			continue
		} else {
			// Nothing interesting to do with enemies, focus on pellets

			// If none pellets are in sight, change strategy
			if len(p.PelletsInSight[pacID]) == 0 {
				// TODO: Go to the closest enemy of go to 0,0 if no enemy in sight
				//if len(p.EnemiesInSight[pacID]) > 0 {
				//	debug("go to random last enemy in sight position")
				//	for _, lastEnemyPosition := range p.EnemiesInSight[pacID] {
				//		action = append(action, pac.Move(lastEnemyPosition.X, lastEnemyPosition.Y))
				//		break
				//	}
				//	continue
				//}

				debug("go to random position")
				action = append(action, pac.Move(p.GenerateRandomCell()))
				continue
			}

			// if pac is stuck for some turn
			if pac.ImmobileSince > 2 {
				// TODO: For now, when it's stuck, he try to go to the farthest. But we can improve this (maybe try the random cell)
				if cell := p.FindRandomCell(pacID); cell != nil {
					action = append(action, pac.Move(cell.X, cell.Y))
					continue
				} else {
					action = append(action, pac.Move(p.GenerateRandomCell()))
				}
				// TODO: Should never occurred but, no farthest cell was found
				debug("NO FARTHEST CELL")
			}

			// Default strategy : go to closest pellets with highest value
			if cell := p.FindClosestCell(pacID, closestCells); cell != nil {
				action = append(action, pac.Move(cell.X, cell.Y))
				continue
			}

			// TODO: Should never occurred but, no closest cell was found
			debug("NO CLOSEST CELL")
			action = append(action, pac.Move(p.GenerateRandomCell()))
		}
	}

	fmt.Printf(strings.Join(action, "|") + "\n")
	p.PacDestination = []*Cell{}
	p.PacIDs = []int{}
}

func (p *PacsManager) FindClosestCell(pacID int, otherClosestCell Cells) *Cell {
	for _, cell := range p.PelletsInSight[pacID] {

		if otherClosestCell != nil {
			skip := false
			for _, otherCell := range otherClosestCell {
				if otherCell.Equals(cell) && otherCell.Dist < cell.Dist {
					skip = true
					break
				}
			}
			if skip {
				continue
			}
		}

		p.PacDestination = append(p.PacDestination, cell)
		return cell
	}
	return nil
}

// FIXME: Not a big deal but 2+ can go to the same cell
func (p *PacsManager) FindFarthestCell(pacID int) *Cell {
	for i := len(p.PelletsInSight[pacID]) - 1; i >= 0; i-- {
		cell := p.PelletsInSight[pacID][i]
		if existingCell := p.PacDestination.Contains(cell); existingCell == nil {
			p.PacDestination = append(p.PacDestination, cell)
			return cell
		}
	}
	return nil
}

func (p *PacsManager) FindRandomCell(pacID int) *Cell {
	length := len(p.PelletsInSight[pacID])
	for i := 0; i < length; i++ {
		k := rand.Intn(length - 1)
		cell := p.PelletsInSight[pacID][k]
		if existingCell := p.PacDestination.Contains(cell); existingCell != nil {
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

func NewPac(ID, x, y int, pacType string, left, cooldown int) *Pac {
	p := Pac{
		ID:              ID,
		X:               x,
		Y:               y,
		Type:            pacType,
		SpeedTurnsLeft:  left,
		AbilityCooldown: cooldown,
	}
	return &p
}

func (p *Pac) Update(x, y int, pacType string, left, cooldown int) {
	if p.X == x && p.Y == y {
		p.ImmobileSince++
	} else {
		p.ImmobileSince = 0
	}
	p.X = x
	p.Y = y
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
	debug("AskForSpeed", p.ID, p.AbilityCooldown)
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

// TODO: unused
func (p *Pac) FindClosestEnemy(pacManager *PacsManager) float64 {
	closestEnemy := float64(1000)
	for _, cell := range pacManager.EnemiesInSight[p.ID] {
		if closestEnemy > cell.Dist {
			closestEnemy = cell.Dist
		}
	}
	return closestEnemy
}

func (p *Pac) CheckEnemies(pacManager *PacsManager) (string, bool) {
	// If no enemy around ask for speed
	// No sure if it's efficient, because if next turn we see an enemy will can kill, we cannot chase him.
	//if p.FindClosestEnemy(pacManager) >= 5 {
	//	if s, ok := p.AskForSpeed(); ok {
	//		return s, ok
	//	}
	//}
	// debug("CheckEnemies", p.ID, pacManager.EnemiesInSight[p.ID])
	for enemy, cell := range pacManager.EnemiesInSight[p.ID] {
		// debug("CheckEnemies", enemy, cell)
		if cell.Trend < 0 && cell.Value < 0 && cell.Dist <= 3 {
			// Danger: an enemy is coming and can eat me. Try to morph into his weak spot
			debug("danger", p.ID, enemy, cell)
			if s, ok := p.AskForMorph(weakSpot(enemy.Type)); ok {
				return s, ok
			} else {
				// TODO: an enemy is coming and we fail switch. Try to go add the opposite
				if cell := pacManager.FindFarthestCell(p.ID); cell != nil {
					return p.Move(cell.X, cell.Y), true
				}
			}
		} else if cell.Trend < 0 && cell.Value > 0 && cell.Dist <= 2 {
			// Go for it. You can eat him
			// FIXME
			//if s, ok := p.AskForSpeed(); ok {
			//	return s, ok
			//}
			//return p.Move(enemy.X, enemy.Y), true
		} else if cell.Trend < 0 && cell.Value == 0 {
			// Cannot hurt me
		} else if cell.Trend > 0 && cell.Value < 0 {
			// A danger is leaving
		} else if cell.Trend > 0 && cell.Value > 0 && cell.Dist <= 2 {
			// An opponent that I can eat is leaving. Maybe chase him
			// FIXME
			//if s, ok := p.AskForSpeed(); ok {
			//	return s, ok
			//}
			//return p.Move(enemy.X, enemy.Y), true
		} else if cell.Trend > 0 && cell.Value == 0 {
			// An opponent that cannot hurt me if leaving. Ignore him.
		} else if cell.Trend == 0 && cell.Value < 0 {
			// I'm following an enemy that can eat me. Consider to go away from him
			if cell := pacManager.FindFarthestCell(p.ID); cell != nil {
				return p.Move(cell.X, cell.Y), true
			}
		} else if cell.Trend == 0 && cell.Value > 0 && cell.Dist <= 2 {
			// I'm following an enemy I can eat. Maybe chase him
			// FIXME
			//debug("Ask to speed up to kill this bastard", p.ID)
			//if s, ok := p.AskForSpeed(); ok {
			//	return s, ok
			//}
			//debug("No speed up to kill this bastard", p.ID, enemy)
			//return p.Move(enemy.X, enemy.Y), true
		} else if cell.Trend == 0 && cell.Value == 0 {
			// I'm following an enemy that does matter. Ignore him
		}
	}
	return "", false
}

func (p *Pac) String() string {
	return fmt.Sprintf("PAC(%v)(%+v;%v)(%v %v)", p.Type, p.X, p.Y, p.SpeedTurnsLeft, p.AbilityCooldown)
}

type Cell struct {
	X     int
	Y     int
	Value int
	Dist  float64
	Trend float64
}

type Cells []*Cell

func (c *Cell) String() string {
	return fmt.Sprintf("C(%v %v %v %.3f)", c.X, c.Y, c.Value, c.Dist)
}

func (c *Cell) Equals(other *Cell) bool {
	return c.X == other.X && c.Y == other.Y
}

func (c *Cell) Update(x, y, value int, dist float64) {
	c.Trend = 0
	if dist < c.Dist {
		c.Trend = -1 * dist
	} else if dist > c.Dist {
		c.Trend = dist
	}

	c.X = x
	c.Y = y
	c.Value = value
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

func (c *Cells) Contains(value *Cell) *Cell {
	if value == nil {
		return nil
	}
	for _, cell := range *c {
		if cell.Equals(value) {
			return cell
		}
	}
	return nil
}

func (c Cells) Len() int      { return len(c) }
func (c Cells) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c Cells) Less(i, j int) bool {
	if c[i].Value > c[j].Value {
		return true
	} else if c[i].Value == c[j].Value {
		return c[i].Dist < c[j].Dist
	}
	return false
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

func debug(values ...interface{}) {
	fmt.Fprintln(os.Stderr, values...)
}

/*
 Sandbox
*/
func main() {
	pacsManager := NewPacsManager(50, 50)
	pacsManager.ManagePac(true, 0, 7, 6, Rock, 0, 0)
	pacsManager.ManagePac(true, 1, 23, 9, Paper, 0, 0)
	fmt.Println(pacsManager.String())

	initPellet(&pacsManager, 2)

	fmt.Println(pacsManager.String())

	pacsManager.Decide()

	pacsManager.ManagePac(true, 0, 7, 5, Rock, 5, 9)
	pacsManager.ManagePac(true, 1, 28, 3, Rock, 5, 9)
	initPellet(&pacsManager, 1)
	pacsManager.Decide()
	//
	//pacsManager.ManagePac(true, 0, 28, 5, Rock, 4, 8)
	//pacsManager.ManagePac(true, 1, 19, 9, Rock, 4, 8)
	//initPellet(&pacsManager, 2)
	//pacsManager.Decide()
}

func initPellet(pacsManager *PacsManager, super int) {
	pacsManager.RegisterPellet(8, 7, 1)
	pacsManager.RegisterPellet(7, 7, 1)
	pacsManager.RegisterPellet(6, 5, 1)
	pacsManager.RegisterPellet(5, 5, 1)
	pacsManager.RegisterPellet(8, 5, 1)
	pacsManager.RegisterPellet(9, 5, 1)

	pacsManager.RegisterPellet(22, 9, 1)
	pacsManager.RegisterPellet(21, 9, 1)
	pacsManager.RegisterPellet(20, 9, 1)
	pacsManager.RegisterPellet(23, 8, 1)
	pacsManager.RegisterPellet(23, 7, 1)
	pacsManager.RegisterPellet(23, 6, 1)
	if super == 2 {
		pacsManager.RegisterPellet(7, 5, 10)
	}
	if super > 0 {
		pacsManager.RegisterPellet(31, 10, 10)
	}
}
