package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
	"webApp/api"
	"webApp/internal/types"
)

// PageData содержит данные для шаблона
type PageData struct {
	Title     string
	Header    string
	Subheader string
	Content   string
	types.IPResponse
}

func main() {
    // Загружаем шаблоны
	tmpl, err := template.ParseGlob("templates/*.html")
	if err != nil {
		fmt.Printf("Ошибка загрузки шаблонов: %v\n", err)
		os.Exit(1)
	}
	// Обработчик для главной страницы
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		 // Перезагружаем шаблоны при каждом запросе (только для разработки!)
        tmpl = template.Must(template.ParseGlob("templates/*.html"))
		// Устанавливаем заголовок Content-Type
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// 1. Инициализация клиента
	client := api.NewClient("https://dummyjson.com")

	// 2. Создание контекста с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 3. Выполнение запроса
	ipInfo, err := client.GetIP(ctx)
	if err != nil {
		log.Fatalf("Error getting IP: %v", err)
	}

		// Данные для шаблона
		data := PageData{
			Title:     "Мое Go приложение",
			Header:    "Добро пожаловать в мое приложение на Go!",
			Subheader: "Это простой пример веб-страницы",
			Content:   "Сервер написан на языке GO!!!",
			IPResponse: *ipInfo,
		}

	// 4. Вывод результата
	fmt.Printf("IP: %s\nUser Agent: %s\n", ipInfo.IP, ipInfo.UserAgent)

	// Рендерим шаблон с данными. Выполняем шаблон	
		err = tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Запускаем сервер на порту 8080
	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}