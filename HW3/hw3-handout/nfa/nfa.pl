
    %% remove fail and add body/other cases for this predicate
%Base case: 
reachable(State, State, []).

% 2) Recursive step: read the first symbol, pick one of the possible next states,
%    and recurse on the tail of the input.
reachable(Start, Final, [Symbol|Rest]) :-
    transition(Start, Symbol, NextStates),  % NextStates is the list of states reachable from Start on Symbol
    member(Next, NextStates),            % choose one of them
    reachable(Next, Final, Rest).          %recursion on the rest of the input