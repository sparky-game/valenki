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
		Position: -1,
		Finished: false,
	}
}

func (p *Piece) Move(n int, b *Board) {
	if !p.Finished {
		newPos := p.Position + n
		if newPos >= 68 {
			newPos = 68
			p.Finished = true
		}
		b.MovePiece(p, newPos-p.Position)
		p.Position = newPos
	}
}

func (p *Piece) IsInHouse() bool {
	return p.Position == -1
}

func (p *Piece) CanMove() bool {
	return !p.IsInHouse() && !p.Finished
}

func (p *Piece) ResetToHouse() {
	p.Position = -1
	p.Finished = false
}
