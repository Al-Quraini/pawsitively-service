package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	alphabet    = "abcdefghijklmnopqrstuvwxyz"
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberBytes = "0123456789"
)

func init() {
	rand.Seed(time.Now().UnixMilli())
}

// RandomInt generate a random integer
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomName() string {
	return RandomString(6)
}

func randStringBytes(n int, allowedBytes string) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = allowedBytes[rand.Intn(len(allowedBytes))]
	}
	return string(b)
}

func RandEmail() string {
	username := randStringBytes(10, letterBytes+numberBytes)
	domain := randStringBytes(8, letterBytes)
	return fmt.Sprintf("%s@%s.com", username, domain)
}

func RandomAge() int64 {
	return RandomInt(0, 20)
}

func RandomGender() string {
	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(2)
	if randomInt == 0 {
		return "Male"
	} else {
		return "Female"
	}
}

func RandomAnimal() string {
	// List of animal types
	animalTypes := []string{"Dog", "Cat", "Horse", "Bird", "Fish", "Hamster", "Rabbit", "Snake", "Turtle", "Lizard", "Ferret"}

	// Initialize random seed with current time
	rand.Seed(time.Now().UnixNano())

	// Generate a random index within the range of the animalTypes slice
	randomIndex := rand.Intn(len(animalTypes))

	// Return the animal type at the random index
	return animalTypes[randomIndex]
}

func RandomImageUrl() string {
	// List of possible image file extensions
	fileExtensions := []string{".jpg", ".jpeg", ".png", ".gif"}

	// Random integer between 1000 and 9999
	randomNum := rand.Intn(9000) + 1000

	// Initialize random seed with current time
	rand.Seed(time.Now().UnixNano())

	// Random index within the range of the fileExtensions slice
	randomIndex := rand.Intn(len(fileExtensions))

	// Return the random URL
	return "https://example.com/images/" + fmt.Sprint(randomNum) + fileExtensions[randomIndex]
}
