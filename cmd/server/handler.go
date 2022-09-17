package main

import "net/http"

type Handler struct {
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/check":
		h.check(w, r)
	case "/reset-bucket":
		h.resetBucket(w, r)
	case "/add-to-blacklist":
		h.addToBlackList(w, r)
	case "/remove-from-blacklist":
		h.removeFromBlackList(w, r)
	case "/add-to-whitelist":
		h.addToWhiteList(w, r)
	case "/remove-from-whitelist":
		h.removeFromWhiteList(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (h *Handler) check(w http.ResponseWriter, r *http.Request) {
	// Попытка авторизации
	//Запрос:
	//* login
	//* password
	//* ip
	//
	//Ответ:
	//* ok (true/false) - сервис должен возвращать ok=true, если считает что запрос нормальный
	//и ok=false, если считает что происходит bruteforce.
}

func (h *Handler) resetBucket(w http.ResponseWriter, r *http.Request) {
	//
	// Сброс bucket
	//Должен очистить bucket-ы соответствующие переданным login и ip.
	//* login
	//* ip
	//
}

func (h *Handler) addToBlackList(w http.ResponseWriter, r *http.Request) {
	// Добавление IP в blacklist
	//* подсеть (IP + маска)
}

func (h *Handler) removeFromBlackList(w http.ResponseWriter, r *http.Request) {
	// Удаление IP из blacklist
	//* подсеть (IP + маска)
}
func (h *Handler) addToWhiteList(w http.ResponseWriter, r *http.Request) {
	// Добавление IP в whitelist
	//* подсеть (IP + маска)
}

func (h *Handler) removeFromWhiteList(w http.ResponseWriter, r *http.Request) {
	// Удаление IP из whitelist
	//* подсеть (IP + маска)
}
