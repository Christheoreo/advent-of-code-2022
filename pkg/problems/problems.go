package problems

type Problem interface {
	Solve(filename string) error
	SolveFirst(filename string) error
	SolveSecond(filename string) error
}
