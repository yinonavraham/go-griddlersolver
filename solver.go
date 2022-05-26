package griddlersolver

func Solve(p Problem) Solution {
	return &solution{
		problem: p,
	}
}