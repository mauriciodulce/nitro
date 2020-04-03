package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/craftcms/nitro/command"
	"github.com/craftcms/nitro/internal/nitro"
)

func init() {
	siteCommand.AddCommand(siteAddCommand)
}

var siteAddCommand = &cobra.Command{
	Use:   "add",
	Short: "Add a site to machine",
	Run: func(cmd *cobra.Command, args []string) {
		if err := nitro.Run(
			command.NewMultipassRunner("multipass"),
			nitro.Empty(flagMachineName),
		); err != nil {
			log.Fatal(err)
		}
	},
}
