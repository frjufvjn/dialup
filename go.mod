module github.com/frjufvjn/dialup

go 1.19

replace github.com/frjufvjn/dialup/tcp => ./tcp

replace github.com/frjufvjn/dialup/udp => ./udp

require (
	github.com/frjufvjn/dialup/tcp v0.0.0-00010101000000-000000000000
	github.com/frjufvjn/dialup/udp v0.0.0-00010101000000-000000000000
)
