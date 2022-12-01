package cmd

import (
	"fmt"
	"os"

	"github.com/christheoreo/advent-of-code-2022/pkg/problems"
	"github.com/spf13/cobra"
)

type Mapping struct {
	files   [2]string
	problem problems.Problem
}

var problemsMap map[int]Mapping = map[int]Mapping{
	1: {
		problem: problems.ProblemOne{},
		files:   [2]string{"problem-1.txt", "problem-1-example.txt"},
	},
	2: {
		problem: problems.ProblemTwo{},
		files:   [2]string{"problem-2.txt", "problem-2-example.txt"},
	},
}

func Execute() {

	var rootCmd = &cobra.Command{Use: "cli"}
	rootCmd.AddCommand(cmdSolve)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
