package main

import (
	"fmt"
	"os"
	"presensi/config"
	"presensi/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	app := fiber.New()
	cfg := config.InitConfig()
	logger := logrus.New()

	// Inisialisasi database
	db, err := config.InitDB(*cfg)
	if err != nil {
		logger.Fatal("Gagal menghubungkan ke database:", err)
		return
	}

	// Migrasi misal sudah di run, comment aja
	// if err := migration.RunMigrations(db); err != nil {
	// 	logger.Fatal("Gagal menjalankan migrasi:", err)
	// 	return
	// }

	// Inisialisasi JWT
	jwt := utils.NewJWT(cfg.JWT_Secret)

	// Bootstrap aplikasi
	config.Bootstrap(&config.BootstrapConfig{
		DB:  db,
		App: app,
		Log: logger,
		JWT: jwt,
	})

	// Start Server di port 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Info("Server berjalan di port", port)
	app.Listen(fmt.Sprintf(":%s", port))
}
