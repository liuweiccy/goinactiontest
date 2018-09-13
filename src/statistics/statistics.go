package statistics

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type statistics struct {
	numbers []float64
	mean    float64
	median  float64
}

func GetState(numbers []float64) (stats statistics) {
	stats.numbers = numbers
	sort.Float64s(stats.numbers)
	stats.mean = sum(numbers) / float64(len(numbers))
	stats.median = median(numbers)
	return
}

func sum(numbers []float64) float64 {
	var sum float64
	sum = 0
	for _, num1 := range numbers {
		sum += num1
	}
	return sum
}

func median(numbers []float64) float64 {
	len1 := len(numbers)
	if len1 == 0 {
		errors.New("数组为空")
	}
	index := (len1+1)/2 - 1
	return numbers[index]
}

func Start() {
	http.HandleFunc("/", homePage)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("Failed to start server", err)
	}
}

const (
	pageTop = `<h1>统计</h1></br>`
	form    = `<form action="/" method="POST>
			<label for="numbers">Number (comma or space-separated):</label><br/>
			<input type="text" name="numbers" size="30"><br/>
			<input type="submit" value="计算"><br/>
			`
	pageBottom = `<h5>统计完成</h5></br>`
	anError    = `<p class = "error">%s</p>`
)

func homePage(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	fmt.Fprint(writer, pageTop, form)
	if err != nil {
		fmt.Fprintf(writer, anError, err)
	} else {
		if numbers, message, ok := processRequest(request); ok {
			stats := GetState(numbers)
			fmt.Fprintln(writer, stats)
		} else if message != "" {
			fmt.Fprintf(writer, anError, message)
		}
	}

	fmt.Fprint(writer, pageBottom)
}

func processRequest(request *http.Request) ([]float64, string, bool) {
	var numbers []float64
	if slice, found := request.Form["numbers"]; found && len(slice) > 0 {
		text := strings.Replace(slice[0], ",", " ", -1)

		for _, filed := range strings.Fields(text) {
			if x, err := strconv.ParseFloat(filed, 64); err != nil {
				return numbers, "'" + filed + "' is invalid", false
			} else {
				numbers = append(numbers, x)
			}
		}
	}

	if len(numbers) == 0 {
		return numbers, "", false
	}

	return numbers, "", true
}
