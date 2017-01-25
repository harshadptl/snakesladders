package board

type Cell struct {
	Occupiers map[string]*Player
}

func (c *Cell) PlacePlayer(p *Player) {

	if c.Occupiers == nil {
		c.Occupiers = make(map[string]*Player)
	}
	c.Occupiers[p.name] = p
}

func (c *Cell) PickPlayer(p *Player) {
	delete(c.Occupiers, p.name)
}
