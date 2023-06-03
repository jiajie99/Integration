package gomock

//go:generate mockgen -source=calculation.go -destination=mock_calculation.go -package=gomock
type Calculation interface {
	Add(a, b int) int
}
