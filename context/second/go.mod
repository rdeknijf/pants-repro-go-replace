module something.org/second

go 1.23

replace something.org/first => ../first

require (
	gotest.tools v2.2.0+incompatible
	something.org/first v0.0.0-00010101000000-000000000000
)

require (
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
)
