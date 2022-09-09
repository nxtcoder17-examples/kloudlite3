package wg

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/kloudlite/kl/lib/common"
	"github.com/kloudlite/kl/lib/common/ui/color"
	"github.com/spf13/cobra"
)

func startServiceInBg() {
	command := exec.Command("kl", "wg", "connect", "--foreground")
	err := command.Start()
	if err != nil {
		fmt.Println(err)
		return
	}
	configFolder, err := common.GetConfigFolder()
	if err != nil {
		common.PrintError(err)
		return
	}
	os.WriteFile(configFolder+"/wgpid", []byte(fmt.Sprintf("%d", command.Process.Pid)), 0644)
}

var foreground bool
var connectVerbose bool

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "connect selected device wireguard",
	Long: `connect wireguard
Examples:
  # connect the selected device
  kl wg connect

  # connect the selected device with verbose
  kl wg connect -v

	`,
	Run: func(_ *cobra.Command, _ []string) {
		if euid := os.Geteuid(); euid != 0 {
			common.PrintError(
				errors.New(
					color.ColorText("make sure you are running command with sudo", 209),
				),
			)
			return
		}

		if foreground {
			if err := startService(connectVerbose); err != nil {
				common.PrintError(err)
				return
			}
		} else {
			startServiceInBg()
			if err := startConfiguration(connectVerbose); err != nil {
				common.PrintError(err)
				return
			}
		}

		common.PrintError(errors.New("[#] connected"))
	},
}

func init() {
	connectCmd.Flags().BoolVar(&foreground, "foreground", false, "")
	connectCmd.Flags().BoolVarP(&connectVerbose, "verbose", "v", false, "show verbose")
}
