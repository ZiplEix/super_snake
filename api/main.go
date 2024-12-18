package main

import (
	"errors"
	"log"
	"os"

	"github.com/ZiplEix/super_snake/api/database"
	"github.com/ZiplEix/super_snake/api/hub"
	"github.com/ZiplEix/super_snake/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func checkEnv() error {
	log.Printf("Checking environment variables...")

	// application
	if _, ok := os.LookupEnv("PORT"); !ok {
		return errors.New("env var 'PORT' is not set")
	}
	if _, ok := os.LookupEnv("VERSION"); !ok {
		return errors.New("env var 'VERSION' is not set")
	}
	// cors
	if _, ok := os.LookupEnv("ALLOWED_ORIGINS"); !ok {
		return errors.New("env var 'ALLOWED_ORIGINS' is not set")
	}
	// database
	if _, ok := os.LookupEnv("POSTGRES_URL"); !ok {
		return errors.New("env var 'POSTGRES_URL' is not set")
	}
	// jwt
	if _, ok := os.LookupEnv("JWT_SECRET"); !ok {
		return errors.New("env var 'JWT_SECRET' is not set")
	}

	return nil
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	_ = godotenv.Load()

	err := checkEnv()
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	err = database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	err = database.Migrate()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("ALLOWED_ORIGINS"),
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))

	app.Use(logger.New(logger.Config{}))

	// hub := websocket.NewHub()
	// go hub.Run()

	hub.MainHub = hub.NewHub()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
