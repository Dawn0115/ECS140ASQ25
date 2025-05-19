isUnion([], Set, Set).
isUnion([H|T], Set, Union) :-
    member(H, Set),           
    isUnion(T, Set, Union).
isUnion([H|T], Set, [H|Union]) :-
    \+ member(H, Set),
    isUnion(T, Set, Union).


isIntersection([], _, []).
isIntersection([H|T], Set2, [H|I]) :-
    member(H, Set2),          % H is in the second set, add to I
    isIntersection(T, Set2, I).
isIntersection([H|T], Set2, I) :-
    \+ member(H, Set2),       % H not in the second set, skip it
    isIntersection(T, Set2, I).

isEqual(S1, S2) :-
    \+ (member(ElementinS1, S1), \+ member(ElementinS1, S2)), \+ (member(ElementinS2, S2), \+ member(ElementinS2, S1)).
