package main

import (
	"crypto/sha1"
	"fmt"
	"os"
)

func hashMySQLPassword(password string) string {
	// Get user input (password)
	fmt.Print("New MySQL Password: ")
	userEnteredPassword := readPassword()

	// Hash the password twice using SHA-1
	hashedSHA1 := sha1.Sum([]byte(userEnteredPassword))
	hashedSHA1Twice := sha1.Sum(hashedSHA1[:])

	// Convert to hexadecimal and prepend "*"
	hashedPassword := "*" + fmt.Sprintf("%x", hashedSHA1Twice)

	return hashedPassword
}

func readPassword() string {
	// Read password from user input (you can customize this part)
	// For example, you can use a library like "github.com/howeyc/gopass"
	// to securely read the password without echoing it to the terminal.
	// For simplicity, I'll use plain text input here.
	var password string
	fmt.Fscanf(os.Stdin, "%s", &password)
	return password
}

func main() {
	newPassword := "user-entered-password"
	hashedMySQLPassword := hashMySQLPassword(newPassword)
	fmt.Println("Hashed MySQL password:", hashedMySQLPassword)
}
