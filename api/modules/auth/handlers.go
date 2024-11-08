package auth

import (
	"net/http"
	"piccolo/api/shared"

	"github.com/labstack/echo/v4"
)

type RegisterReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRes struct {
	Id int `json:"id"`
}

func (m *AuthModule) registerHandler(c echo.Context) error {
	registerReq := new(RegisterReq)

	var err error

	if err = c.Bind(registerReq); err != nil {
		return c.JSON(http.StatusBadRequest, shared.SuccessRes{
			Success: false,
			Message: "Invalid input",
		})
	}

	hash, err := hashPassword(registerReq.Password)
	if err != nil {
		m.server.Logger.Error("unexpected error", err)
		return c.JSON(http.StatusInternalServerError, shared.SuccessRes{
			Success: false,
			Message: "Unexpected error",
		})
	}

	return c.JSON(http.StatusOK, "TODO")
}

// func (server *Server) CreateUser(c echo.Context) error {
// 	userReqBody := new(UserRequestBody)
//
// 	// Parse the JSON body into the UserRequestBody struct
// 	if err := c.Bind(userReqBody); err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Please provide the correct input!!"})
// 	}
//
// 	// Hash the password
// 	hashPassword, err := getHashPassword(userReqBody.Password)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Something bad happened on the server :("})
// 	}
//
// 	// Insert the user into the database
// 	query := `INSERT INTO User (email, hash) VALUES (?, ?)`
// 	result, err := server.DB.Exec(query, userReqBody.Email, hashPassword)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Something bad happened on the server :("})
// 	}
//
// 	// Get the last inserted ID
// 	recordId, _ := result.LastInsertId()
//
// 	// Prepare the response
// 	response := Response{
// 		Id: int(recordId),
// 	}
//
// 	// Return a JSON response with the created user's ID
// 	return c.JSON(http.StatusOK, response)
// }
