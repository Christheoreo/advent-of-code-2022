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
# test-2:
# 	go test -v github.com/christheoreo/advent-of-code-2022/day02
# bench-2:
# 	go test github.com/christheoreo/advent-of-code-2022/day02 -bench=. -benchmem -run=^$