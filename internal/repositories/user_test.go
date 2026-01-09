package repositories

// import (
// 	"address-book-server-v3/internal/models"
// 	// "address-book-server-v3/internal/common/types"
// 	// "go/types"

// 	"github.com/google/uuid"
// )

// func (suite *RepositoryTestSuite) TestCreateUser() {
// 	repoCtx := NewRepoContext(suite.db, suite.suite.GetLogger())
// 	repo := NewUserRepo(repoCtx)
// 	id := uuid.New()

// 	user := &models.User{
// 		Id:       id[:],
// 		Email:    "yash@gmail.com",
// 		Password: "Yash1234...",
// 	}

// 	suite.Run("creates a new user", func() {
// 		result := repo.Create(user)
// 		suite.True(result.IsOk())
// 		createdUser := result.MustGet()
// 		suite.Equal(user.Id, createdUser.Id)
// 		suite.Equal(user.Email, createdUser.Email)
// 	})
// }

// func (suite *RepositoryTestSuite) TestFindByEmail() {
// 	repoCtx := NewRepoContext(suite.db, suite.suite.GetLogger())
// 	repo := NewUserRepo(repoCtx)
// 	id := uuid.New()

// 	user := &models.User{
// 		Id:       id[:],
// 		Email:    "yash@gmail.com",
// 		Password: "Yash1234...",
// 	}

// 	repo.Create(user)

// 	email := "yash@gmail.com"

// 	suite.Run("Find user by email", func() {
// 		result := repo.FindByEmail(email)
// 		suite.True(result.IsOk())
// 		createdUser := result.MustGet()
// 		suite.Equal(user.Id, createdUser.Id)
// 		suite.Equal(user.Email, createdUser.Email)
// 	})
// }

// func (suite *RepositoryTestSuite) TestExistsByEmail() {
// 	repoCtx := NewRepoContext(suite.db, suite.suite.GetLogger())
// 	repo := NewUserRepo(repoCtx)
// 	id := uuid.New()

// 	user := &models.User{
// 		Id:       id[:],
// 		Email:    "1@gmail.com",
// 		Password: "Yash1234...",
// 	}

// 	repo.Create(user)

// 	email := "1@gmail.com"

// 	suite.Run("Find user by email", func() {
// 		result := repo.ExistsByEmail(email)
// 		suite.True(result.IsOk())
// 		isExist := result.MustGet()
// 		suite.Equal(true, *isExist)
// 	})
// }

// // func (suite *RepositoryTestSuite) TestExistsByID() {
// // 	repoCtx := NewRepoContext(suite.db, suite.suite.GetLogger())
// // 	repo := NewUserRepo(repoCtx)
// // 	id := uuid.New()

// // 	user := &models.User{
// // 		Id:       id[:],
// // 		Email:    "yash1@gmail.com",
// // 		Password: "Yash1234...",
// // 	}

// // 	repo.Create(user)

// // 	suite.Run("Find user by email", func() {
// // 		result := repo.ExistsByID(types.UserId(id))
// // 		suite.True(result.IsOk())
// // 		isExist := result.MustGet()
// // 		suite.Equal(true, *isExist)
// // 	})
// // }