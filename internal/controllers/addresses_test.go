package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	// "address-book-server-v3/internal/common/fault"
	"address-book-server-v3/internal/models"
	"address-book-server-v3/internal/services"
)

func (s *ControllerTestSuite) loginAndGetToken(email string) string {
	cmd := services.NewRegisterUserCmd(email, "Password@123")
	cmd.Execute(services.NewCommandContext(
		s.app,
		services.NewMockRequestCtx(),
		s.app.GetLogger(),
	))

	body, _ := json.Marshal(models.LoginRequestBody{
		Email:    email,
		Password: "Password@123",
	})

	req, _ := http.NewRequest(http.MethodPost, "/api/v3/auth/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	s.router.ServeHTTP(w, req)

	var resp struct {
		Data models.LoginResponse `json:"data"`
	}
	json.Unmarshal(w.Body.Bytes(), &resp)

	return resp.Data.Token
}

func authRequest(method, url, token string, body []byte) *http.Request {
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	return req
}

func (s *ControllerTestSuite) TestCreateAddress() {
	s.Run("creates address", func() {
		token := s.loginAndGetToken("create@addr.com")

		reqBody := models.CreateAddressRequestBody{
			FirstName:    "John",
			LastName:     "Doe",
			Email:        "john@doe.com",
			Phone:        "9999999999",
			AddressLine1: "Street 1",
			City:         "Ahmedabad",
			State:        "Gujarat",
			Country:      "India",
			Pincode:      "380001",
		}

		body, _ := json.Marshal(reqBody)
		req := authRequest(http.MethodPost, "/api/v3/addresses", token, body)
		w := httptest.NewRecorder()

		s.router.ServeHTTP(w, req)

		s.Equal(http.StatusOK, w.Code)

		var resp struct {
			Data models.AddressResponse `json:"data"`
		}
		json.Unmarshal(w.Body.Bytes(), &resp)

		s.Equal("John", resp.Data.FirstName)
		s.Equal("Ahmedabad", resp.Data.City)
	})
}

func (s *ControllerTestSuite) TestListAllAddresses() {
	s.Run("lists addresses", func() {
		token := s.loginAndGetToken("list@addr.com")

		createBody := models.CreateAddressRequestBody{
			FirstName:    "A",
			Email:        "a@b.com",
			AddressLine1: "Line 1",
		}
		body, _ := json.Marshal(createBody)

		s.router.ServeHTTP(
			httptest.NewRecorder(),
			authRequest(http.MethodPost, "/api/v3/addresses", token, body),
		)

		req := authRequest(http.MethodGet, "/api/v3/addresses", token, nil)
		w := httptest.NewRecorder()

		s.router.ServeHTTP(w, req)

		s.Equal(http.StatusOK, w.Code)

		var resp struct {
			Data models.ListAddressResponse `json:"data"`
		}
		json.Unmarshal(w.Body.Bytes(), &resp)

		s.Len(resp.Data.Addresses, 1)
	})
}

func (s *ControllerTestSuite) TestGetAddressById() {
	s.Run("gets address by id", func() {
		token := s.loginAndGetToken("get@addr.com")

		createBody := models.CreateAddressRequestBody{
			FirstName:    "Fetch",
			Email:        "fetch@addr.com",
			AddressLine1: "Line",
		}
		body, _ := json.Marshal(createBody)

		w1 := httptest.NewRecorder()
		s.router.ServeHTTP(w1,
			authRequest(http.MethodPost, "/api/v3/addresses", token, body),
		)

		var createResp struct {
			Data models.AddressResponse `json:"data"`
		}
		json.Unmarshal(w1.Body.Bytes(), &createResp)

		req := authRequest(
			http.MethodGet,
			"/api/v3/addresses/"+createResp.Data.Id,
			token,
			nil,
		)

		w := httptest.NewRecorder()
		s.router.ServeHTTP(w, req)

		s.Equal(http.StatusOK, w.Code)

		req = authRequest(
			http.MethodGet,
			"/api/v3/addresses/123",
			token,
			nil,
		)
		w = httptest.NewRecorder()
		s.router.ServeHTTP(w, req)

		s.Equal(http.StatusBadRequest, w.Code)
	})
}

func (s *ControllerTestSuite) TestUpdateAddress() {
	s.Run("updates address", func() {
		token := s.loginAndGetToken("update@addr.com")

		createBody := models.CreateAddressRequestBody{
			FirstName:    "Old",
			Email:        "old@addr.com",
			AddressLine1: "Line",
		}
		body, _ := json.Marshal(createBody)

		w1 := httptest.NewRecorder()
		s.router.ServeHTTP(w1,
			authRequest(http.MethodPost, "/api/v3/addresses", token, body),
		)

		var createResp struct {
			Data models.AddressResponse `json:"data"`
		}
		json.Unmarshal(w1.Body.Bytes(), &createResp)

		newCity := "Baroda"
		updateBody := models.UpdateAddressRequestBody{
			City: &newCity,
		}

		upd, _ := json.Marshal(updateBody)
		req := authRequest(
			http.MethodPut,
			"/api/v3/addresses/"+createResp.Data.Id,
			token,
			upd,
		)

		w := httptest.NewRecorder()
		s.router.ServeHTTP(w, req)

		s.Equal(http.StatusOK, w.Code)
	})
}

func (s *ControllerTestSuite) TestDeleteAddress() {
	s.Run("deletes address", func() {
		token := s.loginAndGetToken("delete@addr.com")

		createBody := models.CreateAddressRequestBody{
			FirstName:    "Del",
			Email:        "del@addr.com",
			AddressLine1: "Line",
		}
		body, _ := json.Marshal(createBody)

		w1 := httptest.NewRecorder()
		s.router.ServeHTTP(w1,
			authRequest(http.MethodPost, "/api/v3/addresses", token, body),
		)

		var createResp struct {
			Data models.AddressResponse `json:"data"`
		}
		json.Unmarshal(w1.Body.Bytes(), &createResp)

		req := authRequest(
			http.MethodDelete,
			"/api/v3/addresses/"+createResp.Data.Id,
			token,
			nil,
		)

		w := httptest.NewRecorder()
		s.router.ServeHTTP(w, req)

		s.Equal(http.StatusOK, w.Code)
	})
}

func (s *ControllerTestSuite) TestExportAddresses() {
	s.Run("exports addresses", func() {
		token := s.loginAndGetToken("export@addr.com")

		reqBody := models.ExportAddressRequestBody{
			Fields: []string{"first_name", "email"},
			Email:  "export@addr.com",
		}

		body, _ := json.Marshal(reqBody)
		req := authRequest(http.MethodPost, "/api/v3/addresses/export", token, body)
		w := httptest.NewRecorder()

		s.router.ServeHTTP(w, req)

		s.Equal(http.StatusOK, w.Code)
	})
}

func (s *ControllerTestSuite) TestAddressUnauthorized() {
	s.Run("fails without token", func() {
		req, _ := http.NewRequest(http.MethodGet, "/api/v3/addresses", nil)
		w := httptest.NewRecorder()

		s.router.ServeHTTP(w, req)

		s.Equal(http.StatusUnauthorized, w.Code)
	})
}
