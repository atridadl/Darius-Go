package lib

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func GetJWT(c *fiber.Ctx) error {

	var user string = c.FormValue("user")
	var pass string = c.FormValue("pass")

	hashed, _ := bcrypt.GenerateFromPassword([]byte(pass), 8)

	pass = string(hashed)

	// Throws Unauthorized error if no user and pass is sent
	if user == "" || pass == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Throws Unauthorized error if user or pass is incorrect
	// This is where you would perform a check against your database for the username and password
	if false {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"username": user,
		"admin":    true,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	// CHANGE THIS SECRET! This is used for demo purposes only! You should never hardcode your secret in your code like this!
	t, err := token.SignedString([]byte("CHANGEME"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if c.IsFromLocal() {
		c.Cookie(&fiber.Cookie{
			Name:   "token",
			Value:  t,
			MaxAge: 3600,
		})
		return c.JSON(fiber.Map{"success": "true"})
	} else {
		c.Cookie(&fiber.Cookie{
			Name:     "token",
			Value:    t,
			Expires:  time.Now().Add(time.Hour * 72),
			HTTPOnly: true,
			Secure:   true,
			MaxAge:   3600,
		})

		return c.JSON(fiber.Map{"success": "true"})
	}

}
