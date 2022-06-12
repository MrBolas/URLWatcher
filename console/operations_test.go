package console

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizeInputs(t *testing.T) {

	testCommand := "testcommand argument1 argument2"

	inputs, err := SanitizeInputs(testCommand)
	if err != nil {
		fmt.Println(err)
	}

	assert.Equal(t, "argument1", inputs[0])
	assert.Equal(t, "argument2", inputs[1])

}

func TestUrlsReadFromFile(t *testing.T) {

	pathToFile := "../testfolder/links"

	urls, err := ReadUrlsFromFile(pathToFile)
	if err != nil {
		fmt.Println(err)
	}

	assert.Equal(t, 5, len(urls))
}

func TestConvertToUrl(t *testing.T) {

	candidateUrls := []string{"https://manpages.debian.org/",
		"http://www.google.com",
		"http://example.com/",
		"http://zealwebtech.com/",
		"http://techvynsys.com/v2/"}

	urls := ConvertToUrl(candidateUrls)

	assert.Equal(t, len(candidateUrls), len(urls))
}
