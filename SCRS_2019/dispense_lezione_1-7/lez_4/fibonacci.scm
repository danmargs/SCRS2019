;readline is keg-only, which means it was not symlinked into /usr/local,
;because macOS provides the BSD libedit library, which shadows libreadline.
;In order to prevent conflicts when programs look for libreadline we are
;defaulting this GNU Readline installation to keg-only.
;
;For compilers to find readline you may need to set:
;  export LDFLAGS="-L/usr/local/opt/readline/lib"
;  export CPPFLAGS="-I/usr/local/opt/readline/include"
;
;For pkg-config to find readline you may need to set:
;  export PKG_CONFIG_PATH="/usr/local/opt/readline/lib/pkgconfig"

(define (fibo n)
    (cond
        ((= n 0) 0)
        ((= n 1) 1)
        ( (+ (fibo (- n 1)) (fibo (- n 2))))
    )
)

(begin
    (display "Sequenza Fibonacci:")
    (define a (fibo 50))
    (newline)
    (display a)
    (newline)
)


