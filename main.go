package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/post04/chesscomgen/chess"
)

var (
	usernames []string
	config    *c
)

type c struct {
	Threads  int    `json:"threads"`
	Proxy    string `json:"proxy"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
	// ! load usernames
	f, err := os.ReadFile("./data/usernames.txt")
	if err != nil {
		panic(err)
	}
	usernames = strings.Split(string(f), "\r\n")
	fmt.Println("Loaded", len(usernames), "usernames!")
	// ! load config
	f, err = os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(f, &config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", config)
}

var (
	writingPool = make(chan string, 999)
)

func writerHandler() {
	accounts := ""
	f, err := os.ReadFile("./data/accounts.txt")
	if err != nil {
		panic(err)
	}
	accounts = string(f)
	for {
		a := <-writingPool
		if accounts != "" {
			accounts += "\r\n" + a
		} else {
			accounts += a
		}
		os.WriteFile("./data/accounts.txt", []byte(accounts), 0064)
	}
}

func main() {
	go writerHandler()
	for i := 0; i < config.Threads; i++ {
		go func() {
			for {
				client := chess.NewClient(getRandomGenProxy())
				username := getUsername()
				email := getRandomEmail()
				err := client.CreateAccount(username, config.Password, email)
				if err != nil {
					fmt.Println(err)
					continue
				}
				fmt.Printf("Generated account\n	Username: %s\n	Email: %s\n	Password: %s\n	Oauth: %s\n	Refresh: %s\n", username, email, config.Password, client.Bearer[:50], client.RefreshToken[:50])
				writingPool <- fmt.Sprintf("%s | %s | %s | %s | %s | %s | %s", username, email, config.Password, client.Bearer, client.RefreshToken, client.LoginToken, client.SessionID)

			}
		}()
	}
	select {}
}
