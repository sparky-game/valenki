package parchis

type Player struct {
	Name   string
	Color  string
	Pieces []*Piece
}

func NewPlayer(name, color string) *Player {
	pieces := make([]*Piece, 4)
	for i := 0; i < 4; i++ {
		pieces[i] = NewPiece(i+1, color)
	}
	return &Player{
		Name:  name,
		Color: color,
		Pieces, pieces,
	}
}

func (p *Player) GetPieceByID(id int) (*Piece, error) {
	if id < 1 || id > 4 {
		return nil, errors.New("Invalid piece ID")
	}
	return p.Pieces[id-1], nil
}

func (p *Player) AllPiecesAtHome() bool {
	for _, piece := range p.Pieces {
		if piece.Position != -1 {
			return false
		}
	}
	return true
}

func (p *Player) CanMoveAnyPiece() bool {
	for _, piece := range p.Pieces {
		if piece.CanMove() {
			return true
		}
	}
}

func (p *Player) HasWon() bool {
	for _, piece := range p.Pieces {
		if !piece.Finished {
			return false
		}
	}
	return true
}
