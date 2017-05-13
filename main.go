package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

//BotConfig stores the configuration
type BotConfig struct {
	Url      string `json:"url"`
	Email    string `json:"email"`
	Password string `json:"password"`
	DeviceId string `json:"deviceId"`
}

//LoginResp stores the token
type LoginResp struct {
	Username string `json:"user.username"`
	Id       string `json:"user._id"`
	Token    string `json:"token"`
}

var botConfig BotConfig
var loginResp LoginResp

func readConfig() {
	file, e := ioutil.ReadFile("botConfig.json")
	if e != nil {
		fmt.Println("error:", e)
	}
	content := string(file)
	json.Unmarshal([]byte(content), &botConfig)
}

func login() bool {
	url := botConfig.Url + "/devices/id/" + botConfig.DeviceId + "/login"
	jsonStr := `{
		"email":"` + botConfig.Email + `",
		"password":"` + botConfig.Password + `",
		"userAgent":"DEVICE:` + botConfig.DeviceId + `"
	}`
	b := strings.NewReader(jsonStr)
	req, _ := http.NewRequest("POST", url, b)
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return false
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal([]byte(body), &loginResp)

	fmt.Println("token: " + loginResp.Token)
	c.Cyan(loginResp.Username)
	return true
}

func sendRegister(dateRaw time.Time, value float64) bool {
	date := dateRaw.Format("2006-01-02 15:04:05.000Z")

	url := botConfig.Url + "/devices/id/" + botConfig.DeviceId + "/registers/"
	jsonStr := `{
		"deviceid":"` + botConfig.DeviceId + `",
		"date":"` + date + `",
		"value":"` + strconv.FormatFloat(value, 'f', 6, 64) + `"
	}`
	b := strings.NewReader(jsonStr)
	req, _ := http.NewRequest("POST", url, b)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-access-token", loginResp.Token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return false
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

	fmt.Print("register sent")
	c.Green(date + ": " + strconv.FormatFloat(value, 'f', 6, 64))
	return true
}

func main() {
	readConfig()
	logged := false
	for logged != true {
		logged = login()
		sleepSec(5)
	}

	args := os.Args[1:]
	fmt.Println(len(args))
	if len(args) > 0 {
		fakeInputData()
	} else {
		inputRead()
	}
}
