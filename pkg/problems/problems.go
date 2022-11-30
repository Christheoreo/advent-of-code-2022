package problems

type Problem interface {
	Solve() error
	SolveFirst() error
	SolveSecond() error
}
