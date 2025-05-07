; You may define helper functions here
(defun reachable-list (transition states final input)
  (cond
    ((null states) nil)
    ((reachable
       transition
       (car states)
       final
       input)
     t)
    (t
     (reachable-list
       transition
       (cdr states)
       final
       input))))
       
(defun reachable (transition start final input)
  (cond
    ((null input) (equal start final))
    (t
     (reachable-list
       transition
       (funcall transition start (car input))
       final
       (cdr input)))))

