package data

// Generator has a Generate and a Reset for managing
// test data
type Generator interface {
	Generate() error
	Reset() error
}
