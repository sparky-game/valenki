package parchis

type Piece struct {
	ID       int
	Color    string
	Position int
	Finished bool
}

func NewPiece(id int, color string) *Piece {
	return &Piece{
		ID:       id,
		Color:    color,
		Position: HOME_POSITION,
		Finished: false,
	}
}

func (p *Piece) IsInHouse() bool {
	return p.Position == HOME_POSITION
}

func (p *Piece) CanMove() bool {
	return !p.IsInHouse() && !p.Finished
}

func (p *Piece) ReturnToHouse() {
	p.Position = HOME_POSITION
	p.Finished = false
}
