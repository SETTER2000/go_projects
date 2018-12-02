// Срез (slice) – это ссылка на скрытый массив. И срез, извлеченный из среза , ссылается на тот же самый скрытый массив.
package main

import (
	"fmt"
)

type Product struct {
	name  string
	price float64
}

type Point struct{ x, y, z int }

func (point Point) String() string {
	return fmt.Sprintf("(%d,%d,%d)", point.x, point.y, point.z)
}

func (p Product) String() string {
	return fmt.Sprintf("%s (%.2f)", p.name, p.price)
}

func main() {
	s := []string{"A", "B", "C", "D", "E", "F", "G"}
	t := s[:5]
	u := s[3 : len(s)-1]
	fmt.Println(s, t, u)
	u[1] = "x"
	fmt.Println(s, t, u)

	b := make([]byte, 20, 60)
	g1 := make([][]int, 3)
	for i := range g1 {
		g1[i] = make([]int, 3)
	}

	g1[1][0], g1[1][1], g1[1][2] = 8, 6, 2
	g2 := [][]int{{4, 3, 0}, {8, 6, 2}, {0, 0, 0}}
	cities := []string{"Shanghai", "Mumbai", "Moscow", "Istambul", "Orel"}
	cities[len(cities)-1] = "Baku"
	fmt.Println("Type     Len Cap Contents")
	fmt.Printf("%-8T %2d %3d %v\n", b, len(b), cap(b), b)
	fmt.Printf("%-8T %2d %3d %q\n", cities, len(cities), cap(cities), cities)
	fmt.Printf("%-8T %2d %3d %v\n", g1, len(g1), cap(g1), g1)
	fmt.Printf("%-8T %2d %3d %v\n", g2, len(g2), cap(g2), g2)

	s2 := []string{"A", "B", "C", "D", "E", "F", "G"}
	d := s2[2:6]
	fmt.Println(d, s2, "=", s2[:4], "+", s2[4:])
	s2[3] = "i"
	d[len(d)-1] = "j"
	fmt.Println(d, s2, "=", s2[:4], "+", s2[4:])

	// Итерация по срезам
	amounts := []float64{237.81, 261.87, 273.93, 279.99, 281.07, 303.17, 231.47, 227.33, 209.23, 197.09}
	sum := 0.0
	// первых два элемента среза
	//for _, amount := range amounts[:2] {
	for _, amount := range amounts {
		sum += amount
	}
	fmt.Printf("\nСумма всех значений среза:\nƩ %.1f → %.1f\n", amounts, sum)

	/*
		Если в цикле необходимо иметь возможность изменять значения
		элементов, следует использовать цикл for , поставляющий только
		допустимые индексы, но не копии элементов в срезе.
	*/
	for i := range amounts {
		amounts[i] *= 1.05 // увеличение на 5% каждого значения
		sum += amounts[1]  // сумма всех изменённых значений среза
	}
	fmt.Printf("\nИзменённые значения среза:\nƩ %.1f → %.1f\n", amounts, sum)

	/*
		Выше отмечалось, что цикл for ...range не может использоваться
		в ситуациях, когда требуется иметь возможность изменять значе-
		ния элементов среза во время итераций. Однако в этом примере
		для всех элементов среза удалось благополучно увеличить значе-
		ния полей price. В каждой итерации переменной product присва-
		ивается копия значения типа *Product – то есть копия указателя
		из соответствующего элемента среза, ссылающегося на фактическое
		значение типа Product в памяти. Поэтому изменения применяются
		к значению Product, на которое указывает указатель, а не к копии
		указателя *Product.
				"Программирование на Go. Разработка приложений XXI века - 2013"
	*/
	// Добавляемые значения должны иметь тот же тип, что и сам срез.
	fmt.Println("\nВывод значений типа Product:")
	products := []*Product{{"Spanner", 3.99}, {"Winch", 2.49}, {"Screwdriver", 1.99}}
	fmt.Println(products)
	for _, product := range products {
		product.price += 0.50
	}
	fmt.Println(products)

	// Для добавления новых значений в конец среза можно использовать встроенную функцию append().
	s3 := []string{"A", "B", "C", "D", "E", "F", "G"}
	t2 := []string{"K", "L", "M", "N"}
	u2 := []string{"m", "n", "o", "p", "q", "r"}
	s3 = append(s3, "h", "i", "j") // Добавить отдельные значения
	s3 = append(s3, t2...)         // Добавить все элементы из среза t2
	s3 = append(s3, u2[2:5]...)    // Добавить часть элементов из среза u2
	r := []byte{'U', 'V'}
	letters := "wxy"
	r = append(r, letters...) // Добавить байты из строки в срез с байтами
	fmt.Printf("\nДобавленные значения в срез:\n%v\n%s\n", s3, r)

	// Примеры демонстрируют создание и заполнение отображений с ключами типа string и значениями типа float64
	massForPlanet := make(map[string]float64) // То же, что и: map[string]float64{}
	massForPlanet["Mercury"] = 0.06
	massForPlanet["Venus"] = 0.82
	massForPlanet["Earth"] = 1.00
	massForPlanet["Mars"] = 0.11
	fmt.Println(massForPlanet)

	// Представлен пример использования указателей в качестве ключей. *Point (указатель типа Point),
	triangle := make(map[*Point]string, 3)
	triangle[&Point{89, 47, 27}] = "α"
	triangle[&Point{86, 65, 86}] = "β"
	triangle[&Point{86, 65, 86}] = "ɣ"
	fmt.Println("\nИспользование указателей в качестве ключей:")
	fmt.Println(triangle)

	/*
		В языке Go допускается использовать
		структуры в качестве ключей отображений, при условии что типы
		их полей поддерживают операции сравнения с помощью операторов
		== и !=

		В примере ниже вывод будет содержать лишь последнюю введённую точку в отображение nameForPoint,
		т.к. ключи отображения, в данном случаи это структура Point, одинаковые.
		Для того чтобы ввести две одинаковые точки типа Point, в качестве ключей
		отображения следует использовать указатели на тип Point (*Point). Как показано
		в примере выше.
	*/
	nameForPoint := make(map[Point]string) // То же, что и: map[Point]string{}
	nameForPoint[Point{54, 91, 78}] = "x"
	nameForPoint[Point{54, 91, 78}] = "y"
	fmt.Println("\nСтруктуры в качестве ключей отображений:")
	fmt.Println(nameForPoint)

	// В этом отображение целиком было создано применением синтаксиса составных литералов
	populationForCity := map[string]int{"Istanbul": 12610000, "Karachi": 10620000, "Mumbai": 12690000, "Shanghai": 13680000}
	fmt.Println("\nCинтаксис составных литералов:")
	for city, population := range populationForCity {
		fmt.Printf("%-15s %8d\n", city, population)
	}

	/*
		Поиск по отображению. Вариант 1
		Если выполняется поиск по ключу, присутствующему в отображе-
		нии, возвращается соответствующее ему значение. Но если искомый
		ключ отсутствует, возвращается нулевое значение данного типа.
	*/

	population := populationForCity["Mumbai"]
	fmt.Println("\nПоиск по отображению. Вариант 1:\nMumbai’s population is", population)
	population = populationForCity["Emerald City"]
	fmt.Println("Emerald City’s population is", population)

	/*
		Поиск по отображению. Вариант 2
		0 для ключа "Emerald City" нельзя
		сказать, означает ли оно, что в городе Emerald City действительно
		отсутствует население или этот город отсутствует в отображении.
		Решение этой проблемы предлагает второй способ поиска.
	*/

	fmt.Printf("\nПоиск по отображению. Вариант 2:\n")
	city := "Istanbul"
	if population, found := populationForCity[city]; found {
		fmt.Printf("%s's population is %d\n", city, population)
	} else {
		fmt.Printf("%s's population data is unavailable\n", city)
	}
	city = "Emerald City"
	_, present := populationForCity[city]
	fmt.Printf("%q is in the map == %t\n", city, present)

}
