package main

import (
	"fmt"
	"goGate/internal/auth/repository"
	"goGate/internal/auth/service"
	"log"
	"net/http"

	httpDelivery "goGate/internal/auth/delivery/http"
)

func main() {
	userRepo := repository.NewInMemoryUserRepo()

	authService := service.NewAuthService(userRepo)

	handler := httpDelivery.NewHandler(authService)

	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/welcome", handler.Welcome)

	fmt.Println("Сервер запущен на порту :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
