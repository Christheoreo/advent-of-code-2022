package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/christheoreo/advent-of-code-2022/internal/timetrack"
	"github.com/spf13/cobra"
)

var partString string = "all"

var cmdSolve = &cobra.Command{
	Use:   "solve [problem number] [part]",
	Short: "Solve a specific problem",
	Long:  `solve is for executing a specific problem.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if partString != "all" && partString != "1" && partString != "2" {
			partString = "all"
		}

		problemNumber, err := strconv.Atoi(args[0])

		if err != nil {
			fmt.Printf("Error - Could not convet %s to a number\n", args[0])
			return
		}

		problem, ok := problemsMap[problemNumber]

		if !ok {
			fmt.Println("Error - could not find problem to solve")
			return
		}

		if partString == "all" || partString == "one" {
			firstTime := time.Now()
			err = problem.SolveFirst()

			timetrack.TimeTrack(firstTime, fmt.Sprintf("%d - part 1", problemNumber))
			if err != nil {
				panic(err)
			}
		}

		if partString == "one" {
			return
		}

		secondTime := time.Now()
		err = problem.SolveSecond()
		timetrack.TimeTrack(secondTime, fmt.Sprintf("%d - part 2", problemNumber))
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	cmdSolve.Flags().StringVarP(&partString, "part", "p", "", "Problem part")
}
