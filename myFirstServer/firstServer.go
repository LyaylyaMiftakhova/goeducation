package myFirstServer

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var helloMap = make(map[string]int)

func MyFirstServer() {
	// Создаем новый роутер с использованием go-chi
	r := chi.NewRouter()

	// Добавляем middleware для логирования
	r.Use(middleware.Logger)

	// Регистрируем обработчики для каждой из трех ручек
	r.Get("/hello/{name}", sayHello)
	r.Get("/hello/count", getCount)
	r.Put("/hello/{deleteCount}", updateCount)

	// Запускаем сервер на порту 3000
	fmt.Println("Server is running on :3000")
	http.ListenAndServe(":3000", r)
}

// Обработчик для ручки GET /hello/{name}
func sayHello(w http.ResponseWriter, r *http.Request) {
	// Получаем значение параметра {name} из URL
	name := chi.URLParam(r, "name")

	// Увеличиваем счетчик для данного имени
	helloMap[name]++
	// Получаем текущее количество запросов для данного имени
	count := helloMap[name]

	// Формируем JSON-ответ
	response := map[string]interface{}{"Hello": name, "Count": count}
	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "application/json")
	// Кодируем и отправляем JSON-ответ
	json.NewEncoder(w).Encode(response)
}

// Обработчик для ручки GET /hello/count
func getCount(w http.ResponseWriter, r *http.Request) {
	// Вычисляем общее количество запросов
	var totalCount int
	for _, count := range helloMap {
		totalCount += count
	}

	// Формируем JSON-ответ с общим количеством запросов
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"TotalCount": totalCount})
}

// Обработчик для ручки PUT /hello/{deleteCount}
func updateCount(w http.ResponseWriter, r *http.Request) {
	// Получаем значение параметра {deleteCount} из URL
	deleteCountParam := chi.URLParam(r, "deleteCount")
	// Преобразуем его в число
	deleteCount, err := strconv.Atoi(deleteCountParam)
	if err != nil {
		// В случае ошибки возвращаем ошибку BadRequest
		http.Error(w, "Invalid deleteCount parameter", http.StatusBadRequest)
		return
	}

	// Вычисляем общее количество запросов
	var totalCount int
	for _, count := range helloMap {
		totalCount += count
	}

	// Уменьшаем счетчики для всех имен пропорционально
	for name := range helloMap {
		helloMap[name] = helloMap[name] * (totalCount - deleteCount) / totalCount
	}

	// Отправляем успешный ответ
	w.WriteHeader(http.StatusOK)
}
