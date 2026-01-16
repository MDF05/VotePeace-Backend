package controllers

import (
	"strconv"
	"strings"
	"time"
	"votepeace/database"
	"votepeace/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret" // In production, use os.Getenv("SECRET_KEY")

func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		NIK:      data["nik"],
		Name:     data["name"],
		Password: string(password),
		Role:     "USER",
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "NIK already registered or invalid data",
		})
	}

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User
	database.DB.Where("nik = ?", data["nik"]).First(&user)

	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}

	claims := jwt.MapClaims{
		"iss": strconv.Itoa(int(user.ID)),
		"exp": time.Now().Add(time.Hour * 24).Unix(), // 1 day
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(SecretKey))

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	return c.JSON(fiber.Map{
		"status":  200,
		"succes":  true,
		"message": "Login Success",
		"content": fiber.Map{
			"token": t,
			"user":  user,
		},
	})
}

func Check(c *fiber.Ctx) error {
	// Token should be in Authorization header: Bearer <token> or Body/Query if needed.
	// But frontend sends as Bearer token.
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	claims := token.Claims.(jwt.MapClaims)
	issuer := claims["iss"].(string)

	var user models.User
	database.DB.Where("id = ?", issuer).First(&user)

	return c.JSON(fiber.Map{
		"content": user,
	})
}
