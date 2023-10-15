package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Инициализируем срез содержащий пути к двум файлам. Обратите внимание, что
	// файл home.page.tmpl должен быть *первым* файлом в срезе.
	files := []string{
		"C:\\Users\\aliha\\OneDrive\\Рабочий стол\\Fourth\\Go\\Project\\snippetbox\\ui\\html\\base.layout.tmpl",
		//"C:\\Users\\aliha\\OneDrive\\Рабочий стол\\Fourth\\Go\\project\\snippetbox\\ui\\html\\base.layout.tmpl",
		//"C:\\Users\\aliha\\OneDrive\\Рабочий стол\\Fourth\\Go\\project\\snippetbox\\ui\\html\\footer.partial.tmpl",
	}

	// Используем функцию template.ParseFiles() для чтения файлов шаблона.
	// Если возникла ошибка, мы запишем детальное сообщение ошибки и
	// используя функцию http.Error() мы отправим пользователю
	// ответ: 500 Internal Server Error (Внутренняя ошибка на сервере)
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	//// Используем функцию template.ParseFiles() для чтения файла шаблона.
	//// Если возникла ошибка, мы запишем детальное сообщение ошибки и
	//// используя функцию http.Error() мы отправим пользователю
	//// ответ: 500 Internal Server Error (Внутренняя ошибка на сервере)
	//ts, err := template.ParseFiles("./ui/html/home.page.tmpl")
	//if err != nil {
	//	log.Println(err.Error())
	//	http.Error(w, "Internal Server Error", 500)
	//	return
	//}
	//
	//// Затем мы используем метод Execute() для записи содержимого
	//// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	//// возможность отправки динамических данных в шаблон.
	//err = ts.Execute(w, nil)
	//if err != nil {
	//	log.Println(err.Error())
	//	http.Error(w, "Internal Server Error", 500)
	//}

	//w.Write([]byte("Привет из Snippetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Отображение определенной заметки с ID %d...", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Метод не дозволен", 405)
		return
	}

	w.Write([]byte("Создание новой заметки..."))
}
