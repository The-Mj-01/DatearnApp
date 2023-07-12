package hasher

import "golang.org/x/crypto/bcrypt"

// PASSWORD_COST_LENGTH defines length of password hashing
const PASSWORD_COST_LENGTH int = 14

// Make hashes given password
func Make(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PASSWORD_COST_LENGTH)
	return string(bytes), err
}

// Check if two give passwords match or not
func Check(password string, hashedPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password))
	return err == nil
}
