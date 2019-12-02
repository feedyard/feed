package cmd

const (
  ClientID = "5kZqhP7fCAaVs2u0fMiMdWrULhTtYCPq"

  DeviceCodeUrl = "https://feedyard-dev.auth0.com/oauth/device/code"
  DeviceScopes = "offline_access openid https://github.org/feedyard/teams"
  DeviceAudience = "https://api.feedyard.xyz/v1"
  AuthUrl = "https://feedyard-dev.auth0.com/oauth/token"
  AuthGrantType = "urn:ietf:params:oauth:grant-type:device_code"
  UserInfoUrl = "https://feedyard-dev.auth0.com/userinfo"

  ConfigEnvDefault = "FEEDYARD"
  ConfigFileDefaultLocation = "/.feedyard"  // path will begin with $HOME dir
  ConfigFileDefaultName = "config"
  ConfigFileDefaultType = "yaml"
  ConfigFileDefaultLocationMsg = "config file (default is $HOME/.feedyard/config.yaml)"
)
