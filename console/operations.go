package console

import (
	"bufio"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"
)

func SanitizeInputs(command string) ([]string, error) {

	arguments := strings.Split(command, " ")
	if len(arguments) < 2 {
		return []string{}, errors.New("not enough arguments")
	}

	return arguments[1:], nil
}

func ReadUrlsFromFile(path string) ([]url.URL, error) {

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var fileLines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	urls := ConvertToUrl(fileLines)

	return urls, nil
}

func ConvertToUrl(lines []string) []url.URL {
	var urls []url.URL
	var failedUrls []string

	for _, line := range lines {
		u, err := url.Parse(line)
		if err != nil {
			failedUrls = append(failedUrls, fmt.Sprintf("%s parsing failed: %s\n", line, err))
			continue
		}

		urls = append(urls, *u)
	}

	if len(failedUrls) > 0 {
		printLoadFails(failedUrls)
	}

	return urls
}

func printLoadFails(failedUrls []string) {
	fmt.Println("Failure loading URL(s):")
	for _, url := range failedUrls {
		fmt.Println(url)
	}
}
