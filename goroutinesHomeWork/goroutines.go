package goroutinesHomeWork

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func GoroutinesHomeWork() {
	urls := []string{
		"pornoHub.com",
		"XVideos.com",
		"Porno365.com",
		"Redtube.com",
	}

	// Получение статусов через параллельную проверку
	statuses := parallelCheck(urls)

	// Вывод результатов
	for _, status := range statuses {
		fmt.Println(status)
	}
}

func parallelCheck(urls []string) []int {
	// Создание WaitGroup для отслеживания завершения горутин
	var wg sync.WaitGroup

	// Срез для хранения статусов проверки
	statuses := make([]int, len(urls))

	// Канал для передачи статусов из горутин в основную горутину
	statusChan := make(chan int, len(urls))

	// Запуск горутин для проверки каждого URL
	for i, url := range urls {
		wg.Add(1)
		go func(i int, url string) {
			defer wg.Done()
			// Выполнение проверки и отправка результата в канал
			status := healthCheck(url)
			statusChan <- status
		}(i, url)
	}

	// Завершение работы всех горутин после их выполнения
	go func() {
		wg.Wait()
		close(statusChan)
	}()

	// Чтение результатов из канала и сохранение в срезе
	index := 0
	for status := range statusChan {
		statuses[index] = status
		index++
	}

	// Возвращение среза статусов
	return statuses
}

func healthCheck(url string) int {
	time.Sleep(2 * time.Second)

	response, err := http.Get("http://" + url)
	if err != nil {
		// В случае ошибки недоступности сайта по какой-то причине
		return http.StatusServiceUnavailable
	}
	defer response.Body.Close()

	// Проверка статуса ответа
	if response.StatusCode == http.StatusOK {
		// В случае успешной проверки, возвращаем статус StatusOK
		return http.StatusOK
	} else {
		// В случае неуспешной проверки, возвращаем статус StatusCode
		return response.StatusCode
	}
}
