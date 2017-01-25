package main

import (
	"board"
	"fmt"
)

func main() {
	b := board.NewBoard(10, 10)

	b.AddSnake(5, 5, 2, 2)
	b.AddLadder(3, 3, 6, 6)

	tortoiseMove := func(currentLoc board.Location, diceResult uint) *board.Location {

		nStep := diceResult * 2

		newLoc := b.StepN(&currentLoc, nStep)

		return newLoc
	}
	b.AddAnimal("tortoise", 4, 4, tortoiseMove)

	b.AddPlayerToStart("harshad", "yellow")
	b.AddPlayerToStart("jai", "red")

	for true {
		for np := range b.Players {
			fmt.Println(np)
			d := board.DiceRoll()
			b.Move(np, d)
			if b.CheckWin() {
				fmt.Println("We have a winner, " + b.GetWinner())
				return
			}
		}
	}
}
