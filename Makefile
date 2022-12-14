test:
	go test ./...
test-v:
	go test -v ./...
bench:
	go test ./... -bench=. -benchmem -run=^$
test-1:
	go test -v github.com/christheoreo/advent-of-code-2022/day01
bench-1:
	go test github.com/christheoreo/advent-of-code-2022/day01 -bench=. -benchmem -run=^$
test-2:
	go test -v github.com/christheoreo/advent-of-code-2022/day02
bench-2:
	go test github.com/christheoreo/advent-of-code-2022/day02 -bench=. -benchmem -run=^$
test-3:
	go test -v github.com/christheoreo/advent-of-code-2022/day03
bench-3:
	go test github.com/christheoreo/advent-of-code-2022/day03 -bench=. -benchmem -run=^$
test-4:
	go test -v github.com/christheoreo/advent-of-code-2022/day04
bench-4:
	go test github.com/christheoreo/advent-of-code-2022/day04 -bench=. -benchmem -run=^$
test-5:
	go test -v github.com/christheoreo/advent-of-code-2022/day05 -count=1 # remove cache
bench-5:
	go test github.com/christheoreo/advent-of-code-2022/day05 -bench=. -benchmem -run=^$
test-6:
	go test -v github.com/christheoreo/advent-of-code-2022/day06 -count=1 # remove cache
bench-6:
	go test github.com/christheoreo/advent-of-code-2022/day06 -bench=. -benchmem -run=^$
test-7:
	go test -v github.com/christheoreo/advent-of-code-2022/day07 -count=1 # remove cache
bench-7:
	go test github.com/christheoreo/advent-of-code-2022/day07 -bench=. -benchmem -run=^$
test-8:
	go test -v github.com/christheoreo/advent-of-code-2022/day08 -count=1 # remove cache
bench-8:
	go test github.com/christheoreo/advent-of-code-2022/day08 -bench=. -benchmem -run=^$
test-9:
	go test -v github.com/christheoreo/advent-of-code-2022/day09 -count=1 # remove cache
bench-9:
	go test github.com/christheoreo/advent-of-code-2022/day09 -bench=. -benchmem -run=^$
test-10:
	go test -v github.com/christheoreo/advent-of-code-2022/day10 -count=1 # remove cache
bench-10:
	go test github.com/christheoreo/advent-of-code-2022/day10 -bench=. -benchmem -run=^$