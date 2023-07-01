package gomock

//go:generate mockgen -source=calculation.go -destination=mock_calculation.go -package=gomock
type calculation interface {
	add(a, b int) int
	addFunc(a, b int, f func(a, b int) int) int
}

type example struct {
	c calculation
}

func (e *example) executeAdd(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return e.c.addFunc(a, b, f)
}
