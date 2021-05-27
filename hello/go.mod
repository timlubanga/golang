module example.com/hello

go 1.16

replace example.com/greetings => ../greetings

require (
	example.com/calculate v0.0.0-00010101000000-000000000000
	example.com/greetings v0.0.0-00010101000000-000000000000
)

replace example.com/calculate => ../calculate
