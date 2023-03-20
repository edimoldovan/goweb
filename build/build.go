package build

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
)

func MinifyCSS(body string) string {
	result := bytes.NewBuffer([]byte{})
	bodyReader := strings.NewReader(body)
	err := css.Minify(minify.New(), result, bodyReader, map[string]string{})
	if err != nil {
		log.Println(fmt.Sprintf("CSS minify error error: %v", err))
	}
	return result.String()
}

func BuildCSS() {
	content, err := ioutil.ReadFile("./assets/css/global.css")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")

	log.Println("building and minifying global css with tdewolff/minify...")
	path := ""
	nestedString := ""
	for _, line := range lines {
		if line != "" && line[0] == 64 {
			path = fmt.Sprintf(`./assets/css/%s`, strings.ReplaceAll(strings.Split(line, " ")[1], "\"", ""))
			path = strings.ReplaceAll(path, ";", "")
			fileContent, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}
			if fileContent != nil {
				nestedString = nestedString + string(fileContent)
			}
		}
	}

	minifiedGlobalCssFile, err := os.Create("./public/css/global.min.css")
	if err != nil {
		panic(err)
	}
	nestedN, err := minifiedGlobalCssFile.WriteString(MinifyCSS(nestedString))
	if err != nil {
		log.Println(nestedN)
		panic(err)
	}
	defer minifiedGlobalCssFile.Close()
}
