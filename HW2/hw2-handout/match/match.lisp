(defun match (pattern assertion)
  (cond
    ((null pattern)
     (null assertion))

    ((eq (car pattern) '?)
     (and (consp assertion)
          (match (cdr pattern) (cdr assertion))))

    ((eq (car pattern) '!)
     (and (consp assertion)
          (or (match (cdr pattern) (cdr assertion))
              (match pattern (cdr assertion)))))

    ((and (consp assertion)
          (equal (car pattern) (car assertion)))
     (match (cdr pattern) (cdr assertion)))

    (t nil)))
