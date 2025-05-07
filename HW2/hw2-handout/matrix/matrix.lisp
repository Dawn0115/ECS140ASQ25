; You may define helper functions here

(defun row_adder (r1 r2)
  (cond
    ((null r1) nil)
    (t
     (cons (+ (car r1) (car r2))
           (row_adder (cdr r1) (cdr r2))))))

(defun matrix-add (A B)
  (cond
    ((null A) nil)
    (t
     (cons (row_adder (car A) (car B))
           (matrix-add (cdr A) (cdr B))))))

(defun first_col_remover (M)
  (cond
    ((null M) nil)
    (t
     (cons (car (car M)) ;; first column
           (first_col_remover (cdr M))))))

(defun rest_cols (M)
  (cond
    ((null M) nil)
    (t
     (cons (cdr (car M)) ;; rest of the first row
           (rest_cols (cdr M))))))

(defun matrix-transpose (M)
  (cond
    ((null (car M)) nil)
    (t
     (cons (first_col_remover M)
           (matrix-transpose (rest_cols M))))))

(defun product (a b)
  (if (null a)
      0
    (+ (* (car a) (car b)) ;; return the product of two vector
       (product (cdr a) (cdr b)))))

(defun row_mult (row B_transpose) ;; multiple of row of A and first "column" of  B
  (if (null B_transpose)
      nil
    (cons
      (product row (car B_transpose))
      (row_mult row (cdr B_transpose)))))

(defun matrix-multiply (A B)
  (if (null A)
      nil
    (let ((B_transpose (matrix-transpose B)))
      (cons
        (row_mult (car A) B_transpose)
        (matrix-multiply (cdr A) B)))))
