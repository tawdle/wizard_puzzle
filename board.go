package main

type Slot [8]Pipe

func (s *Slot) Pipes() [8]Pipe {
	return [8]Pipe(*s)
}

func (s Slot) Pipe(i int) Pipe {
	return s[i]
}

type Board [9]*Slot

func (b Board) Slot(i int) *Slot {
	return b[i]
}

type Terminal int

func (t Terminal) Value(b *State) int {
	return int(t)
}

type Connector struct {
	slot     int
	position int
}

func (c Connector) Target() Connector {
	slot := TheBoard.Slot(c.slot)
	pipe := slot.Pipe(c.position)
	return pipe.(Connector)
}

func (c Connector) Value(s *State) int {
	return s.ValueAt(c.slot, c.position)
}

type Pipe interface {
	Value(b *State) int
}

var TheBoard Board

func init() {
	T := func(value int) Terminal {
		return Terminal(value)
	}

	C := func(slotIndex, positionIndex int) Connector {
		return Connector{slotIndex, positionIndex}
	}

	TheBoard = Board{
		&Slot{T(1), T(2), C(1, 7), C(1, 6), C(7, 1), C(7, 0), T(7), T(12)},
		&Slot{T(3), T(4), C(2, 7), C(2, 6), C(8, 1), C(8, 0), C(0, 3), C(0, 2)},
		&Slot{T(5), T(6), T(5), T(7), C(3, 1), C(3, 0), C(1, 3), C(1, 2)},
		&Slot{C(2, 5), C(2, 4), T(6), T(8), C(4, 1), C(4, 0), C(8, 3), C(8, 2)},
		&Slot{C(3, 5), C(3, 4), T(9), T(2), T(4), T(9), C(5, 3), C(5, 2)},
		&Slot{C(8, 5), C(8, 4), C(4, 7), C(4, 6), T(8), T(3), C(6, 3), C(6, 2)},
		&Slot{C(7, 5), C(7, 4), C(5, 7), C(5, 6), T(10), T(11), T(10), T(1)},
		&Slot{C(0, 5), C(0, 4), C(8, 7), C(8, 6), C(6, 1), C(6, 0), T(11), T(12)},
		&Slot{C(1, 5), C(1, 4), C(3, 7), C(3, 6), C(5, 1), C(5, 0), C(7, 3), C(7, 2)},
	}
}
