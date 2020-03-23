// Copyright © 2010-12 Qtrac Ltd.
// 
// This program or package and any associated files are licensed under the
// Apache License, Version 2.0 (the "License"); you may not use these files
// except in compliance with the License. You can get a copy of the License
// at: http://www.apache.org/licenses/LICENSE-2.0.
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

const (
	pageTop = `<!DOCTYPE HTML><html><head>
<style>.error{color:#FF0000;}</style></head><title>Statistics</title>
<body><h3>Statistics</h3>
<p>Computes basic statistics for a given list of numbers</p>`
	form = `<form action="/" method="POST">
<label for="numbers">Numbers (comma or space-separated):</label><br />
<input type="text" name="numbers" size="30"><br />
<input type="submit" value="Calculate">
</form>`
	pageBottom = `</body></html>`
	anError    = `<p class="error">%s</p>`
)

type statistics struct {
	numbers    []float64
	mode       []float64
	mean       float64
	median     float64
	stdDev     float64
	passOption []float64
	nk         []float64
	countDig         []float64
	err        error
}

func main() {
	http.HandleFunc("/", homePage)
	fmt.Println("Run server to port :9001")
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}

func factorial(num int) (factorial int) {
	factorial = 1
	ch := num
	for num > 1 {
		factorial *= num
		num--
	}
	fmt.Printf("factorial числа %v = %v\n", ch, factorial)
	return factorial
}

func passOption(numbers []float64) (passOption []float64) {
	for _, n := range numbers {
		fmt.Printf("%v вариантов комбинаций паролей из %v-значного пароля\n", factorial(int(n)), n)
		passOption = append(passOption, float64(factorial(int(n))))
	}
	return passOption
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

func formatStats(stats statistics) string {
	return fmt.Sprintf(`<table border="1">
<tr><th colspan="2">Results</th></tr>
<tr><td>Numbers</td><td>%v</td></tr>
<tr><td>Count</td><td>%d</td></tr>
<tr><td>Mean</td><td>%f</td></tr>
<tr><td>Median</td><td>%f</td></tr>
<tr><td>Std. Dev</td><td>%f</td></tr>
<tr><td>Mode</td><td>%v</td></tr>
<tr><td>Factorials</td><td>%v</td></tr>
<tr><td>Кол-во комбинаций:</td><td>%v</td></tr>
<tr><td>Вариаций пароля:</td><td>%v</td></tr>

</table><hr>%v`, stats.numbers, len(stats.numbers), stats.mean, stats.median, stats.stdDev, stats.mode, stats.passOption, stats.nk, stats.countDig, stats.err)
}

func getStats(numbers []float64) (stats statistics) {
	stats.numbers = numbers
	sort.Float64s(stats.numbers)
	stats.mean = sum(numbers) / float64(len(numbers))
	stats.median = median(numbers)
	stats.stdDev = stdDev(numbers)
	stats.mode = mode(numbers)
	stats.passOption = passOption(numbers)
	stats.nk, stats.err = nk(numbers)
	stats.countDig, stats.err = countDig(numbers)
	return stats
}

func nk(numbers []float64) (nk []float64, err error) {
	if len(numbers) < 2 {
		return nil, errors.New("Ошибка! Кол-во чисел во входящих данных не должно быть меньше двух!")
	}
	n := numbers[0]
	k := numbers[1]
	x := float64(factorial(int(n)) / (factorial(int(k)) * factorial(int(n)-int(k))))
	nk = append(nk, x)
	fmt.Printf("Кол-во вариантов %v-символьного пароля из двух букв, где первая буква повторяется %v раз, а вторая повторяется %v раз: %v\n", n, k, int(n)-int(k), nk)

	return nk, nil
}

func countDig(numbers []float64) (countDig []float64, err error) {
	/**
	Сколько всех четырехзначных чисел можно составить из цифр 1, 5, 6, 7, 8?
	Решение. Для каждого разряда четырехзначного числа имеется пять возможностей, значит N=5*5*5*5=5степень4=625.
	 */
	if len(numbers) < 2 {
		return nil, errors.New("Ошибка! Кол-во чисел во входящих данных не должно быть меньше двух!")
	}
	y := numbers[0]
	n := numbers[1]
	x := math.Pow(n, y)
	countDig = append(countDig, x)
	fmt.Printf("Кол-во вариантов %v-символьного пароля из %v цифр: %v\n", y, n,  countDig)

	return countDig, nil
}

func sum(numbers []float64) (total float64) {
	for _, x := range numbers {
		total += x
	}
	return total
}

func median(numbers []float64) float64 {
	middle := len(numbers) / 2
	result := numbers[middle]
	if len(numbers)%2 == 0 {
		result = (result + numbers[middle-1]) / 2
	}
	return result
}

func stdDev(numbers []float64) float64 {
	var s float64
	for _, x := range numbers {
		s += math.Pow(sum(numbers)/float64(len(numbers))-x, 2)
	}
	s /= float64(len(numbers) - 1)
	s = math.Sqrt(s)
	return s
}

func mode(numbers []float64) (s []float64) {
	var n float64
	fmt.Println("unicode.MaxRune", unicode.MaxRune)
	for k, x := range numbers {
		var i int
		for _, y := range numbers[k:] {
			if x == y {
				i++
			}
			if i == 2 {
				if n != y {
					s = append(s, y)
					n = s[len(s)-1]
					break
				}
			}
		}
	}
	return s
}
