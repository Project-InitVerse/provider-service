package operator

import (
	"github.com/spf13/cobra"

	"providerService/src/operator/inventory"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "operator",
		Short:        "kubernetes operators control",
		SilenceUsage: true,
	}

	cmd.AddCommand(inventory.Cmd())

	return cmd
}
