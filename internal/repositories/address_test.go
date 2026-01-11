package repositories

import (
	"address-book-server-v3/internal/models"
	// "fmt"

	"github.com/google/uuid"
)

func createTestUser(s *RepositoryTestSuite, email string) (*models.User, uuid.UUID) {
	userID := uuid.New()
	user := &models.User{
		Id:       userID[:],
		Email:    email,
		Password: "Yash1234...",
	}
	repo := NewUserRepo(NewRepoContext(s.db, s.suite.GetLogger()))
	repo.Create(user)
	return user, userID
}

func createTestAddress(
	s *RepositoryTestSuite,
	userID uuid.UUID,
	email string,
) (*models.Address, uuid.UUID) {

	addressID := uuid.New()
	address := &models.Address{
		Id:     addressID[:],
		UserId: userID[:],

		FirstName: "y",
		LastName:  "s",
		Email:     email,
		Phone:     "9999999999",

		AddressLine1: "Street 1",
		City:         "Surat",
		State:        "Gujarat",
		Country:      "India",
		Pincode:      "395006",
	}

	repo := NewAddressRepo(NewRepoContext(s.db, s.suite.GetLogger()))
	repo.Create(address)

	return address, addressID
}

func (s *RepositoryTestSuite) TestCreateAddress() {
	repoCtx := NewRepoContext(s.db, s.suite.GetLogger())
	userRepo := NewUserRepo(repoCtx)
	addressRepo := NewAddressRepo(repoCtx)

	userId := uuid.New()
	user := &models.User{
		Id:       userId[:],
		Email:    "yash1@gmail.com",
		Password: "Yash1234...",
	}
	userRepo.Create(user)

	addressId := uuid.New()
	address := &models.Address{
		Id:     addressId[:],
		UserId: userId[:],

		FirstName: "y11",
		LastName:  "y11",
		Email:     "y11s11@gmail.com",
		Phone:     "9876543210",

		AddressLine1: "123 - New street",
		AddressLine2: "near city mall",
		City:         "surat",
		State:        "gujrat",
		Country:      "india",
		Pincode:      "395006",
	}

	s.Run("creates a new address", func() {
		result := addressRepo.Create(address)
		s.True(result.IsOk())
		createdAddress := result.MustGet()

		// fmt.Printf("==============> Inserted addressId: %x\n", createdAddress.Id)
		// fmt.Printf("==============> Original addressId: %x\n", addressId[:])
		// fmt.Printf("==============> Inserted userId: %x\n", createdAddress.UserId)
		// fmt.Printf("==============> Original userId: %x\n", userId[:])

		s.Equal(addressId, uuid.UUID(createdAddress.Id))
		s.Equal(userId, uuid.UUID(createdAddress.UserId))
	})
}

func (s *RepositoryTestSuite) TestFindByUser() {
	_, userID := createTestUser(s, "yash1@gmail.com")
	createTestAddress(s, userID, "y11s11@gmail.com")
	createTestAddress(s, userID, "y12s12@gmail.com")

	repo := NewAddressRepo(NewRepoContext(s.db, s.suite.GetLogger()))

	s.Run("Find address by user", func() {
		result := repo.FindByUser(userID)
		
		s.True(result.IsOk())
		addresses := result.MustGet()
		s.Len(*addresses, 2)
	})

	_, wrongUserID := createTestUser(s, "ys@gmail.com")
	s.Run("Find address by user", func() {
		result := repo.FindByUser(wrongUserID)

		s.True(result.IsOk())
		addresses := result.MustGet()
		s.Len(*addresses, 0)
	})

}

func (s *RepositoryTestSuite) TestFindByID() {
	_, userID1 := createTestUser(s, "yash1@gmail.com")
	_, userID2 := createTestUser(s, "yash2@gmail.com")
	_, addressID := createTestAddress(s, userID1, "y11s11@gmail.com")

	repo := NewAddressRepo(NewRepoContext(s.db, s.suite.GetLogger()))

	s.Run("Find address by Id", func() {
		result := repo.FindByID(addressID, userID1)

		s.True(result.IsOk())
		address := result.MustGet()
		s.Equal(addressID, uuid.UUID(address.Id))
	})

	s.Run("Find address by Id", func() {
		result := repo.FindByID(addressID, userID2)

		s.True(result.IsError())
	})

}

func (s *RepositoryTestSuite) TestUpdateAddress() {
	_, userID := createTestUser(s, "yash1@gmail.com")
	address, _ := createTestAddress(s, userID, "y11s11@gmail.com")

	repo := NewAddressRepo(NewRepoContext(s.db, s.suite.GetLogger()))

	address.City = "Vadodara"
	s.Run("Update address", func() {
		result := repo.Update(address)

		s.True(result.IsOk())
		updated := result.MustGet()
		s.Equal("Vadodara", updated.City)
	})

}

func (s *RepositoryTestSuite) TestDeleteAddress() {
	_, userID := createTestUser(s, "yash1@gmail.com")
	address, addressID := createTestAddress(s, userID, "y11s11@gmail.com")

	repo := NewAddressRepo(NewRepoContext(s.db, s.suite.GetLogger()))

	s.Run("Delete address", func() {
		delResult := repo.Delete(address)

		s.True(delResult.IsOk())
		findResult := repo.FindByID(addressID, userID)
		s.True(findResult.IsError())
	})

}

func (s *RepositoryTestSuite) TestFindAllForExport() {
	_, userID := createTestUser(s, "yash1@gmail.com")
	createTestAddress(s, userID, "y11s11@gmail.com")
	createTestAddress(s, userID, "y12s12@gmail.com")

	repo := NewAddressRepo(NewRepoContext(s.db, s.suite.GetLogger()))

	s.Run("Find all for export", func() {
		fields := []string{"email", "city"}
		result := repo.FindAllForExport(fields, userID)

		s.True(result.IsOk())
		data := result.MustGet()
		s.Len(*data, 2)
	})

}

func (s *RepositoryTestSuite) TestFindFiltered() {
	_, userID := createTestUser(s, "yash1@gmail.com")
	createTestAddress(s, userID, "y11s11@gmail.com")
	createTestAddress(s, userID, "y12s12@gmail.com")

	repo := NewAddressRepo(NewRepoContext(s.db, s.suite.GetLogger()))

	query := &models.FilterAddrQuery{
		Body: &models.FilterAddrQueryBody{
			Page:   1,
			Limit:  10,
			Search: "y11s11",
		},
	}
	s.Run("Find filtered addresses", func() {
		result := repo.FindFiltered(userID, query)

		s.True(result.IsOk())
		data := result.MustGet()
		s.Len(data.Addresses, 1)
		s.Equal(int64(1), data.Total)
	})

	wrongQuery := &models.FilterAddrQuery{
		Body: &models.FilterAddrQueryBody{
			Page:   1,
			Limit:  10,
			Search: "notfound",
		},
	}
	s.Run("Find filtered addresses", func() {
		result := repo.FindFiltered(userID, wrongQuery)

		s.True(result.IsOk())
		data := result.MustGet()
		s.Len(data.Addresses, 0)
		s.Equal(int64(0), data.Total)
	})

}
