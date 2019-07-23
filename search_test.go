package googleit

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	log "github.com/schollz/logger"
	"github.com/stretchr/testify/assert"
)

func TestSearchStartPage(t *testing.T) {
	log.SetLevel("trace")
	urls, err := StartPage("banana chocolate chip cookie recipe", Options{NumPages: 30, MustInclude: []string{"banana", "chocolate", "chip", "cookie"}})
	assert.Nil(t, err)
	assert.True(t, len(urls) > 10)
	fmt.Println(strings.Join(urls, "\n"))
}

func TestSearchBing(t *testing.T) {
	log.SetLevel("trace")
	urls, err := Bing("banana chocolate chip cookie recipe", Options{NumPages: 30, MustInclude: []string{"banana", "chocolate", "chip", "cookie"}})
	assert.Nil(t, err)
	assert.True(t, len(urls) > 100)
	fmt.Println(strings.Join(urls, "\n"))
}

func TestSearchDuckDuckGo(t *testing.T) {
	log.SetLevel("trace")
	urls, err := DuckDuckGo("banana chocolate chip cookie recipe", Options{NumPages: 30, MustInclude: []string{"banana", "chocolate", "chip", "cookie"}})
	assert.Nil(t, err)
	assert.True(t, len(urls) > 100 && len(urls) < 300)
	fmt.Println(strings.Join(urls, "\n"))
}

func TestSearchBoth(t *testing.T) {
	log.SetLevel("trace")
	urls, err := Search("banana chocolate chip cookie recipe", Options{NumPages: 30, MustInclude: []string{"banana", "chocolate", "chip", "cookie"}})
	assert.Nil(t, err)
	assert.True(t, len(urls) > 10)
}

func TestSearchBingWithTor(t *testing.T) {
	log.SetLevel("trace")
	urls, err := Bing("cat animal wiki", Options{NumPages: 1, UseTor: true})
	assert.Nil(t, err)
	assert.True(t, len(urls) >= 9)
}

func TestRun(t *testing.T) {
	log.SetLevel("trace")
	urls, err := Search("banana chocolate chip cookie recipe", Options{NumPages: 3, MustInclude: []string{"banana", "chocolate", "chip", "cookie"}})
	assert.Nil(t, err)
	ioutil.WriteFile("urls.txt", []byte(strings.Join(urls, "\n")), 0644)
}
