package bootstrap

import (
	"io"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/kooroshh/fiber-boostrap/app/ws"
	"github.com/kooroshh/fiber-boostrap/pkg/database"
	"github.com/kooroshh/fiber-boostrap/pkg/env"
	"github.com/kooroshh/fiber-boostrap/pkg/router"
	"go.elastic.co/apm"
)

func NewApplication() *fiber.App {
	env.SetupEnvFile()
	SetupLogfile()

	database.SetupDatabase()
	database.SetupMongoDB()

	apm.DefaultTracer.Service.Name = "simple-messaging-app"
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	app.Use(recover.New()) // middleware untuk recover jd appnya ga mati kalau panic
	app.Use(logger.New())  // middleware logger
	app.Get("/dashboard", monitor.New())

	go ws.ServeWSMessaging(app) // biar berjalan di thread terpisah

	router.InstallRouter(app)

	return app
}

func SetupLogfile() {
	err := os.MkdirAll("./logs", os.ModePerm) // Create the folder if it doesn't exist
	if err != nil {
		log.Fatal("Failed to create logs directory: ", err)
	}
	logFile, err := os.OpenFile("./logs/simple_messaging_app.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal("Failed to open log file: ", err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
}
