package utils

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

func RandomUsername() string {
	adjectives := []string{"Quick", "Lazy", "Sleepy", "Clever", "Brave"}
	nouns := []string{"Fox", "Dog", "Cat", "Mouse", "Bird"}
	rand.Seed(uint64(time.Now().UnixNano())) // Khởi tạo nguồn ngẫu nhiên
	return fmt.Sprintf("%s_%s_%d", adjectives[rand.Intn(len(adjectives))], nouns[rand.Intn(len(nouns))], rand.Intn(1000))
}

// Hàm tạo email ngẫu nhiên
func RandomEmail() string {
	domains := []string{"example.com", "test.com", "demo.com"}
	username := RandomUsername() // Sử dụng hàm RandomUsername để tạo phần username cho email
	return fmt.Sprintf("%s@%s", username, domains[rand.Intn(len(domains))])
}

// Hàm tạo mật khẩu ngẫu nhiên
func RandomPassword() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+"
	passwordLength := 12
	password := make([]byte, passwordLength)
	rand.Seed(uint64(time.Now().UnixNano())) // Khởi tạo nguồn ngẫu nhiên

	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}
	return string(password)
}

// Hàm tạo ngày sinh ngẫu nhiên
func RandomDOB() string {
	rand.Seed(uint64(time.Now().UnixNano()))
	year := rand.Intn(30) + 1975 // Tạo năm ngẫu nhiên từ 1970 đến 1999
	month := rand.Intn(12) + 1   // Tháng từ 1 đến 12
	day := rand.Intn(28) + 1     // Ngày từ 1 đến 28 (để đơn giản)
	dob := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return dob.Format("2006-01-02")
}
