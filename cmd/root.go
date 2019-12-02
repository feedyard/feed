package cmd

import (
  "fmt"
  "os"
  "github.com/spf13/cobra"
  "log"
  homedir "github.com/mitchellh/go-homedir"
  "github.com/spf13/viper"

)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
  Use:   "feedyard",
  Short: "delivery platform cli",
  Long: `The feedyard cli is an example command line tool to demonstrate
many common interactions with a delivery infrastruture, platform product.`,
  // Uncomment the following line to define actions for
  // calling the cli with no commands:
  //	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
  cobra.OnInitialize(initConfig)

  rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", ConfigFileDefaultLocationMsg)
  rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig sets the config values based on the following order of precedent:
// ENV variables
// Config file definitions
// Default values from constant.go
func initConfig() {
  // auth0 application that manages the device authentication for the feedyard cli
  viper.SetDefault("ClientID", ClientID)

  viper.SetDefault("DeviceScopes", DeviceScopes)
  viper.SetDefault("DeviceAudience", DeviceAudience)
  viper.SetDefault("DeviceCodeUrl", DeviceCodeUrl)
  viper.SetDefault("DeviceCodePayload",
    "client_id=" + viper.GetString("ClientID") +
    "&scope=" + viper.GetString("DeviceScopes") +
    "&audience=" + viper.GetString("DeviceAudience"))
  viper.SetDefault("AuthUrl", AuthUrl)
  viper.SetDefault("AuthGrantType", AuthGrantType)
  viper.SetDefault("UserInfoUrl", UserInfoUrl)

  viper.SetEnvPrefix(ConfigEnvDefault)
  viper.AutomaticEnv() // read in environment variables that match
  if cfgFile != "" {
    // Use config file from the flag.
    viper.SetConfigFile(cfgFile)
  } else {
    viper.AddConfigPath(defaultConfigLocation())
    viper.SetConfigName(ConfigFileDefaultName)
  }

  // If a config file is found, read it in, else write a blank.
  if err := viper.ReadInConfig(); err != nil {
    home := defaultConfigLocation()
    os.MkdirAll(home, 600)
    fmt.Println(home+"/"+ConfigFileDefaultName+"."+ConfigFileDefaultType)
    emptyFile, err := os.Create(home+"/"+ConfigFileDefaultName+"."+ConfigFileDefaultType)
  	if err != nil {
  		log.Fatal(err)
  	}
  	emptyFile.Close()
  }
  
}

func defaultConfigLocation() string {
  home, err := homedir.Dir()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  return home + ConfigFileDefaultLocation
}
