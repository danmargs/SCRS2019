    (let (
            ;lista associazioni
            (v1 10)
            (alfa 20)
            (elenco '(1 2 3))
            (resto "Buongiorno")
        )
        ; le seguenti operazioni utilizzano quei valori
        (list resto (+ v1 alfa) (cadr elenco))
        ;operatore list crea lista
    )

    ;scrivere una funzione che torna l'n-simo elemento di una lista
    (define (nth l n)
        ;lista ha n elementi?
        (if (nil? l)
            #nil
            (if (= n 0)
                (car l)
                (nth (cdr l) (- n 1))
            )
        )
    )

    ;accumulatore
    (define (adder)
        (let (
            (tot 0)
            )
            ;valore di ritorno
            (lambda (cmd . args) 
                (cond 
                ;uguaglianza eqv? se command è add allora tot associa la somma di tot col primo dei parametri
                    ((eqv? cmd 'add)
                    (set! tot (+ tot (car args)))
                    )
                    ((eqv? cmd 'sub)
                    (set! tot (- tot (car args)))
                    )
                    ;per altri comandi torni tot
                    (#t tot)
                )
            )
        )
    )

#|  ;questo da errore perchè cerchi di utilizzare v1 in fase di inizializzazione
    (let ((v1 10)
        (v2 (1+ v1))
        )
        (display v1) (newline)
        ) |#

    ;questo non dà errore
    (let* ((v1 10)
        (v2 (1+ v1))
        )
        (newline)
        (display "Valore let:")
        (display v1) 
        (newline)
    )

    ;loop fatto con la let (è meglio)
    (let loop ((v1 10) (v2 20))
        ;voglio calcolare questo
        (if (and (> v1 0) (> v2 0))
            (loop (- v1 1) (* v2 v1))
            1
        )
    )
    ;funzione ricorsiva che si chiama loop
    ;se voglio fare for
    (display "ciclo for:")
    (newline)
    (let ciclofor ((n 10))
        (if (zero? n)
            'fine
            (begin
                (display n)
                (newline)
                (ciclofor (- n 1))
            )
        )
    )

    ;si può fare con un altra funzione passata (lambda (i) (* i i)) mettendo (Squarel l fun)
    (define (Squarel l)
        (cond
            ((nil? l) #nil)
            (#t (cons 
                    (* (car l) (car l) )
                    (Squarel (cdr l))
                )
            )
        )
    )

;map può essere eseguita in parallelo
;torna una lista da 0 a 9 con i quadrati
(map (lambda (i) (* i i)) ((iota 10)))
;ci sta anche reduce
(define (Addl l)
    (cond 
        ((nil? l) 0)
        (#t (+ (car l) (Addl (cdr l))))
    )
)
;dovrebbe funzionare, ma non mi va di provarla, somma i numeri da 0 a n
;più carina della reduce c'è la fold (compatta, fa la stessa cosa)
(fold (lambda (curr prev) (+ curr prev)) 0n (iota 10))
;per ogni elemento della lista chiama lambda e il valore di ritorno dienta il prev della chiamata siccessiva, dando la somma dei numeri da 1 a n
;leggermente più veloce di quella sopra perchè implementata internamente


;una specie di main
(begin

    (display "lista quadrati :")
    (define c (Squarel '(1 2 3 4 5)))
    (display c)
    (newline)

    (display "lista :")
    (define a (nth '(1 2 3 4 5) 2))
    (newline)
    (display a)
    (newline)

    (display "adder :")
    (define b (adder) )
    (b 'add 3)
    (b 'add 2)
    (newline)
    (display (b 'altroComando))
    (newline)

)
