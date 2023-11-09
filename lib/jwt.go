package lib

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func GenerateJWTHandler(c *fiber.Ctx) error {

	var user string = c.FormValue("user")
	var pass string = c.FormValue("pass")
	var json = c.FormValue("json")

	isJSON, isJSONInvalid := strconv.ParseBool(json)

	if isJSONInvalid != nil {
		isJSON = false
	}

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

		if isJSON {
			return c.JSON(fiber.Map{"success": "true"})
		} else {
			return c.SendString(`
		<p id="hello">😊 Success! 😊</p>
		<p>Try going go this page to test it out!:</p>
		<a
			class="text-white bg-gradient-to-r from-purple-500 to-pink-500 hover:bg-gradient-to-l focus:ring-4 focus:outline-none focus:ring-purple-200 dark:focus:ring-purple-800 font-medium rounded-lg text-sm px-5 py-2.5 text-center mr-2 mb-2"
			href="/restricted"
		>
			Shhh! Secret!
		</a>`)
		}
	} else {
		c.Cookie(&fiber.Cookie{
			Name:     "token",
			Value:    t,
			Expires:  time.Now().Add(time.Hour * 72),
			HTTPOnly: true,
			Secure:   true,
			MaxAge:   3600,
		})

		if isJSON {
			return c.JSON(fiber.Map{"success": "true"})
		} else {
			return c.SendString(`
		<p id="hello">😊 Success! 😊</p>
		<p>Try going go this page to test it out!:</p>
		<a
			class="text-white bg-gradient-to-r from-purple-500 to-pink-500 hover:bg-gradient-to-l focus:ring-4 focus:outline-none focus:ring-purple-200 dark:focus:ring-purple-800 font-medium rounded-lg text-sm px-5 py-2.5 text-center mr-2 mb-2"
			href="/restricted"
		>
			Shhh! Secret!
		</a>`)
		}
	}

}
