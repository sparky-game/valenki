package parchis

type Cell struct {
	OccupiedBy *Piece
}

type Board struct {
	Cells     []Cell
	SafeZones map[string][]Cell
	SafeCells []int
}

func NewBoard() *Board {
	return &Board{
		Cells: make([]Cell, 68),
		SafeZones: map[string][]Cell{
			"Red":    make([]Cell, 8),
			"Blue":   make([]Cell, 8),
			"Yellow": make([]Cell, 8),
			"Green":  make([]Cell, 8),
		},
		SafeCells: []int{
			5, 12, 17, 22, 29, 34, 39, 46, 51, 56, 63, 68,
		},
	}
}

func (b *Board) IsSafeCell(p int) bool {
	for _, sp := range b.SafeCells {
		if sp == p {
			return true
		}
	}
	return false
}

func (b *Board) IsSafeZone(pc *Piece) bool {
	sz, ok := b.SafeZones[pc.Color]
	if !ok {
		return false
	}
	for _, cell := range sz {
		for _, p := range cell.OccupiedBy {
			if p == pc {
				return true
			}
		}
	}
	return false
}

func (b *Board) MovePiece(p *Piece, n int) error {
	// Remove piece from current cell
	if p.Position != -1 && p.Position < len(b.Cells) {
		b.removePieceFromCell(p.Position, p)
	}
	// Update piece position
	p.Position += n
	if p.Position >= 68 {
		p.Position = 68
		p.Finished = true
	}
	// Place piece on new cell
	if p.Position < 68 {
		if b.IsSafeCell(p.Position) {
			b.addPieceToCell(p.Position, p)
		} else {
			if len(b.Cells[p.Position].OccupiedBy) >= 2 && b.Cells[p.Position].OccupiedBy[0].Color == p.Color {
				return errors.New("A barrier blocks your path")
			} else {
				b.addPieceToCell(p.Position, p)
			}
		}
	}
}

func (b *Board) addPieceToCell(p int, pc *Piece) {
	b.Cells[p].OccupiedBy = append(b.Cells[p].OccupiedBy, pc)
}

func (b *Board) removePieceFromCell(p int, pc *Piece) {
	cell := &b.Cells[p]
	for i, piece := range cell.OccupiedBy {
		if piece == pc {
			cell.OccupiedBy = append(cell.OccupiedBy[:i], cell.OccupiedBy[i+1:]...)
			return
		}
	}
}
