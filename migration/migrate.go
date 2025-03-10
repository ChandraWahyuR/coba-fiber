package migration

import (
	"database/sql"
	"fmt"
	"os"
)

// CheckIfTableExists memeriksa apakah tabel sudah ada di database
func CheckIfTableExists(db *sql.DB, tableName string) (bool, error) {
	var exists bool
	query := `SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = DATABASE() AND table_name = ?`
	err := db.QueryRow(query, tableName).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("gagal memeriksa tabel %s: %v", tableName, err)
	}
	return exists, nil
}

// 002 buat alter table yang sudah ada
// ALTER TABLE users
// ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
// RunMigrations akan mengeksekusi semua file SQL di folder migrations
func RunMigrations(db *sql.DB) error {
	files := []string{
		"migration/001_create_users.sql",
		// "migration/0002_add_columns.sql",
	}

	// Mengecek apakah tabel 'users' sudah ada
	tableExists, err := CheckIfTableExists(db, "users")
	if err != nil {
		return err
	}

	// Jika tabel 'users' sudah ada, skip file migrasi yang membuat tabel
	if tableExists {
		fmt.Println("Tabel 'users' sudah ada, melewati migrasi pembuatan tabel.")
		files = append(files[1:], files[0])
	}

	for _, file := range files {
		query, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("gagal membaca file migrasi %s: %v", file, err)
		}

		_, err = db.Exec(string(query))
		if err != nil {
			return fmt.Errorf("gagal menjalankan migrasi %s: %v", file, err)
		}
		fmt.Println("Berhasil menjalankan migrasi:", file)
	}
	return nil
}
