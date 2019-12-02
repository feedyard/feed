package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createCmd represents the create command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:`,
	Run: func(cmd *cobra.Command, args []string) {
    if viper.IsSet("access_token") {
      fmt.Println("is set")
    } else {
      authenticate()
    }
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
