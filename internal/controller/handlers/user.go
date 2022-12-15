package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"forum/internal/entity"
	"forum/internal/service"
)

type userHandler struct {
	service service.Authentication
}

func NewUserHandler(service service.Authentication) *userHandler {
	log.Println("| | user handler is done!")
	return &userHandler{
		service: service,
	}
}

func (h *userHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	switch r.Method {
	// case http.MethodGet:
	// 	tmpl, _ := template.ParseFiles("./template/signin.html")
	// 	tmpl.Execute(w, nil)
	// 	return
	case http.MethodPost:
		var user entity.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		session, err := h.service.SetSession(r.Context(), user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "session_token",
			Value:   session.Token,
			Expires: session.ExpireTime,
		})

		// json.NewEncoder(w).Encode(session.UserID)

	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (h *userHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	switch r.Method {
	// case http.MethodGet:
	// 	tmpl, _ := template.ParseFiles("./template/signup.html")
	// 	tmpl.Execute(w, nil)
	// 	return
	case http.MethodPost:
		var user entity.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		_, err := h.service.CreateUser(r.Context(), user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// w.Write([]byte(strconv.Itoa(int(id))))
		// json.NewEncoder(w).Encode(id)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (h *userHandler) LogOut(w http.ResponseWriter, r *http.Request) {
}
