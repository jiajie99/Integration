package mockery

// nolint:unused
//
//go:generate mockery --name=calculation --with-expecter --inpackage --filename=mock_calculation.go
type calculation interface {
	Add(a, b int) int
}
