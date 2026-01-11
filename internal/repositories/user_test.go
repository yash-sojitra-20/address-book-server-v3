package repositories

import (
	// "address-book-server-v3/internal/common/types"
	"address-book-server-v3/internal/models"
	// "fmt"

	"github.com/google/uuid"
)

func (suite *RepositoryTestSuite) TestCreateUser() {
	repoCtx := NewRepoContext(suite.db, suite.suite.GetLogger())
	repo := NewUserRepo(repoCtx)

	id := uuid.New()
	user := &models.User{
		Id:       id[:],
		Email:    "yash1@gmail.com",
		Password: "Yash1234...",
	}

	suite.Run("creates a new user", func() {
		result := repo.Create(user)
		suite.True(result.IsOk())
		createdUser := result.MustGet()
		suite.Equal(user.Id, createdUser.Id)
		suite.Equal(user.Email, createdUser.Email)
	})
}

func (suite *RepositoryTestSuite) TestFindByEmail() {
	repoCtx := NewRepoContext(suite.db, suite.suite.GetLogger())
	repo := NewUserRepo(repoCtx)

	id := uuid.New()
	user := &models.User{
		Id:       id[:],
		Email:    "yash2@gmail.com",
		Password: "Yash1234...",
	}
	repo.Create(user)

	email := "yash2@gmail.com"
	suite.Run("Find user by email", func() {
		result := repo.FindByEmail(email)
		suite.True(result.IsOk())
		createdUser := result.MustGet()
		suite.Equal(user.Id, createdUser.Id)
		suite.Equal(user.Email, createdUser.Email)
	})

	wrongEmail := "yash22@gmail.com"
	suite.Run("Find user by email", func() {
		result := repo.FindByEmail(wrongEmail)
		suite.Equal(false, result.IsOk())
	})
}

func (suite *RepositoryTestSuite) TestExistsByEmail() {
	repoCtx := NewRepoContext(suite.db, suite.suite.GetLogger())
	repo := NewUserRepo(repoCtx)

	id := uuid.New()
	user := &models.User{
		Id:       id[:],
		Email:    "yash3@gmail.com",
		Password: "Yash1234...",
	}
	repo.Create(user)

	email := "yash3@gmail.com"
	suite.Run("Find user exist by email", func() {
		result := repo.ExistsByEmail(email)
		suite.True(result.IsOk())
		isExist := result.MustGet()
		suite.Equal(true, *isExist)
	})

	wrongEmail := "yash33@gmail.com"
	suite.Run("Find user exist by email", func() {
		result := repo.ExistsByEmail(wrongEmail)
		// fmt.Println("==============> result:", result, user.Email, wrongEmail)
		suite.Equal(true, result.IsOk())
		isExist := result.MustGet()
		suite.Equal(false, *isExist)
	})
}

func (suite *RepositoryTestSuite) TestExistsByID() {
	repoCtx := NewRepoContext(suite.db, suite.suite.GetLogger())
	repo := NewUserRepo(repoCtx)

	id := uuid.New()
	user := &models.User{
		Id:       id[:],
		Email:    "yash4@gmail.com",
		Password: "Yash1234...",
	}
	repo.Create(user)

	// fmt.Printf("==============> Inserted ID: %x\n", user.Id)
	// fmt.Printf("==============> Original ID: %x\n", id[:])

	suite.Run("Find user by Id", func() {
		result := repo.ExistsByID(id)
		// fmt.Println("===========> result:", result)
		suite.True(result.IsOk())
		isExist := result.MustGet()
		suite.Equal(true, *isExist)
	})

	wrondId := uuid.New()
	suite.Run("Find user by Id", func() {
		result := repo.ExistsByID(wrondId)
		// fmt.Println("===========> result:", result)
		suite.True(result.IsOk())
		isExist := result.MustGet()
		suite.Equal(false, *isExist)
	})
}
