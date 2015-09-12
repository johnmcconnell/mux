# mux
Generic mux using a got file [gotgo](https://github.com/droundy/gotgo)

Mux is short for multiplexer, it takes an input to decide which handler
to use and then handles the message for a given input and output.

## Usage
Create your own
```
install [gotgo](https://github.com/droundy/gotgo)

gotgo -pack-name="yourpackage" path/to/template > path/to/out type1
type2 type 3

gotgo -package-name="intmux" template.got > intmux/intmux.go int int *int
```

## Example
```go
type Handler1 struct{}

func (t Handler1) Handle(x int, y *int) {
  *y = x
}

type Handler2 struct{}

func (t Handler2) Handle(x int, y *int) {
  *y = (2 * x)
}

m := qikmux.New(nil) 

m.HandleWith(
  1,
  testHandler1{},
)

m.HandleWith(
  2,
  testHandler2{},
)

r := 0

m.Select(1, 2, &r)
r //= 2

m.Select(2, 5, &r)
r //= 10
```
