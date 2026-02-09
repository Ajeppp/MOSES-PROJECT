package main

import (
	"fmt"
	"log"

	"moses/database"
	"moses/routes"
	"moses/scheduler"
	"moses/seeders"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
	============================
	  CONFIG

============================
*/
var (
	DB_HOST = "localhost"
	DB_PORT = "5432"
	DB_USER = "postgres"
	DB_PASS = "jef"
	DB_NAME = "moses"
)

/* ============================
   DB BOOTSTRAP (POSTGRES)
============================ */

func ensureDatabase() {
	// connect ke default postgres db
	dsnRoot := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=postgres port=%s sslmode=disable",
		DB_HOST, DB_USER, DB_PASS, DB_PORT,
	)

	rootDB, err := gorm.Open(postgres.Open(dsnRoot), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Gagal connect ke Postgres server:", err)
	}

	var exists bool
	checkQuery := `
		SELECT EXISTS(
			SELECT 1 
			FROM pg_database 
			WHERE datname = ?
		)
	`

	if err := rootDB.Raw(checkQuery, DB_NAME).Scan(&exists).Error; err != nil {
		log.Fatal("❌ Gagal cek database:", err)
	}

	if !exists {
		log.Println("🧱 Database belum ada, creating:", DB_NAME)
		createQuery := fmt.Sprintf(`CREATE DATABASE "%s";`, DB_NAME)

		if err := rootDB.Exec(createQuery).Error; err != nil {
			log.Fatal("❌ Gagal create database:", err)
		}

		log.Println("✅ Database created:", DB_NAME)
	} else {
		log.Println("✅ Database exists:", DB_NAME)
	}
}

/* ============================
   MAIN
============================ */

func main() {
	// ========================
	// 1. DB BOOTSTRAP
	// ========================
	ensureDatabase()

	// ========================
	// 2. CONNECT DB
	// ========================
	database.Connect()

	// ========================
	// 3. SEEDER
	// ========================
	seeders.SeedAll()

	// ========================
	// 4. ROUTER
	// ========================
	r := gin.Default()

	// ✅ CORS CONFIG
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	routes.RegisterRoutes(r)

	// ========================
	// 5. SCHEDULER
	// ========================
	go scheduler.GenerateSchedule(2, 2026)
	// async biar server ga nge-freeze

	// ========================
	// 6. RUN SERVER
	// ========================
	log.Println("🚀 Server running on :8080")
	r.Run(":8080")
}
