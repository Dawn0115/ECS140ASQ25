package nfa

// A state in the NFA is labeled by a single integer.
type state uint

// TransitionFunction tells us, given a current state and some symbol, which
// other states the NFA can move to.
//
// Deterministic automata have only one possible destination state,
// but we're working with non-deterministic automata.
type TransitionFunction func(st state, act rune) []state

func Reachable(
	// `transitions` tells us what our NFA looks like
	transitions TransitionFunction,
	// `start` and `final` tell us where to start, and where we want to end up
	start, final state,
	// `input` is a (possible empty) list of symbols to apply.
	input []rune,


) bool {
	// TODO Write the Reachable function,
	// return true if the nfa accepts the input and can reach the final state with that input,
	// return false otherwise
	if len(input) == 0 {
		return start == final
	}
	currentStates := []state{start}
	for _, symbol := range input {
		var nextStates []state
	
		for _, st := range currentStates {
			resultStates := transitions(st, symbol)
			for _, newSt := range resultStates {
				nextStates = append(nextStates, newSt)
			}
		}
		currentStates = nextStates
	
		// If there are no reachable states, we can stop processing early.
		if len(currentStates) == 0 {
			return false
		}
	}
	for _, st := range currentStates {
		if st == final {
			return true
		}
	}
	return false

	panic("TODO: implement this!")
}
