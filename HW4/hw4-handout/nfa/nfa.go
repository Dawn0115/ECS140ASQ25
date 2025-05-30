package nfa

import "sync"
 
// A nondeterministic Finite Automaton (NFA) consists of states,
// symbols in an alphabet, and a transition function.

// A state in the NFA is represented as an unsigned integer.
type state uint

// Given the currentStatesent state and a symbol, the transition function
// of an NFA returns the set of next states the NFA can transition to
// on reading the given symbol.
// This set of next states could be empty.
type TransitionFunction func(st state, sym rune) []state

// Reachable returns true if there exists a sequence of transitions
// from `transitions` such that if the NFA starts at the start state
// `start` it would reach the final state `final` after reading the
// entire sequence of symbols `input`; Reachable returns false otherwise.
func Reachable(
	// `transitions` tells us what our NFA looks like
	transitions TransitionFunction,
	// `start` and `final` tell us where to start, and where we want to end up
	start, final state,
	// `input` is a (possible empty) list of symbols to apply.
	input []rune,
) bool {
	// TODO
	if start == final && len(input) == 0 {
		return true
	}
	currentStates := map[state]struct{}{start: {}}

	// Process each symbol in the input sequence.
	for _, symbol := range input {
		next := make(map[state]struct{})
		var wg sync.WaitGroup
		var mu sync.Mutex

		// Launch a goroutine for each currentStatesent 
		for st := range currentStates {
			wg.Add(1)
			st := st  
			go func() {
				defer wg.Done()
				for _, ns := range transitions(st, symbol) {
					mu.Lock()
					next[ns] = struct{}{}
					mu.Unlock()
				}
			}()
		}

		wg.Wait()
		if len(next) == 0 {
			return false
		}
		currentStates = next
	}
	_, ok := currentStates[final]
	return ok
}
