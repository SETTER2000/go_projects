package quadratic

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const (
	pageTop    = `<!DOCTYPE HTML><html><head>
<style>.error{color:#FF0000;}</style></head>
<title>Quadratic Equation Solver</title><body>
<h3>Quadratic Equation Solver</h3><p>Solves equations of the form
a<i>x</i>² + b<i>x</i> + c</p>`
	form       = `<form action="/" method="POST">
<input type="text" name="a" size="1"><label for="a"><i>x</i>²</label> +
<input type="text" name="b" size="1"><label for="b"><i>x</i></label> +
<input type="text" name="c" size="1"><label for="c"> →</label>
<input type="submit" name="calculate" value="Calculate">
</form>`
	pageBottom = "</body></html>"
	error      = `<p class="error">%s</p>`
	solution   = "<p>%s → %s</p>"
)

func main() {
	http.HandleFunc("/", homePage)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm() // Должна вызываться перед записью в ответ
	fmt.Fprint(writer, pageTop, form)
	if err != nil {
		fmt.Fprintf(writer, anError, err)
	} else {
		if numbers, message, ok := processRequest(request); ok {
			stats := getStats(numbers)
			fmt.Fprint(writer, formatStats(stats))
		} else if message != "" {
			fmt.Fprintf(writer, anError, message)
		}
	}
	fmt.Fprint(writer, pageBottom)

	formatQuestion()
	solve()
	formatSolutions()
}

func processRequest(request *http.Request) ([]float64, string, bool) {
	var numbers []float64
	if slice, found := request.Form["numbers"]; found && len(slice) > 0 {
		text := strings.Replace(slice[0], ",", " ", -1)
		for _, field := range strings.Fields(text) {
			if x, err := strconv.ParseFloat(field, 64); err != nil {
				return numbers, "'" + field + "' is invalid", false
			} else {
				numbers = append(numbers, x)
			}
		}
	}
	if len(numbers) == 0 {
		return numbers, "", false // no data first time form is shown
	}
	return numbers, "", true
}