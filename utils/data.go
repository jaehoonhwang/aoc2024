package utils

type Solution interface {
	Part1(string) string
	Part2(string) string
}

type Problem struct {
	Part1Input []string
	Part2Input []string
}
