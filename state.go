package main

import "fmt"

type Placement struct {
	tile     int
	rotation int
}

func (p *Placement) MapFrom(pipeIndex int) int {
	in := (pipeIndex - 2*p.rotation)
	if in < 0 {
		in += 8
	}
	out := (Tiles[p.tile][in] + 2*p.rotation) % 8
	return out
}

func (p *Placement) String() string {
	if p == nil {
		return "----  "
	}
	return fmt.Sprintf("%d, %d  ", p.tile, p.rotation)
}

type State [9]*Placement

func (s State) IsComplete() bool {
	for _, placement := range s {
		if placement == nil {
			return false
		}
	}
	return true
}

func (s State) IsValid() bool {
	for slotIndex, placement := range s {
		if placement == nil {
			continue
		}
		for pipeIndex, pipe := range TheBoard.Slot(slotIndex) {
			if terminal, ok := pipe.(Terminal); ok {
				destValue := s.ValueAt(slotIndex, pipeIndex)
				if destValue == 0 || destValue == int(terminal) {
					continue
				}
				//fmt.Printf("invalid: %d != %d at (%d, %d)\n", destValue, int(terminal), slotIndex, pipeIndex)
				return false
			}
		}
	}
	return true
}

func (s State) String() string {
	r := ""

	for _, i := range []int{0, 1, 2} {
		r += s[i].String()
	}
	r += "\n"

	for _, i := range []int{7, 8, 3} {
		r += s[i].String()
	}
	r += "\n"

	for _, i := range []int{6, 5, 4} {
		r += s[i].String()
	}
	r += "\n"
	return r
}

func (s *State) ValueAt(slotIndex, pipeIndex int) int {
	placement := s[slotIndex]
	if placement == nil {
		return 0
	}
	targetIndex := placement.MapFrom(pipeIndex)
	return TheBoard[slotIndex].Pipe(targetIndex).Value(s)
}

func (s *State) FirstOpenSlot() int {
	for index, placement := range s {
		if placement == nil {
			return index
		}
	}
	panic(fmt.Errorf("no available slots"))
}

func (s *State) AvailableTiles() []int {
	var taken [9]bool

	for _, placement := range s {
		if placement == nil {
			continue
		}
		taken[placement.tile] = true
	}

	var result []int

	for index, val := range taken {
		if !val {
			result = append(result, index)
		}
	}
	return result
}
