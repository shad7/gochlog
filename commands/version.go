package commands

import (
	"fmt"

	// The fork has multiple enhancements being used by Docker but not merged
	// into the base project as of yet. Once these are included then the switch
	// back can be made. This project also leverages tagged versions/releases
	// which make stability a strong option.
	// "github.com/spf13/cobra"
	"github.com/dnephin/cobra"

	"github.com/shad7/gochlog/core"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Long:  "Print version",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print(core.GetVersionDisplay())
		return nil
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
