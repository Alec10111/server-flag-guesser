package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gocolly/colly/v2"
)

func removeBrackets(text string, opener string, closer string) string {
	var res string
	var stack []string
	for _, char := range text {
		if string(char) == opener {
			stack = append(stack, string(char))
		} else if string(char) == closer {
			stack = stack[1:]
		} else if len(stack) == 0 {
			res += string(char)
		}
	}
	return res
}

func fetchDescription(name string) string {

	var result string
	c := colly.NewCollector()
	c.OnHTML(".mw-parser-output", func(e *colly.HTMLElement) {

		found := false
		e.ForEach("p", func(_ int, elem *colly.HTMLElement) {
			if !found && strings.Contains(elem.Text, name) {
				result = elem.Text
				found = true
			}
		})
	})
	fmt.Println("Scrapping Complete")

	c.Visit("https://en.wikipedia.org/wiki/" + name)
	return removeBrackets(removeBrackets(result, "(", ")"), "[", "]")
}

func (app *application) writejson(w http.ResponseWriter, status int, data interface{}, wrap string) error {
	wrapper := make(map[string]interface{})
	wrapper[wrap] = data

	js, err := json.Marshal(wrapper)
	if err != nil {
		return err
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func (app *application) errorjson(w http.ResponseWriter, err error) {
	type jsonerror struct {
		Message string `json:"message"`
	}
	theError := jsonerror{
		Message: err.Error(),
	}
	app.writejson(w, http.StatusBadRequest, theError, "error")

}
