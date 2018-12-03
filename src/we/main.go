package main

import (
	"encoding/json"
	"github.com/auth0-community/auth0"
	"github.com/gorilla/mux"
	"net/http"
	// Добавить зависимость пакетов обработчиков
	//jwtmiddleware "github.com/auth0/go-jwt-middleware"
	//jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"os"
	//	"time"
	"fmt"
	jose "gopkg.in/square/go-jose.v2"
)

/*
Сначала мы создадим новый тип под названием Product
Этот тип будет содержать информацию о опыте VR
*/
type Product struct {
	Id          int
	Name        string
	Slug        string
	Description string
}

// Мы создадим наш каталог опыта VR и храним их в кусочке.
var products = []Product{
	Product{Id: 1, Name: "Hover Shooters", Slug: "hover-shooters", Description: "Shoot your way to the top on 14 different hoverboards"},
	Product{Id: 2, Name: "Ocean Explorer", Slug: "ocean-explorer", Description: "Explore the depths of the sea in this one of a kind underwater experience"},
	Product{Id: 3, Name: "Dinosaur Park", Slug: "dinosaur-park", Description: "Go back 65 million years in the past and ride a T-Rex"},
	Product{Id: 4, Name: "Cars VR", Slug: "cars-vr", Description: "Get behind the wheel of the fastest cars in the world."},
	Product{Id: 5, Name: "Robin Hood", Slug: "robin-hood", Description: "Pick up the bow and arrow and master the art of archery"},
	Product{Id: 6, Name: "Real World VR", Slug: "real-world-vr", Description: "Explore the seven wonders of the world in VR"},
}

func main() {
	// Здесь мы создаем маршрутизатор gorilla / mux
	r := mux.NewRouter()

	// На странице по умолчанию мы просто будем обслуживать страницу статического индекса.
	r.Handle("/", http.FileServer(http.Dir("./views/")))
	// Мы настроим наш сервер, чтобы мы могли обслуживать статические ресурсы,
	// такие как изображения, css из маршрута / static / {file}
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	/* Мы добавим промежуточное ПО к нашим продуктам и маршрутам обратной связи.
	Маршрут статуса будет общедоступным */
	r.Handle("/products", authMiddleware(ProductsHandler)).Methods("GET")
	r.Handle("/products/{slug}/feedback", authMiddleware(AddFeedbackHandler)).Methods("POST")

	//r.Handle("/get-token", GetTokenHandler).Methods("GET")
	// Наш API будет состоять из трех маршрутов
	// / status - мы будем звонить, чтобы убедиться, что наш API запущен и запущен
	// / products - который будет извлекать список продуктов, которые пользователь может оставить отзыв на
	// / products / {slug} / feedback - которая будет отображать отзывы пользователей о продуктах
	//	r.Handle("/status", NotImplemented).Methods("GET")

	// Наше приложение будет работать на порту 3000.
	// Здесь мы объявляем порт и передаем наш маршрутизатор.
	// Оберните функцию LoggingHandler вокруг нашего маршрутизатора,
	// чтобы регистратор был вызван первым в каждом запросе маршрута
	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		secret := []byte("{API_CLIENT_SECRET}")
		secretProvider := auth0.NewKeyProvider(secret)
		audience := []string{"{AUTH0_API_AUDIENCE}"}

		configuration := auth0.NewConfiguration(secretProvider, audience, "https://{AUTH0_DOMAIN}.auth0.com/", jose.HS256)
		validator := auth0.NewValidator(configuration)

		token, err := validator.ValidateRequest(r)

		if err != nil {
			fmt.Println(err)
			fmt.Println("Token is not valid:", token)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

/*
Обработчик статуса будет вызываться, когда пользователь вызывает маршрут / status
Он просто вернет строку с сообщением «API запущен и работает»,
*/
//var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte("API is up and running"))
//})

/*
Обработчик продукта вызывается, когда пользователь делает запрос GET конечной точке / products.
Этот обработчик вернет список продуктов, доступных пользователям для просмотра
*/
var ProductsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// Here we are converting the slice of products to JSON
	payload, _ := json.Marshal(products)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
})

/*
Обработчик обратной связи добавит положительную или отрицательную обратную связь к продукту
Обычно мы будем сохранять эти данные в базе данных, но для этой демонстрации мы подделаем ее
так что, пока запрос будет успешным, и мы можем сопоставить продукт с нашим каталогом продуктов
мы вернем статус ОК.
*/
var AddFeedbackHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var product Product
	vars := mux.Vars(r)
	slug := vars["slug"]

	for _, p := range products {
		if p.Slug == slug {
			product = p
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if product.Slug != "" {
		payload, _ := json.Marshal(product)
		w.Write([]byte(payload))
	} else {
		w.Write([]byte("Product Not Found"))
	}
})

/* Настройте глобальную строку для нашего тайного */
//var mySigningKey = []byte("secret")

/* Обработчики */
//var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//	/* Создать токен */
//	token := jwt.New(jwt.SigningMethodHS256)
//
//	// Создайте карту для хранения наших заявок
//	claims := token.Claims.(jwt.MapClaims)
//
//	/* Установите требования к токенам */
//	claims["admin"] = true
//	claims["name"] = "Ado Kukic"
//	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
//
//	/* Подпишите токен с нашей тайной */
//	tokenString, _ := token.SignedString(mySigningKey)
//
//	/* Наконец, напишите токен в окне браузера */
//	w.Write([]byte(tokenString))
//})
//
//// Здесь мы реализуем обработчик NotImplemented. Всякий раз, когда удаляется конечная точка API
//// мы просто вернем сообщение «Не реализовано»,
//var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte("Not Implemented"))
//})
//
//var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
//	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
//		return mySigningKey, nil
//	},
//	SigningMethod: jwt.SigningMethodHS256,
//})
