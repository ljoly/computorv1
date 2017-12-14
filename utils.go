package main

type ErrorMsg int

const (
	noArgument ErrorMsg = iota
	tooManyArguments
	wrongFormat
	degreeTooHigh
	degreeNegative
	degreeNotInt
	coefNotInt
)

const formatRegex = `(?m)([+-])(?:(?:(\d+(?:\.\d+)?)(?:(?:\*X(?:\^(?:[+-])?(\d+))?)|X(?:\^(?:[+-])?(\d+))?))|(\d+(?:\.\d+)?)|(?:(?:\*X(?:\^(?:[+-])?(\d+))?)|X(?:\^(?:[+-])?(\d+))?))`
const negativeDegreeRegex = `(?m)(?:\^(?:\-))`
const degreeRegex = `(?m)(?:\^)`

type Env struct {
	left, right string
	degree      int64
	a, b, c     float64
}
