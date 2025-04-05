package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	// set through ldflags
	version string
	date    string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints version information",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Version:    ", version)
		fmt.Println("Date:       ", date)
		fmt.Println("Go Version: ", runtime.Version())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
