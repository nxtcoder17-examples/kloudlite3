package del

import (
	"fmt"
	"github.com/kloudlite/kl/domain/client"
	common_util "github.com/kloudlite/kl/pkg/functions"

	"github.com/kloudlite/kl/constants"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/spf13/cobra"
)

var deleteConfigCommand = &cobra.Command{
	Use:   "config",
	Short: "remove one config environment from your " + constants.CmdName + "-config",
	Long: `This command help you to delete environment that that is comming from config

Examples:
  # remove config
  kl del config
	`,
	Run: func(_ *cobra.Command, _ []string) {
		err := removeConfig()
		if err != nil {
			common_util.PrintError(err)
			return
		}
	},
}

func removeConfig() error {

	klFile, err := client.GetKlFile(nil)
	if err != nil {
		common_util.PrintError(err)
		es := "please run '" + constants.CmdName + " init' if you are not initialized the file already"
		common_util.PrintError(fmt.Errorf(es))
	}

	if len(klFile.Configs) == 0 {
		es := "no configs added yet in your file"
		return fmt.Errorf(es)
	}

	selectedConfigIndex, err := fuzzyfinder.Find(
		klFile.Configs,
		func(i int) string {
			return klFile.Configs[i].Name
		},
		fuzzyfinder.WithPromptString("Select Config Group >"),
	)

	if err != nil {
		return err
	}

	selectedConfig := klFile.Configs[selectedConfigIndex]

	if len(selectedConfig.Env) == 1 {
		newConfigs := make([]client.ResType, 0)
		for i, rt := range klFile.Configs {
			if i == selectedConfigIndex {
				continue
			}
			newConfigs = append(newConfigs, rt)
		}

		klFile.Configs = newConfigs

		fmt.Printf("removed config %s form your %s-file\n", selectedConfig.Name, constants.CmdName)

	} else {

		selectedKeyIndex, e := fuzzyfinder.Find(
			selectedConfig.Env,
			func(i int) string {
				return fmt.Sprintf("%s, %s", selectedConfig.Env[i].Key, selectedConfig.Env[i].RefKey)
			},
			fuzzyfinder.WithPromptString("Select Config Key >"),
		)

		if e != nil {
			return e
		}

		newEnvs := make([]client.ResEnvType, 0)
		for i, ret := range selectedConfig.Env {
			if i == selectedKeyIndex {
				continue
			}
			newEnvs = append(newEnvs, ret)
		}

		klFile.Configs[selectedConfigIndex].Env = newEnvs

		fmt.Printf("removed key %s/%s form your %s-file\n", selectedConfig.Name, selectedConfig.Name, constants.CmdName)
	}

	err = client.WriteKLFile(*klFile)

	return err
}
