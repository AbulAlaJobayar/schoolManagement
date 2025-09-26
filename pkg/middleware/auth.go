package middleware

// import (
// 	"strings"
// 	"github.com/gofiber/fiber/v2"
// )

// type Role string

// const (
// 	ADMIN  Role = "ADMIN"
// 	USER   Role = "USER"
// 	SUPERADMIN Role = "SUPERADMIN"
// )

// // Auth middleware
// func Auth(allowedRoles ...Role) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		authHeader := c.Get("Authorization")
// 		if authHeader == "" {
// 			return &AppError{StatusCode: fiber.StatusUnauthorized, Message: "User Unauthorized"}
// 		}

// 		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

// 		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 			// Secret key
// 			return []byte(utils.AppConfig.JwtAccessSecret), nil
// 		})
// 		if err != nil || !token.Valid {
// 			return &AppError{StatusCode: fiber.StatusUnauthorized, Message: "Invalid token"}
// 		}

// 		claims, ok := token.Claims.(jwt.MapClaims)
// 		if !ok || claims["id"] == nil {
// 			return &AppError{StatusCode: fiber.StatusUnauthorized, Message: "Invalid token claims"}
// 		}

// 		userID := claims["id"].(string)

// 		// Fetch user from DB
// 		user := new(models.User)
// 		if err := utils.DB.First(user, "id = ? AND is_deleted = false", userID).Error; err != nil {
// 			return &utils.AppError{StatusCode: fiber.StatusNotFound, Message: "User not found"}
// 		}

// 		if user.AccountStatus == "BLOCK" || user.AccountStatus == "SUSPEND" || user.AccountStatus == "IN_PROGRESS" {
// 			return &utils.AppError{StatusCode: fiber.StatusForbidden, Message: "User is not active"}
// 		}

// 		// Role check
// 		if len(allowedRoles) > 0 {
// 			allowed := false
// 			for _, r := range allowedRoles {
// 				if string(r) == user.Role {
// 					allowed = true
// 					break
// 				}
// 			}
// 			if !allowed {
// 				return &utils.AppError{StatusCode: fiber.StatusUnauthorized, Message: "You are not authorized"}
// 			}
// 		}

// 		// Attach user to locals
// 		c.Locals("user", user)

// 		return c.Next()
// 	}
// }
