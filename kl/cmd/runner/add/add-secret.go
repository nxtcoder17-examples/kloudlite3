package add

import (
	"errors"
	"fmt"
	"strings"

	"github.com/kloudlite/kl/domain/client"
	"github.com/kloudlite/kl/domain/server"
	fn "github.com/kloudlite/kl/pkg/functions"
	"github.com/kloudlite/kl/pkg/ui/fzf"

	"github.com/spf13/cobra"
)

var secCmd = &cobra.Command{
	Use:   "secret [name]",
	Short: "add secret references to your kl-config",
	Long:  `This command will add secret entry references from current environement to your kl-config file.`,
	Example: `
  kl add secret 		# add secret and entry by selecting from list (default)
  kl add secret [name] 	# add entry by providing secret name
	`,
	Run: func(cmd *cobra.Command, args []string) {
		err := selectAndAddSecret(cmd, args)
		if err != nil {
			fn.PrintError(err)
			return
		}
	},
}

func selectAndAddSecret(cmd *cobra.Command, args []string) error {
	//TODO: add changes to the klbox-hash file
	// m := fn.ParseStringFlag(cmd, "map")
	filePath := fn.ParseKlFile(cmd)

	name := ""
	if len(args) >= 1 {
		name = args[0]
	}

	klFile, err := client.GetKlFile(filePath)
	if err != nil {
		fn.PrintError(err)
		es := "please run 'kl init' if you are not initialized the file already"
		return fmt.Errorf(es)
	}

	secrets, err := server.ListSecrets()
	if err != nil {
		return err
	}

	if len(secrets) == 0 {
		return fmt.Errorf("no secrets created yet on server")
	}

	selectedSecretGroup := server.Secret{}

	if name != "" {
		for _, c := range secrets {
			if c.Metadata.Name == name {
				selectedSecretGroup = c
				break
			}
		}
		return errors.New("can't find secrets with provided name")

	} else {
		selectedGroup, err := fzf.FindOne(
			secrets,
			func(item server.Secret) string {
				return item.Metadata.Name
			},
			fzf.WithPrompt("Select Secret Group >"),
		)
		if err != nil {
			return err
		}

		selectedSecretGroup = *selectedGroup
	}

	if len(selectedSecretGroup.StringData) == 0 {
		return fmt.Errorf("no secrets added yet to %s secret", selectedSecretGroup.Metadata.Name)
	}

	type KV struct {
		Key   string
		Value string
	}

	selectedSecretKey := &KV{}

	m := ""

	if m != "" {
		kk := strings.Split(m, "=")
		if len(kk) != 2 {
			return errors.New("map must be in format of secret_key=your_var_key")
		}

		for k, v := range selectedSecretGroup.StringData {
			if k == kk[0] {
				selectedSecretKey = &KV{
					Key:   k,
					Value: v,
				}
				break
			}
		}

		return errors.New("secret_key not found in selected secret")

	} else {
		selectedSecretKey, err = fzf.FindOne(
			func() []KV {
				var kvs []KV

				for k, v := range selectedSecretGroup.StringData {
					kvs = append(kvs, KV{
						Key:   k,
						Value: v,
					})
				}

				return kvs
			}(),
			func(val KV) string {
				return val.Key
			},
			fzf.WithPrompt(fmt.Sprintf("Select Key of %s >", selectedSecretGroup.Metadata.Name)),
		)
		if err != nil {
			return err
		}
	}

	currSecs := klFile.EnvVars.GetSecrets()

	matchedGroupIndex := -1
	for i, rt := range currSecs {
		if rt.Name == selectedSecretGroup.Metadata.Name {
			matchedGroupIndex = i
			break
		}
	}

	if matchedGroupIndex != -1 {
		matchedKeyIndex := -1

		for i, ret := range currSecs[matchedGroupIndex].Env {
			if ret.RefKey == selectedSecretKey.Key {
				matchedKeyIndex = i
				break
			}
		}

		if matchedKeyIndex == -1 {
			currSecs[matchedGroupIndex].Env = append(currSecs[matchedGroupIndex].Env, client.ResEnvType{
				Key: RenameKey(func() string {
					if m != "" {
						kk := strings.Split(m, "=")
						return kk[1]
					}
					return selectedSecretKey.Key
				}()),
				RefKey: selectedSecretKey.Key,
			})
		}
	} else {
		currSecs = append(currSecs, client.ResType{
			Name: selectedSecretGroup.Metadata.Name,
			Env: []client.ResEnvType{
				{
					Key: RenameKey(func() string {
						if m != "" {
							kk := strings.Split(m, "=")
							return kk[1]
						}
						return selectedSecretKey.Key
					}()),
					RefKey: selectedSecretKey.Key,
				},
			},
		})

	}

	klFile.EnvVars.AddResTypes(currSecs, client.Res_secret)
	err = client.WriteKLFile(*klFile)
	if err != nil {
		return err
	}

	fn.Log(fmt.Sprintf("added secret %s/%s to your kl-file\n", selectedSecretGroup.Metadata.Name, selectedSecretKey.Key))

	if err := server.SyncBoxHash(); err != nil {
		return err
	}

	//if err := server.SyncDevboxJsonFile(); err != nil {
	//	return err
	//}
	//
	//if err := client.SyncDevboxShellEnvFile(cmd); err != nil {
	//	return err
	//}
	return nil
}

func init() {
	// secCmd.Flags().StringP("map", "m", "", "secret_key=your_var_key")

	secCmd.Aliases = append(secCmd.Aliases, "sec")
	fn.WithKlFile(secCmd)
}
