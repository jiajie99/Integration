package mockery

//go:generate mockery --name=Calculation --with-expecter --inpackage --filename=mock_calculation.go
type Calculation interface {
	Add(a, b int) int
}
