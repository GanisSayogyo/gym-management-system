package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv membaca konfigurasi dari file .env.
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("File .env tidak ditemukan, menggunakan environment sistem")
	}
}

// GetEnv mengambil nilai environment berdasarkan nama key.
// Jika tidak ditemukan, gunakan nilai default.
func GetEnv(key string, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}
