package middleware

import (
	"context"
	"encoding/json"
	jwtToken "go-mux-api/pkg/jwt"
	"net/http"
)

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// func Auth(w http.ResponseWriter, r *http.Request) error {
// 	token := r.Header.Get("token")

// 	if token == "" {
// 		// return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 		// 	"message": "unauthorized",
// 		// })
// 		fmt.Println("unauthorized")
// 	}

// 	// _, err := utils.VerifyToken(token)
// 	claims, err := jwtToken.DecodeToken(token)

// 	if err != nil {
// 		// return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 		// 	"message": "unauthorized",
// 		// })
// 		fmt.Println("unauthorized")
// 	}

// 	// role := claims["role"].(string)
// 	// if role != "admin" {
// 	// 	return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
// 	// 		"message": "forbiden access",
// 	// 	})
// 	// }

// 	ctx.Locals("userId", claims)
// 	r.Response.Location("userId", claims)

// 	return Next
// }
func Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("token")

		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			response := Result{Code: http.StatusUnauthorized, Message: "unauthorized"}
			json.NewEncoder(w).Encode(response)
			return
		}

		// _, err := jwtToken.VerifyToken(token)
		claims, err := jwtToken.DecodeToken(token)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			response := Result{Code: http.StatusUnauthorized, Message: "unauthorized"}
			json.NewEncoder(w).Encode(response)
			return
		}
		// fmt.Println(payload)
		// ctx := context.WithValue(r.Context(), "payloadUser", payload)
		// next.ServeHTTP(w, r.WithContext(ctx))

		ctx := context.WithValue(context.Background(), "userInfo", claims)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
