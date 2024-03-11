# go-roe

Golang pattern to return Result or Error useful 
when array of values need to be returned from the function.

Inspired from [Rust Result unwrap](https://doc.rust-lang.org/rust-by-example/error/result.html)

# Usage

**Simple**

```
import (
    "github.com/prasad83/goroe"
)

func helloWorld() ResultOrError[string] {
    return goroe.Result("hello world")
}

func main() {
    if result, err := helloWorld().Unwrap(); err != nil {
        panic(err)
    } else {
        println(result)
    }
}
```

**Complex**

```
import (
    "github.com/prasad83/goroe"
)

func helloPlanets(planets ...string) []ResultOrError[string] {
	returns := make([]ResultOrError, len(planets))
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
			returns[index] = Result[string](fmt.Sprintf("Hello %s", planet))
		default:
			returns[index] = Errorf[string]("Not a planet")
		}
	}
	return returns
}

func main() {
	results := helloPlanets("Earth", "Moon")
	for _, roe := range results {
		if r, err := roe.Unwrap(); err != nil {
			println("Error: " + err.Error())
		} else {
			println(r)
		}
	}
}
```