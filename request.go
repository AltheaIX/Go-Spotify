package main

import (
	"bufio"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"time"
)

func ResponseReader(response *http.Response) ([]byte, error) {
	var body []byte
	var err error

	reader := bufio.NewReader(response.Body)
	for {
		chunk, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				body = append(body, chunk...)
				break
			}
			return nil, err
		}
		body = append(body, chunk...)
	}
	return body, err
}

func GenerateNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(max-min+1) + min
	return randomInt
}

func GetEmailDomain() ([]string, error) {
	client := http.Client{Timeout: 10 * time.Second}
	var domainList []string

	req, err := http.NewRequest("GET", "https://www.1secmail.com/api/v1/?action=getDomainList", nil)
	if err != nil {
		return domainList, err
	}

	response, err := client.Do(req)
	if err != nil {
		return domainList, err
	}

	scanner, _ := ResponseReader(response)

	err = json.Unmarshal(scanner, &domainList)
	if err != nil {
		return domainList, err
	}

	return domainList, err
}

func GenerateEmail() string {
	return ""
}
