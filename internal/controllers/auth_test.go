package controllers

import (
	"address-book-server-v3/internal/models"
	"address-book-server-v3/internal/services"
	// "fmt"
	// "address-book-server-v3/internal/services"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func (s *ControllerTestSuite) TestRegisterUser() {
	s.Run("it registers a user", func() {
		user := models.RegisterRequest{
			Body: &models.RegisterRequestBody{
				Email:    "yash11@gmail.com",
				Password: "Password@123",
			},
		}

		body, _ := json.Marshal(user.Body)

		// if err != nil {
		// 	fmt.Println("=======================> err :", err.Error())
		// }

		// fmt.Println("=======================> Body :", string(body))
		req, _ := http.NewRequest(http.MethodPost, "/api/v3/auth/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		s.router.ServeHTTP(w, req)

		s.Equal(http.StatusOK, w.Code)
		var response struct {
			Data models.RegisterResponse `json:"data"`
		}

		json.Unmarshal(w.Body.Bytes(), &response)
		s.Equal("yash11@gmail.com", response.Data.Email)
	})
}

func (s *ControllerTestSuite) TestLoginUser() {
	s.Run("it logins a user", func() {

		cmd := services.NewRegisterUserCmd("yash12@gmail.com", "Password@123")
		cmd.Execute(services.NewCommandContext(s.app, services.NewMockRequestCtx(), s.app.GetLogger()))

		user := models.LoginRequest{
			Body: &models.LoginRequestBody{
				Email:    "yash12@gmail.com",
				Password: "Password@123",
			},
		}
		body, _ := json.Marshal(user.Body)
		// if err != nil {
		// 	fmt.Println("=======================> err :", err.Error())
		// }

		// fmt.Println("=======================> Body :", string(body))
		req, _ := http.NewRequest(http.MethodPost, "/api/v3/auth/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		s.router.ServeHTTP(w, req)

		s.Equal(http.StatusOK, w.Code)
		var response struct {
			Data models.LoginResponse `json:"data"`
		}
		json.Unmarshal(w.Body.Bytes(), &response)
		s.NotEmpty(response.Data.Token)

		// User 2 for Test
		// fmt.Println("============> test2")
		user2 := models.LoginRequest{
			Body: &models.LoginRequestBody{
				Email:    "yash12@gmil.com",
				Password: "Password@123",
			},
		}
		
		body2, _ := json.Marshal(user2.Body)
		req2, _ := http.NewRequest(http.MethodPost, "/api/v3/auth/login", bytes.NewBuffer(body2))
		req2.Header.Set("Content-Type", "application/json")
		w2 := httptest.NewRecorder()
		s.router.ServeHTTP(w2, req2)

		s.Equal(http.StatusNotFound, w2.Code)
		var response2 struct {
			Data models.LoginResponse `json:"data"`
		}
		json.Unmarshal(w2.Body.Bytes(), &response2)
		s.Empty(response2.Data.Token)
	})
}
