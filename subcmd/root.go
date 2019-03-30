package subcmd

import (
	"github.com/spf13/cobra"
)

var Version string

func init() {
	cobra.OnInitialize()
}

var RootCommand = &cobra.Command{
	Use:     "align",
	Short:   "align is text align command",
	Example: "align right -n 10 a",
	Version: Version,
	Long: `
align is text align command.

Repository: https://github.com/jiro4989/align
    Author: jiro4989
	`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}
