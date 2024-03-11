package goroe

import (
	"fmt"
	"strings"
	"testing"
)

func helloWorld() ResultOrError[string] {
	return Result("hello world")
}

func helloPlanets(planets ...string) []ResultOrError[string] {
	returns := make([]ResultOrError[string], len(planets))
	for index, planet := range planets {
		switch strings.ToUpper(planet) {
		case "MERCURY":
			fallthrough
		case "VENUS":
			fallthrough
		case "EARTH":
			fallthrough
		case "MARS":
			fallthrough
		case "JUPITER":
			fallthrough
		case "SATURN":
			fallthrough
		case "NEPTUNE":
			fallthrough
		case "PLUTO":
			returns[index] = Result(fmt.Sprintf("Hello %s", planet))
		default:
			returns[index] = Errorf[string]("Not a planet")
		}
	}
	return returns
}

func TestHelloWorld(t *testing.T) {
	if result, err := helloWorld().Unwrap(); err != nil {
		panic(err)
	} else {
		println(result)
	}
}

func TestHelloPlanets(t *testing.T) {
	results := helloPlanets("Earth", "Moon")
	for _, roe := range results {
		if r, err := roe.Unwrap(); err != nil {
			println("Error: " + err.Error())
		} else {
			println(r)
		}
	}
}
