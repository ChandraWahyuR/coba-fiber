package utils

import (
	"presensi/constant"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

const (
	MinCost     int = 4
	MaxCost     int = 31
	DefaultCost int = 14
)

func ValidateEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`)
	return regex.MatchString(email)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), DefaultCost)
	if err != nil {
		return "", constant.ErrHashPassword
	}
	return string(bytes), nil
}

func ValidatePassword(password string) (string, error) {
	// ^ => diawali
	// (?=.*[a-z]) => satu huruf kecil
	// (?=.*[A-Z]) => satu huruf besar
	// (?=.*\d) => satu angka
	// (?=.*[@$!%*?&#]) => satu karakter spesial
	// [A-Za-z\d@$!%*?&#]{8,} => karakter yang ada harus sesuai ketentuan diatas
	// .{8,} => minimal 8
	// $ => diakhiri
	// passwordValid := regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&#])[A-Za-z\d@$!%*?&#]{8,}$`).MatchString(password)
	containsLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	containsUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	containsNumber := regexp.MustCompile(`\d`).MatchString(password)
	containsSpecial := regexp.MustCompile(`[@$!%*?&#]`).MatchString(password)

	// Panjang password 8 - 16
	if len(password) < 8 || len(password) > 16 {
		return "", constant.ErrLenPassword
	}
	if !containsLower || !containsUpper || !containsNumber || !containsSpecial {
		return "", constant.ErrInvalidPassword
	}
	return password, nil
}
