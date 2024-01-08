package vpn

import (
	"github.com/kloudlite/kl/cmd/vpn/intercept"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "vpn",
	Short: "work with vpn",
	Long: `work with vpn
Examples:
	# start vpn
  kl vpn start

	# stop vpn
	kl vpn stop

	# restart vpn
	kl vpn restart

	# status vpn
	kl vpn status

	# list all vpn
	kl vpn list

	# switch to vpn
	kl vpn switch <vpn_name>

	# remove vpn
	kl vpn remove <vpn_name>
	`,
}

func init() {
	Cmd.Aliases = append(Cmd.Aliases, "dev")

	Cmd.AddCommand(newCmd)
	Cmd.AddCommand(listCmd)
	Cmd.AddCommand(switchCmd)
	Cmd.AddCommand(restartCmd)
	Cmd.AddCommand(startCmd)
	Cmd.AddCommand(stopCmd)
	Cmd.AddCommand(statusCmd)
	Cmd.AddCommand(activateCmd)
	Cmd.AddCommand(intercept.Cmd)
}
