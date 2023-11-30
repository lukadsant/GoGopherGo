module helloworld

go 1.21.4
require (
    github.com/lukadsant/calculator v0.0.0
    go mod download rsc.io/sampler
)

replace github.com/lukadsant/calculator => ../calculator
