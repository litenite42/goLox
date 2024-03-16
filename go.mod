module golox

go 1.22.1

replace golox/loxerr => ./loxerr

require golox/core v0.0.0-00010101000000-000000000000

require (
	golox/loxerr v0.0.0-00010101000000-000000000000 // indirect
	golox/tern v0.0.0-00010101000000-000000000000 // indirect
)

replace golox/core => ./core

replace golox/tern => ./tern
