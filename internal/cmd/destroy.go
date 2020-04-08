package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/craftcms/nitro/config"
	"github.com/craftcms/nitro/internal/action"
)

var (
	flagPermanent bool

	destroyCommand = &cobra.Command{
		Use:   "destroy",
		Short: "Destroy a machine",
		RunE: func(cmd *cobra.Command, args []string) error {
			name := config.GetString("name", flagMachineName)

			if flagPermanent {
				fmt.Println("permanently deleting", name)
			} else {
				fmt.Println("soft deleting", name)
			}

			destroyAction, err := action.Destroy(name, flagPermanent)
			if err != nil {
				return err
			}

			return Run(NewMultipassRunner("multipass"), []action.Action{*destroyAction})
		},
	}
)

func init() {
	destroyCommand.Flags().BoolVarP(&flagPermanent, "permanent", "p", false, "permanently destroy the machine")
}
