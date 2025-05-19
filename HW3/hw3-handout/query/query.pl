/* All novels published during the year 1953 or 1996*/
year_1953_1996_novels(Book) :- novel(Book, 1953);novel(Book, 1996).
    %% remove fail and add body/other cases for this predicate


/* List of all novels published during the period 1800 to 1900*/
period_1800_1900_novels(Book) :- novel(Book, Year), Year >= 1800, Year =< 1900.
    %% remove fail and add body/other cases for this predicate


/* Characters who are fans of LOTR */
lotr_fans(Fan) :-fan(Fan, Books), member(the_lord_of_the_rings, Books).
    %% remove fail and add body/other cases for this predicate

/* Authors of the novels owned by Chris */
author_names(Author) :- fan(chris,  LikedBooks), member(Oneof,  LikedBooks), author(Author,WrittenBooks), member(Oneof,WrittenBooks).

    %% remove fail and add body/other cases for this predicate

/* Characters who are fans of Brandon Sanderson's novels */
fans_names(Fan) :- author(brandon_sanderson, WrittenBooks), member(Oneof, WrittenBooks), fan(Fan, LikedBooks), member(Oneof, LikedBooks).
    % author(Brandon_Sanderson, Books), fan(Fan, FanBooks), member(Books, FanBooks).
    %% remove fail and add body/other cases for this predicate

/* Novels common between either of Alex, Logan, and Charlotte */
mutual_novels(Book) :- fan(alex, ALikedBooks), member(Book, ALikedBooks), fan(logan, LLikedBooks), member(Book, LLikedBooks);
                       fan(alex, ALikedBooks), member(Book, ALikedBooks), fan(charlotte, CLikedBooks), member(Book, CLikedBooks);
                       fan(logan, LLikedBooks), member(Book, LLikedBooks), fan(charlotte, CLikedBooks), member(Book, CLikedBooks).

