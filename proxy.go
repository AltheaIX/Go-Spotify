package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
)

var proxyList []string

func ReadProxyFromFile(fileName string) error {
	file, err := os.Open(fileName + ".txt") // Replace with the path to your file
	if err != nil {
		file, _ = os.Create(fileName + ".txt")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, ":")
		re := regexp.MustCompile("[0-9]+")
		proxyFormat := fmt.Sprintf("%s:%s", splitLine[0], re.FindAllString(splitLine[1], -1)[0])
		proxyList = append(proxyList, proxyFormat)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return nil
}

func BuildProxyURL(proxy string) string {
	return strings.TrimRight("http://"+proxy, "\x00")
}

func GetProxy() string {
	var proxy string

	proxy = proxyList[rand.Intn(len(proxyList)-1)]

	return proxy
}
