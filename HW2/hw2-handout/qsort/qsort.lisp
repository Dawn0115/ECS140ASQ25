; You may define helper functions here

(defun pivot (n xs)
  (cond
    ((null xs) (list nil nil))
    (t
     (let* ((x      (car xs))
            (rest-p (pivot n (cdr xs)))
            (lt     (car rest-p))
            (ge     (car (cdr rest-p))))
       (if (< x n)
           (list (cons x lt) ge)
           (list lt (cons x ge)))))))

(defun quicksort (xs)
  (cond
    ((null xs) nil)
    (t
     (let* ((p     (car xs))
            (parts (pivot p (cdr xs)))
            (lt    (car parts))
            (ge    (cadr parts)))
       (append (quicksort lt)
               (list p)
               (quicksort ge))))))