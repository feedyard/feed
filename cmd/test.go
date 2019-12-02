package cmd

import (
	//"os"
	"fmt"
	//"reflect"
	//"encoding/json"
	//"net/http"
	  //"strings"
	//"io/ioutil"
	"github.com/spf13/cobra"
	//"github.com/spf13/viper"
	//"time"
)

type PollsResp struct {
  Error string `json:"error"`
  ErrorDescription string `json:"error_description"`
}

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "testing device fllow auth",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("running test")
		authorize()
		// var testvar = new(PollsResp)
		// t2(testvar)
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	testCmd.Flags().StringP("team", "t", "", "team name")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
