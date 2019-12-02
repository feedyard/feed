package cmd

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure feedyard cli options.",
	Long: `Configure feedyard cli options. Use to view or set the feedyard cli
default configuration. If your config file does not exist the init command will
create it for you (the default location is ~/.feedyard/config.yaml). To keep an
existing value, hit enter when prompted for the value. When you are prompted
for information, the current value will be displayed in [brackets]. If the
config item has no value, it be displayed as []. The configure command will
include ENVIRONMENT variable overrides if present.

Configuration variables

       The following configuration variables are supported in the config file.
       A solid bullet (•) means you can set the value with the config command.
       Other setting must be added to the config file directly:

       • team - The team membership to assume when executing commands. You may
         be a member of more tha one team, however the cli requires a single
         team be assumed when running commands.

       o access_token - Device authorization access token. The feedyard cli calls
         platform apis to perform commands. In order to use the tool, you
         first perform a device authorization (see login command). The access
         and id tokens are stored in the local device configuration file.
         The cli will automatically request refresh tokens so long as the
         access_token is present. Deleting the configuration file or logging
         in from another device will require you to repeat the device
         authorization. Normally, you will not need to directly set this config
         value, however you may manually set the access and id tokens.

       o id_token  - The id_token is requested as part of the login authorization
         and includes the scopes associated with your identity.

       o clientdevicecodeurl - The Auth0 device code authorization url.

       o clientid - The Auth0 application client id.

       o clientscopes - The requested scopes.

       o clientaudience - API resouces`,
	// Run: func(cmd *cobra.Command, args []string) {
	//
	// },
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "display current configuration",
	Long: `Print the current configuration to stdout.`,
	Run: func(cmd *cobra.Command, args []string) {
		//a := []string{"Foo", "Bar"}
		for _, s := range viper.AllKeys() {
      fmt.Printf("%s: %s\n", s, viper.GetString(s))
		}
	},
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "set current configuration",
	Long: `Set the current cli configuration values.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Set current feedyard cli configuration")
    for _, setting := range args {
      switch setting {
      case "team":
        config_set_team()
      default:
        fmt.Println("Not supported")
      }
    }
	},
}

var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "write all configuration settings to current config file",
	Long: `Most configurations are managed for you by the cli. Use this commands
to explicitly add all configurations to the config file. This can
assisst in future upgrades and remote changes to maintain backward
compatibility`,
	Run: func(cmd *cobra.Command, args []string) {
    viper.WriteConfigAs(viper.ConfigFileUsed())
    fmt.Println("All config settings written to config file")
	},
}

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "clear current configuration",
	Long: `Clear all configuration setting and reset to defaults. Will require re-athentication.`,
	Run: func(cmd *cobra.Command, args []string) {
    var err = os.Remove(viper.ConfigFileUsed())
    if err != nil {
      fmt.Println(err.Error())
    }
    fmt.Println("All config setting reset to default")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(getCmd)
	configCmd.AddCommand(setCmd)
  configCmd.AddCommand(writeCmd)
  configCmd.AddCommand(resetCmd)
}

func config_set_team() {
  viper.Set("team", getInput("team", viper.GetString("team")))
  viper.WriteConfigAs(viper.ConfigFileUsed())
}

func getInput(message string, current string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n%s: [%s] ", message, current)
	response, _ := reader.ReadString('\n')
  response = strings.TrimSuffix(response, "\n")
  if response == "" {
    return current
  } else {
    return response
  }
}
