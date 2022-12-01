package cmd

import (
	"fmt"
	"os"

	"github.com/christheoreo/advent-of-code-2022/pkg/problems"
	"github.com/spf13/cobra"
)

var problemsMap map[int]problems.Problem = map[int]problems.Problem{
	1: problems.ProblemOne{},
	2: problems.ProblemTwo{},
}

func Execute() {

	var rootCmd = &cobra.Command{Use: "cli"}
	rootCmd.AddCommand(cmdSolve)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
