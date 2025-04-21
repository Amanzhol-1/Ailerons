package main

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"

	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"goGate/internal/auth/config"
	"goGate/internal/auth/delivery/http"
	"goGate/internal/auth/middleware"
	"goGate/internal/auth/repository"
	"goGate/internal/auth/service"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Almaty",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect DB: %v", err)
	}

	if err := db.AutoMigrate(
		&repository.UserModel{},
		&repository.CustomerProfileModel{},
		&repository.DriverProfileModel{},
	); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	userRepo := repository.NewGormUserRepo(db)
	custRepo := repository.NewGormCustomerProfileRepo(db)
	drvRepo := repository.NewGormDriverProfileRepo(db)

	authSvc := service.NewAuthService(userRepo, []byte(cfg.AuthSecret))
	profSvc := service.NewProfileService(custRepo, drvRepo, userRepo)

	e := echo.New()
	e.POST("/register", http.NewHandler(authSvc, profSvc).Register)
	e.POST("/login", http.NewHandler(authSvc, profSvc).Login)

	jwtMW := echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(cfg.AuthSecret),
		TokenLookup: "header:Authorization",
		ContextKey:  "user",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return &middleware.Claims{}
		},
	})

	grp := e.Group("/profile", jwtMW)
	h := http.NewHandler(authSvc, profSvc)
	grp.GET("", h.GetProfile)
	grp.POST("/customer", h.UpdateCustomer)
	grp.POST("/driver", h.UpdateDriver)

	addr := fmt.Sprintf(":%s", cfg.HTTPPort)
	log.Printf("Auth service listening on %s", addr)
	log.Fatal(e.Start(addr))
}
