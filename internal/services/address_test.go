package services

import (
	"address-book-server-v3/internal/common/types"
	"address-book-server-v3/internal/models"

	"github.com/google/uuid"
)

func createTestUser(s *ServiceTestSuit, email string) types.UserId {
	cmdCtx := NewCommandContext(s.suite, NewMockRequestCtx(), s.suite.GetLogger())

	cmd := NewRegisterUserCmd(email, "Yash1234...")
	result := cmd.Execute(cmdCtx)

	s.True(result.IsOk())
	user := result.MustGet()
	s.Equal(email, user.Email)
	return types.UserId(user.Id)
}

func createTestAddress(s *ServiceTestSuit, userId types.UserId, email string) types.AddressId {
	req := models.CreateAddressRequest{
		Body: &models.CreateAddressRequestBody{
			FirstName:    "y",
			LastName:     "s",
			Email:        email,
			Phone:        "9999999999",
			AddressLine1: "Line 1",
			AddressLine2: "Line 2",
			City:         "Surat",
			State:        "Gujarat",
			Country:      "India",
			Pincode:      "395006",
		},
	}

	cmd := NewCreateAddrCmd(req, userId)
	result := cmd.Execute(NewCommandContext(
		s.suite,
		NewMockRequestCtx(),
		s.suite.GetLogger(),
	))

	s.True(result.IsOk())
	addr := result.MustGet()
	s.Equal(email, addr.Email)
	return types.AddressId(addr.Id)
}

func (s *ServiceTestSuit) TestCreateAddress() {
	// create user
	userId := createTestUser(s, "yash1@gmail.com")

	// create address
	_ = createTestAddress(s, userId, "y11s11@gmail.com")
}

func (s *ServiceTestSuit) TestGetAllAddresses() {
	// create user
	userId := createTestUser(s, "yash1@gmail.com")

	// create address
	_ = createTestAddress(s, userId, "y11s11@gmail.com")

	cmd := NewGetAllAddrCmd(userId)
	result := cmd.Execute(NewCommandContext(
		s.suite,
		NewMockRequestCtx(),
		s.suite.GetLogger(),
	))

	s.True(result.IsOk())
	addrs := result.MustGet()
	s.GreaterOrEqual(len(addrs.Addresses), 1)
}

func (s *ServiceTestSuit) TestGetAddressByID() {
	// create user
	userId1 := createTestUser(s, "yash1@gmail.com")

	// create address
	addrId := createTestAddress(s, userId1, "y11s11@gmail.com")

	cmd1 := NewGetByIdAddrCmd(addrId, userId1)
	result1 := cmd1.Execute(NewCommandContext(
		s.suite,
		NewMockRequestCtx(),
		s.suite.GetLogger(),
	))

	s.True(result1.IsOk())
	s.Equal("y11s11@gmail.com", result1.MustGet().Email)

	userId2 := createTestUser(s, "yash2@gmail.com")

	cmd := NewGetByIdAddrCmd(types.AddressId(uuid.New()), userId2)
	result := cmd.Execute(NewCommandContext(
		s.suite,
		NewMockRequestCtx(),
		s.suite.GetLogger(),
	))

	s.True(result.IsError())
}

func (s *ServiceTestSuit) TestUpdateAddress() {
	// create user
	userId := createTestUser(s, "yash1@gmail.com")

	// create address
	addrId := createTestAddress(s, userId, "y11s11@gmail.com")

	// Update
	newCity := "Baroda"
	updateReq := models.UpdateAddressRequest{
		Body: &models.UpdateAddressRequestBody{
			City: &newCity,
		},
	}

	updateCmd := NewUpdateAddrCmd(addrId, userId, updateReq)
	result := updateCmd.Execute(NewCommandContext(
		s.suite,
		NewMockRequestCtx(),
		s.suite.GetLogger(),
	))

	s.True(result.IsOk())
	s.Equal("Baroda", result.MustGet().City)
}

func (s *ServiceTestSuit) TestDeleteAddress() {
	// create user
	userId := createTestUser(s, "yash1@gmail.com")

	// create address
	addrId := createTestAddress(s, userId, "y11s11@gmail.com")

	// delete
	cmd := NewDeleteAddrCmd(addrId, userId)
	result := cmd.Execute(NewCommandContext(
		s.suite,
		NewMockRequestCtx(),
		s.suite.GetLogger(),
	))

	s.True(result.IsOk())
	s.Equal("Address deleted successfully", result.MustGet().Message)
}

func (s *ServiceTestSuit) TestFilterAddresses() {
	// create user
	userId := createTestUser(s, "yash1@gmail.com")

	// create address
	_ = createTestAddress(s, userId, "y11s11@gmail.com")
	_ = createTestAddress(s, userId, "y12s12@gmail.com")

	query := models.FilterAddrQuery{
		Body: &models.FilterAddrQueryBody{
			Page:  1,
			Limit: 10,
		},
	}
	cmd := NewFilterAddrCmd(query, userId)
	result := cmd.Execute(NewCommandContext(
		s.suite,
		NewMockRequestCtx(),
		s.suite.GetLogger(),
	))

	s.True(result.IsOk())
	resp := result.MustGet()
	s.GreaterOrEqual(len(resp.Data), 2)
	s.GreaterOrEqual(resp.Total, int64(2))

	query2 := models.FilterAddrQuery{
		Body: &models.FilterAddrQueryBody{
			City:  "Surat",
			Page:  1,
			Limit: 10,
		},
	}
	cmd2 := NewFilterAddrCmd(query2, userId)
	result2 := cmd2.Execute(NewCommandContext(s.suite, NewMockRequestCtx(), s.suite.GetLogger()))

	s.True(result.IsOk())
	resp2 := result2.MustGet()
	s.Equal(int64(2), resp2.Total)
	s.Equal("Surat", resp2.Data[0].City)

	query3 := models.FilterAddrQuery{
		Body: &models.FilterAddrQueryBody{
			Search: "y11s11",
			Page:   1,
			Limit:  10,
		},
	}
	cmd3 := NewFilterAddrCmd(query3, userId)
	result3 := cmd3.Execute(NewCommandContext(s.suite, NewMockRequestCtx(), s.suite.GetLogger()))

	s.True(result.IsOk())
	resp3 := result3.MustGet()
	s.Equal(int64(1), resp3.Total)
}

func (s *ServiceTestSuit) TestExportAddressesAsync() {
	// create user
	userId := createTestUser(s,"yash1@gmail.com")

	// create address so export has data
	createTestAddress(s, userId, "y11s11@gmail.com")

	cmd := NewExportAsyncAddrCmd(
		userId,
		[]string{"email", "city"},
		"receiver@gmail.com",
	)

	result := cmd.Execute(NewCommandContext(
		s.suite,
		NewMockRequestCtx(),
		s.suite.GetLogger(),
	))

	s.True(result.IsOk())
	resp := result.MustGet()
	s.Equal("Export started", resp.Message)
}
