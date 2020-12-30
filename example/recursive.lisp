(defun rcsv (n)
    (
        if
        (<= n 1)
        1
        (+ (rcsv (- n 1)) n)
    )

)
(rcsv 4)
