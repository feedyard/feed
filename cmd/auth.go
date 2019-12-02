package cmd

import (
	"fmt"
  "os"
  //"reflect"
  "encoding/json"
	"net/http"
	"strings"
	"io/ioutil"
	"github.com/spf13/viper"
  "time"
)

type DeviceCodeResp struct {
    DeviceCode string `json:"device_code"`
    UserCode string `json:"user_code"`
    VerificationUri string `json:"verification_uri"`
    ExpiresIn int `json:"expires_in"`
    Interval int `json:"interval"`
    VerificationUriComplete string `json:"verification_uri_complete"`
}

type AuthResp struct {
    AccessToken string `json:"access_token"`
    RefreshToken string `json:"refresh_token"`
    IdToken string `json:"id_token"`
    Scope string `json:"scope"`
    ExpiresIn int `json:"expires_in"`
    TokenType string `json:"token_type"`
}

type PollResp struct {
  Error string `json:"error"`
  ErrorDescription string `json:"error_description"`
}

type UserInfoResp struct {
  Sub string `json:"sub"`
  Claims []string `json:"https://github.org/feedyard/teams"`
}


func authenticate() {
  var body []byte
  var status string
  var authorization = new(AuthResp)
  var poll_result = new(PollResp)

  // first step in device code flow,
  device_code, interval := getDeviceCode()

  for {
    body, status = poll(device_code)

    if status == "200 OK" {
      parseResponse([]byte(body), &authorization)
      viper.Set("access_token", authorization.AccessToken)
      viper.Set("refresh_token", authorization.RefreshToken)
      viper.Set("id_token", authorization.IdToken)
      viper.Set("expires_in", authorization.ExpiresIn)
      viper.Set("LastLogin", time.Now())
      viper.WriteConfigAs(viper.ConfigFileUsed())
      break
    } else {
      parseResponse([]byte(body), &poll_result)
      switch poll_result.Error {
      case "authorization_pending":
        fmt.Println("waiting...")
        time.Sleep(time.Duration(interval) * time.Second)
      case "expired_token":
        fmt.Println("code expired")
        os.Exit(1)
      case "access_denied":
        fmt.Println("Access denied")
        os.Exit(1)
      default:
        fmt.Printf("error: unexpected response %s", poll_result.Error)
        os.Exit(1)
      }
    }
  }
}

func poll(device_code string) ([]byte, string) {
  var netClient = &http.Client{
    Timeout: time.Second * 30,
  }
  fmt.Println("before")
  //url := viper.GetString("ClientAuthUrl")
  payload := strings.NewReader("client_id=" + viper.GetString("ClientID") +
                               "&device_code=" + device_code +
                               "&grant_type=" + viper.GetString("AuthGrantType"))

  req, _ := http.NewRequest("POST", viper.GetString("AuthUrl"), payload)
  req.Header.Add("content-type", "application/x-www-form-urlencoded")

  resp, err := netClient.Do(req)
  if err != nil {
      fmt.Println("error")
  }
  defer resp.Body.Close()
  body, _ := ioutil.ReadAll(resp.Body)
  return body, resp.Status
}

func getDeviceCode() (string, int) {
  var device = new(DeviceCodeResp)
  var netClient = &http.Client{
    Timeout: time.Second * 30,
  }

  req, _ := http.NewRequest("POST", viper.GetString("DeviceCodeUrl"), strings.NewReader(viper.GetString("DeviceCodePayload")))
  req.Header.Add("content-type", "application/x-www-form-urlencoded")

  resp, err := netClient.Do(req)
  if err != nil {
    fmt.Println("error")
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)

  err = parseResponse([]byte(body), &device)

  fmt.Println("feedyard will attempt to open a browser window where you can authenticate and verify your laptop.")
  fmt.Println("If the windows does not open, go to the link below and enter the code.\n")
  fmt.Printf("%s\ncode: %s\n", device.VerificationUri, device.UserCode)

  return device.DeviceCode, device.Interval
}

func authorize() {
  if !viper.IsSet("team") {
    config_set_team()
  }
  if tokenExpired() || !viper.IsSet("refresh_token") {
    authenticate()
  }
  
}

func validate(team string) bool {
  var userinfo = new(UserInfoResp)
  var netClient = &http.Client{
    Timeout: time.Second * 30,
  }

  req, err := http.NewRequest("GET", viper.GetString("UserInfoUrl"), nil)
  req.Header.Add("Authorization", "Bearer " + viper.GetString("access_token"))

  resp, err := netClient.Do(req)
  if err != nil {
    fmt.Println("error")
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)

  err = parseResponse([]byte(body), &userinfo)
  if !teamInClaims(team, userinfo.Claims) {
    fmt.Printf("Access Denied to %s\n\n", team)
    fmt.Println("Your teams:")
    for _, val := range userinfo.Claims {
      fmt.Println(val)
    }
    return false
  }
  return true
}

func parseResponse(body []byte, class interface{}) (error) {
  err := json.Unmarshal(body, class)
  if(err != nil){
      fmt.Println("whoops:", err)
  }
  return err
}

func teamInClaims(team string, claims []string) bool {
  for _, val := range claims {
      if val == team {
          return true
      }
  }
  return false
}

func tokenExpired() bool {
  last_login, _ := time.Parse(time.RFC3339, viper.GetString("LastLogin"))
  expire_time := last_login.Add(time.Duration(viper.GetInt("expires_in")) * time.Second)

  return time.Now().After(expire_time)
}
