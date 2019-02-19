package main

import "fmt"

func main() {
	// initalize our state
	var s State
	s[8] = &Placement{1, 1} // the middle tile is our tile 1 at rotation 1

	search(&s)
	fmt.Printf("%d total tries\n", count)
}

var count int

func search(s *State) bool {
	var recurse func() bool
	min := 10

	recurse = func() bool {
		if s.IsComplete() && s.IsValid() {
			// found a solution
			fmt.Printf("found a solution:\n")
			fmt.Printf("%s", s)
			return true
		}

		// which slot are we working on?
		slotIndex := s.FirstOpenSlot()

		// which tiles are unused?
		availableTiles := s.AvailableTiles()

		if len(availableTiles) < min {
			min = len(availableTiles)
			fmt.Printf("%d tiles left -- %d tries\n", min, count)
		}

		for _, tileIndex := range availableTiles {
			for rot := 0; rot < 4; rot++ {
				s[slotIndex] = &Placement{tileIndex, rot}
				count++
				if s.IsValid() {
					recurse()
				}
				s[slotIndex] = nil
			}
		}
		return false
	}
	return recurse()
}
