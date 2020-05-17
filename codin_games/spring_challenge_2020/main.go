package main

import (
	"fmt"
	"strings"
)

// [x] Optimiser le fait que c'est au Pac le plus proche d'aller chercher la cell avec la plus de value
// [ ] Il nous faudrait une info pour savoir si l'info et la position connu de l'ennemy est fraiche et si on peut s'y fier
// [ ] Garder une trace des pellets meme si on ne les voit plus, avoir un système de "est-ce à jour ?"
// [ ] Prévoir un buffer d'action pour un PAC ? Genre il faut speed + aller là, ça éviterai de devoir recalculer un truc qu'on sait déjà et se tromper
// [ ] Pouvoir garder la meme direction d'un tour à l'autre
// [ ] Déclencher un coup de speed si on voit 5 pellet devant nous sans obstacle
// [ ] On devrait retirer aussi des pellets in sights si on a la vision sur une ligne et que la pellet ni est plus

/*
 Sandbox
*/
func main() {
	width, height := 30, 15
	pacsManager := NewPacsManager(width, height)
	for i := 0; i < height; i++ {
		pacsManager.ManageGrid(i, strings.Repeat(" ", width))
	}

	pacsManager.NewRound()
	pacsManager.ManagePac(true, 0, 11, 1, Rock, 0, 5)
	pacsManager.ManagePac(true, 1, 3, 11, Paper, 0, 5)
	pacsManager.ManagePac(true, 2, 9, 7, Scissors, 0, 5)
	pacsManager.ManagePac(true, 3, 9, 5, Rock, 0, 5)
	fmt.Println(pacsManager.String())
	pacsManager.NotifyPacs()

	initPellet(&pacsManager, 0)

	fmt.Println(pacsManager.String())

	pacsManager.Decide()

	// New turn
	//pacsManager.NewRound()
	//pacsManager.ManagePac(true, 0, 21, 11, Rock, 5, 10)
	//pacsManager.ManagePac(true, 1, 3, 5, Paper, 5, 10)
	//pacsManager.ManagePac(true, 2, 15, 11, Scissors, 5, 10)
	//pacsManager.ManagePac(true, 3, 9, 3, Rock, 5, 10)
	//pacsManager.NotifyPacs()
	//
	//initPellet(&pacsManager, 1)
	//pacsManager.Decide()

}

func initPellet(pacsManager *PacsManager, super int) {
	pacsManager.RegisterPellet(12, 1, 1)
	pacsManager.RegisterPellet(13, 1, 1)
	pacsManager.RegisterPellet(14, 1, 1)
	pacsManager.RegisterPellet(15, 1, 1)
	pacsManager.RegisterPellet(16, 1, 1)
}
