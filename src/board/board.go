package board

import "sync"

type Board struct {
	height uint
	width  uint

	CellBoard [1000][1000]Cell
	AnimalMap [1000][1000]*Animal
	Players   map[string]*Player

	movementLock sync.Mutex
}

type Location struct {
	height uint
	width  uint
}

func NewBoard(height, width uint) *Board {
	b := &Board{height: height,
		width: width}

	b.CellBoard = [1000][1000]Cell{}
	b.AnimalMap = [1000][1000]*Animal{}
	b.Players = make(map[string]*Player)

	return b
}

func (b *Board) StepN(l *Location, n uint) *Location {
	pos := l.width + l.height*b.width
	pos = pos + n

	if pos >= b.width*b.height {
		return &Location{
			height: b.height - 1,
			width:  b.width - 1,
		}
	}
	return &Location{
		height: pos / b.width,
		width:  pos % b.width,
	}
}

func (b *Board) AddPlayerToStart(name string, pieceColor string) {
	start := Location{0, 0}

	p := Player{
		name:       name,
		pieceColor: pieceColor,
		loc:        start,
	}

	b.CellBoard[0][0].PlacePlayer(&p)
	b.Players[p.name] = &p
}

func (b *Board) Move(playerName string, diceResult uint) {
	b.movementLock.Lock()

	p := b.Players[playerName]

	newLoc := b.StepN(&p.loc, diceResult)

	animal := b.AnimalMap[newLoc.width][newLoc.height]
	if animal != nil {
		newLoc = animal.move(*newLoc, diceResult)
	}

	oldCell := &b.CellBoard[p.loc.width][p.loc.height]
	newCell := &b.CellBoard[newLoc.width][newLoc.height]

	oldCell.PickPlayer(p)
	p.loc = *newLoc
	newCell.PlacePlayer(p)

	b.movementLock.Unlock()
}

func (b *Board) CheckWin() bool {

	endCell := b.CellBoard[b.width-1][b.height-1]
	if len(endCell.Occupiers) != 0 {
		return true
	}
	return false
}

func (b *Board) GetWinner() string {
	endCell := b.CellBoard[b.width-1][b.height-1]
	for n := range endCell.Occupiers {
		return n
	}
	return ""
}

func NewLocation(height, width uint) *Location {
	return &Location{
		height: height,
		width:  width,
	}
}
