package board

type AnimalMove func(currentLoc Location, diceResult uint) *Location

type Animal struct {
	name string
	loc  Location
	move AnimalMove
}

func (b *Board) AddSnake(sourceH, sourceW, sinkH, sinkW uint) {
	//TODO: check valid snake

	b.AddPipeAnimal("snake", sourceH, sourceW, sinkH, sinkW)
}

func (b *Board) AddLadder(sourceH, sourceW, sinkH, sinkW uint) {
	//TODO: check valid ladder

	b.AddPipeAnimal("ladder", sourceH, sourceW, sinkH, sinkW)
}

func (b *Board) AddPipeAnimal(pipeName string, sourceH, sourceW, sinkH, sinkW uint) {

	pipeMove := func(currentLoc Location, diceResult uint) *Location {
		return &Location{
			height: sinkH,
			width:  sinkW,
		}
	}

	pipe := Animal{
		name: pipeName,
		loc: Location{
			height: sourceH,
			width:  sourceW,
		},
		move: pipeMove,
	}

	b.AnimalMap[sourceW][sourceH] = &pipe
}

func (b *Board) AddPython(pipeName string, sourceH, sourceW, sinkH, sinkW uint) {

	pipeMove := func(currentLoc Location, diceResult uint) *Location {
		return &Location{
			height: sinkH,
			width:  sinkW,
		}
	}

	pipe := Animal{
		name: pipeName,
		loc: Location{
			height: sourceH,
			width:  sourceW,
		},
		move: pipeMove,
	}

	b.AnimalMap[sourceW][sourceH] = &pipe
}

func (b *Board) AddAnimal(name string, sourceH, sourceW uint, move AnimalMove) {
	anim := Animal{
		name: name,
		loc: Location{
			height: sourceH,
			width:  sourceW,
		},
		move: move,
	}

	b.AnimalMap[sourceW][sourceH] = &anim
}
