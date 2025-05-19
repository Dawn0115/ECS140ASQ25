isUnion([], S, S).
isUnion([X|Xs], S, U) :-
    member(X, S),             % X already in S, skip it
    isUnion(Xs, S, U).
isUnion([X|Xs], S, [X|U]) :-
    \+ member(X, S),          % X not in S, include it
    isUnion(Xs, S, U).

isIntersection(Set1,Set2,Intersection) :-
    %% remove fail and add body/other cases for this predicate
    fail.

isEqual(Set1,Set2) :-
    %% remove fail and add body/other cases for this predicate
    fail.
