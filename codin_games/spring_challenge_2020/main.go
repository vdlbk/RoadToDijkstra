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
	width, height := 50, 50
	pacsManager := NewPacsManager(width, height)
	for i := 0; i < height; i++ {
		pacsManager.ManageGrid(i, strings.Repeat(" ", width))
	}

	pacsManager.NewRound()
	pacsManager.ManagePac(true, 0, 7, 6, Rock, 0, 0)
	pacsManager.ManagePac(true, 1, 23, 9, Paper, 0, 0)
	pacsManager.ManagePac(false, 1, 11, 11, Paper, 0, 0)
	fmt.Println(pacsManager.String())
	pacsManager.NotifyPacs()

	initPellet(&pacsManager, 2)

	fmt.Println(pacsManager.String())

	pacsManager.Decide()

	pacsManager.ManagePac(true, 0, 7, 5, Rock, 5, 9)
	pacsManager.ManagePac(true, 1, 28, 3, Rock, 5, 9)
	pacsManager.ManagePac(false, 1, 11, 12, Paper, 0, 0)
	pacsManager.NotifyPacs()
	initPellet(&pacsManager, 1)
	pacsManager.Decide()
	//
	//pacsManager.ManagePac(true, 0, 28, 5, Rock, 4, 8)
	//pacsManager.ManagePac(true, 1, 19, 9, Rock, 4, 8)
	//initPellet(&pacsManager, 2)
	//pacsManager.Decide()

	fmt.Println(pacsManager.GenerateCellFarFrom(5, pacsManager.Pacs[0]))
}

func initPellet(pacsManager *PacsManager, super int) {
	pacsManager.RegisterPellet(15, 11, 1)
	pacsManager.RegisterPellet(16, 11, 10)
	pacsManager.RegisterPellet(17, 11, 1)
	pacsManager.RegisterPellet(18, 11, 10)
	pacsManager.RegisterPellet(19, 11, 1)
	pacsManager.RegisterPellet(20, 11, 1)
	pacsManager.RegisterPellet(21, 11, 1)
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

	pacsManager.RegisterPellet(3, 5, 1)
	pacsManager.RegisterPellet(3, 6, 1)
	pacsManager.RegisterPellet(3, 7, 1)
	pacsManager.RegisterPellet(3, 8, 1)
	pacsManager.RegisterPellet(3, 9, 1)
	pacsManager.RegisterPellet(3, 10, 1)
	pacsManager.RegisterPellet(4, 9, 1)
	pacsManager.RegisterPellet(5, 9, 1)
	pacsManager.RegisterPellet(5, 10, 1)
	pacsManager.RegisterPellet(5, 11, 1)
	pacsManager.RegisterPellet(6, 11, 1)
	pacsManager.RegisterPellet(7, 11, 1)

	pacsManager.RegisterPellet(10, 3, 1)
	pacsManager.RegisterPellet(9, 3, 1)
	pacsManager.RegisterPellet(8, 3, 1)
	pacsManager.RegisterPellet(7, 3, 1)
	pacsManager.RegisterPellet(11, 4, 1)
	pacsManager.RegisterPellet(11, 7, 1)
}
