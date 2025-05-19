:- initialization main.

main :-
    consult(["facts.pl", "query.pl"]),
    (show_coverage(run_tests) ; true),
    halt.

:- begin_tests(query).

test(year_1953_1996_novels, set(X == [
        a_song_of_ice_and_fire_series,
        childhoods_end,
        fahrenheit451,
        neverwhere,
        the_caves_of_steel
        ])) :-
    year_1953_1996_novels(X).

test(period_1800_1900_novels, set(X == [
        frankenstein,
        little_women,
        the_20000_leagues_under_the_sea,
        the_journey_to_the_center_of_the_earth,
        the_time_machine,
        the_war_of_the_worlds
        ])) :-
    period_1800_1900_novels(X).

test(lotr_fans, set(X == [
        henry,
        bob,
        charlotte,
        susan,
        sasha
        ])) :-
    lotr_fans(X).

test(author_names, set(X == [
        arthur_clarke,
        george_rr_martin,
        isaac_asimov,
        robert_heinlein,
        william_gibson
        ])) :-
    author_names(X).

test(fans_names, set(X == [
        henry,
        mia,
        carlos,
        charlotte,
        maria,
        adam
        ])) :-
    fans_names(X).

test(mutual_novels, set(X == [
        something_wicked_this_way_comes,
        the_princess_bride,
        the_time_machine,
        the_wheel_of_time_series
        ])) :-
    mutual_novels(X).

test(novel_frankenstein_year_1818, true(Year == 1818)) :-
    novel(frankenstein, Year).

test(novel_20000_leagues_year_1870, true(Year == 1870)) :-
    novel(the_20000_leagues_under_the_sea, Year).

test(fan_bob_likes_lotr, true(member(the_lord_of_the_rings, Books))) :-
    fan(bob, Books).

test(fan_mia_likes_mistborn, true(member(mistborn_trilogy, Books))) :-
    fan(mia, Books).

test(author_asimov_works, set(Works == [[the_foundation_trilogy, i_robot, the_caves_of_steel]])) :-
    author(isaac_asimov, Works).

test(author_wells_works, set(Works == [[the_time_machine, the_war_of_the_worlds]])) :-
    author(hg_wells, Works).

test(year_1953_includes_fahrenheit451, true) :-
    year_1953_1996_novels(fahrenheit451).

test(year_1953_includes_caves, true) :-
    year_1953_1996_novels(the_caves_of_steel).

test(year_1996_includes_asoiaf, true) :-
    year_1953_1996_novels(a_song_of_ice_and_fire_series).

test(year_1953_excludes_lotr, [fail]) :-
    year_1953_1996_novels(the_lord_of_the_rings).

test(period_1800_1900_includes_war_of_worlds, true) :-
    period_1800_1900_novels(the_war_of_the_worlds).

test(period_1800_1900_excludes_dune, [fail]) :-
    period_1800_1900_novels(the_dune_chronicles).

test(lotr_fans_includes_henry, true) :-
    lotr_fans(henry).

test(lotr_fans_excludes_mia, [fail]) :-
    lotr_fans(mia).

test(fans_names_includes_adam, true) :-
    fans_names(adam).

test(fans_names_excludes_jon, [fail]) :-
    fans_names(jon).

test(author_names_includes_clarke, true) :-
    author_names(arthur_clarke).

test(author_names_excludes_king, [fail]) :-
    author_names(stephen_king).

test(mutual_novels_includes_time_machine, true) :-
    mutual_novels(the_time_machine).

test(mutual_novels_excludes_dune, [fail]) :-
    mutual_novels(the_dune_chronicles).

:- end_tests(query).
