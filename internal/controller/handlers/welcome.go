package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"forum/internal/controller"
	"forum/internal/tool/config"
)

type welcomeHandler struct {
	service controller.Authentication
}

func NewWelcomeHandler(service controller.Authentication) *welcomeHandler {
	log.Println("| | welcome handler is done!")
	return &welcomeHandler{
		service: service,
	}
}

func (h *welcomeHandler) WelcomePage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	// userId := r.Context().Value(ctx)
	// tmpl, _ := template.ParseFiles("./template/welcome.html")
	// tmpl.Execute(w, userId)
	userId := r.Context().Value(config.UserID)
	json.NewEncoder(w).Encode(userId)
	// w.Write([]byte(strconv.Itoa(userId)))
}

func (h *welcomeHandler) HomePage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("home page"))
}

func (h *welcomeHandler) Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("middleware")
		c, err := r.Cookie("session_token")
		// fmt.Println(c.Valid())
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				// next.ServeHTTP(w, r)
				return
			}
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		session, err := h.service.GetSession(r.Context(), c.Value)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		// session, err = h.service.UpdateSession(r.Context(), session)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		// http.SetCookie(w, &http.Cookie{
		// 	Name:    "session_token",
		// 	Value:   session.Token,
		// 	Expires: session.ExpireTime,
		// })

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), config.UserID, session.UserID)))
	}
}
