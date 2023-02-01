package flags

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//Define k8s default config path
const (
	KubeConfigDefaultPath = "$HOME/.kube/config"
)

//AddKubeConfigPathFlag is function add config flag
func AddKubeConfigPathFlag(cmd *cobra.Command) error {
	cmd.Flags().String(FlagKubeConfig, "$HOME/.kube/config", "kubernetes configuration file path")
	return viper.BindPFlag(FlagKubeConfig, cmd.Flags().Lookup(FlagKubeConfig))
}
