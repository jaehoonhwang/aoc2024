package day

import (
	utils "github.com/jaehoonhwang/aoc2024/utils"
)

var (
	p1_filepath = "day1/part1.txt"
	p2_filepath = "day1/part2.txt"
)

type Day1 utils.Problem

func NewDay1() *Day1 {
	ret := Day1{}
	return &ret
}

func (d *Day1) Part1() string {
	return "solution for part 1"
}

func (d *Day1) Part2() string {
	return "solution for part 2"
}
