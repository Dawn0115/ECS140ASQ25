:- initialization main.

main :-
    consult(['transitions.pl', 'nfa.pl']),
    (show_coverage(run_tests) ; true),
    halt.
     
:- begin_tests(nfa).

test(nfa1, [nondet]) :- reachable(0, 2, [a]).
test(nfa2, [nondet]) :-  reachable(0, 2, [b]).
test(nfa3, [nondet]) :-  reachable(0, 1, [a, b, a]).
test(nfa4, [fail]) :- reachable(0, 1, [a, b, a, b]).
test(nfa5, [nondet]) :- reachable(0, 2, [a, b, a]).

test(nfa6, [nondet]) :- reachable(10, 13, [a,b]).
test(nfa7, [nondet]) :-  reachable(10, 13, [a,c]).
test(nfa8, [nondet]) :-  reachable(10, 11, [a]).
test(nfa9, [fail]) :- reachable(10, 13, [a,a]).
test(nfa10, [fail]) :- reachable(10, 13, [a]).
test(nfa11, [fail]) :- reachable(10, 11, [b]).

test(nfa1, [nondet]) :- reachable(20, 20, [a,b,b]).
test(nfa2, [nondet]) :-  reachable(20, 21, [a,a,b]).
test(nfa3, [nondet]) :-  reachable(20, 20, [a, a, a, a, a]).
test(nfa4, [fail]) :- reachable(20, 21, [a, a]).
test(nfa5, [fail]) :- reachable(20, 20, [a, b, a, a]).


test(nfa12, [nondet]) :- reachable(0, 1, [a]).
test(nfa13, [fail])   :- reachable(0, 1, [a, b]).
test(nfa14, [nondet]) :- reachable(0, 0, [a, b]).
test(nfa15, [fail])   :- reachable(1, 0, []).
test(nfa16, [nondet]) :- reachable(1, 1, []).
test(nfa17, [nondet]) :- reachable(2, 2, []).
test(nfa18, [fail])   :- reachable(2, 0, [b]).
test(nfa19, [fail])   :- reachable(1, 0, [b, b]).
test(nfa20, [fail])   :- reachable(1, 0, [b, a]).
test(nfa21, [nondet]) :- reachable(1, 0, [b, a, b]).
test(nfa22, [fail])   :- reachable(1, 2, [b, a, b]).
test(nfa23, [nondet]) :- reachable(1, 1, [b, a, b, a]).
test(nfa24, [fail])   :- reachable(2, 1, [a]).
test(nfa25, [fail])   :- reachable(2, 2, [a]).
test(nfa26, [fail])   :- reachable(2, 2, [a, b]).

test(nfa27, [nondet]) :- reachable(10, 12, [a]).
test(nfa28, [fail])   :- reachable(10, 11, [a, b, c]).
test(nfa29, [fail])   :- reachable(10, 13, [a, b, c]).
test(nfa30, [fail])   :- reachable(10, 13, [a, b, b]).
test(nfa31, [fail])   :- reachable(10, 13, [a, c, b]).
test(nfa32, [fail])   :- reachable(10, 13, [a, c, a]).
test(nfa33, [nondet]) :- reachable(10, 10, []).
test(nfa34, [fail])   :- reachable(10, 13, []).

test(nfa35, [nondet]) :- reachable(20, 20, [a]).
test(nfa36, [nondet]) :- reachable(20, 20, [b, a, b, a]).
test(nfa37, [nondet]) :- reachable(20, 21, [b]).
test(nfa38, [nondet]) :- reachable(21, 20, [b]).
test(nfa39, [fail])   :- reachable(21, 21, [b, a, a]).
test(nfa40, [nondet]) :- reachable(21, 21, [a, a, a]).
test(nfa41, [nondet]) :- reachable(21, 21, []).
test(nfa42, [nondet]) :- reachable(20, 21, [a, b, a]).
test(nfa43, [nondet]) :- reachable(20, 20, [b, b]).
test(nfa44, [fail])   :- reachable(20, 21, [b, b, a]).
test(nfa45, [nondet]) :- reachable(20, 20, [b, b, a]).
test(nfa46, [fail])   :- reachable(20, 21, [b, b]).
test(nfa47, [nondet]) :- reachable(20, 21, [b, b, b]).
test(nfa48, [nondet]) :- reachable(20, 20, [b, b, b, b]).
test(nfa49, [fail])   :- reachable(20, 21, [a, a, b, b, a]).
test(nfa50, [fail])   :- reachable(20, 21, [b, a, b]).
test(nfa51, [nondet]) :- reachable(20, 21, [b, a, b, a, b]).

test(nfa52, [true])  :- reachable(1,1,[]).
test(nfa53, [fail])  :- reachable(1,0,[]).
test(nfa54, [fail])  :- reachable(1,0,[a]).   % transition(1,a,[]) so no move
test(nfa55, [fail])  :- reachable(0,2,[c]).   % no transition(_,c,_)
test(nfa56, [nondet]) :- reachable(0,0,[a,b,a,b,a,b]).
test(nfa57, [fail]) :- reachable(0,2,[a,a,b]).
test(nfa58, [fail])  :- reachable(0,2,[a,b,a,a]).  
test(nfa59, [true])  :- reachable(10,10,[]).
test(nfa60, [fail])  :- reachable(10,11,[]).
test(nfa61, [nondet]) :- reachable(10,11,[a]).
test(nfa62, [nondet]) :- reachable(10,12,[a]).
test(nfa63, [fail])  :- reachable(10,13,[b]).
test(nfa64, [fail])  :- reachable(10,13,[c]).
test(nfa65, [nondet]) :- reachable(10,13,[a,b]).
test(nfa66, [nondet]) :- reachable(10,13,[a,c]).
test(nfa67, [fail])  :- reachable(10,13,[a,c,b]).
test(nfa68, [fail])  :- reachable(10,13,[a,b, extra]).  
test(nfa69, [true])  :- reachable(20,20,[]).
test(nfa70, [fail])  :- reachable(20,21,[]).
test(nfa71, [nondet]) :- reachable(20,20,[a,a,a]).
test(nfa72, [nondet]) :- reachable(20,21,[b]).
test(nfa73, [nondet]) :- reachable(21,20,[b]).
test(nfa74, [nondet]) :- reachable(20,20,[b,a,b]).
test(nfa75, [nondet]) :- reachable(20,21,[a,b,a]).
test(nfa76, [nondet]) :- reachable(21,21,[a,a]).
test(nfa77, [fail])  :- reachable(20,21,[c]).  
test(nfa78, [nondet]) :- reachable(20,20,[a,b,b,a,b,a]).test(nfa79, [fail])  :- reachable(20,21,[b,b,b,a,b]).
test(nfa80, [fail])  :- reachable(21,20,[a,c,b]).

:- end_tests(nfa).
