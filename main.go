package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	domain "github.com/schoolboybru/git-notifier/internal/domain"
)

func main() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.github.com/notifications", nil)

	if err != nil {
		fmt.Print("Something went wrong")
	}

	var userName string = os.Getenv("USERNAME")
	var password string = os.Getenv("PASSWORD")

	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.SetBasicAuth(userName, password)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	var notifications []domain.Notification

	json.Unmarshal(bodyBytes, &notifications)
	fmt.Printf("API response as struct %+v\n", notifications)

}
