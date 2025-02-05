package internal

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func DecodeToken(c *fiber.Ctx) (err error) {
	authorizationHeader := c.Get("Authorization")
	if !strings.Contains(authorizationHeader, "Bearer") {
		return ReturnTheResponse(c, true, 401, "Unauthorized", nil)
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (_ interface{}, err error) {
		return nil, err
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return ReturnTheResponse(c, true, int(500), "Failed to parse data token", nil)
	}

	username := claims["username"].(string)
	role := claims["role"].(string)

	c.Locals("username", username)
	c.Locals("role", role)

	return c.Next()
}
