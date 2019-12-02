package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:`,
	// Run: func(cmd *cobra.Command, args []string) {
	//
	// },
}

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Create repo command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(viper.GetString("team"))
		if validate(viper.GetString("team")) {
			fmt.Println("authorized to create repo")
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(repoCmd)

}
