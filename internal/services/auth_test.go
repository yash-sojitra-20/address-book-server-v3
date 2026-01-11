package services

import (
	"address-book-server-v3/internal/common/utils"
	"address-book-server-v3/internal/models"
	"address-book-server-v3/internal/repositories"

	"github.com/google/uuid"
)

func (s *ServiceTestSuit) TestRegisterUser() {
	s.Run("Register new user", func() {
		cmdCtx := NewCommandContext(s.suite, NewMockRequestCtx(), s.suite.GetLogger())

		cmd1 := NewRegisterUserCmd("yash1@gmail.com", "Yash1234...")
		result1 := cmd1.Execute(cmdCtx)

		s.True(result1.IsOk())
		registeredUser := result1.MustGet()
		s.Equal("yash1@gmail.com", registeredUser.Email)

		cmd2 := NewRegisterUserCmd("yash1@gmail.com", "Yash1234...")
		result2 := cmd2.Execute(cmdCtx)

		s.True(result2.IsError())
	})
}

func (s *ServiceTestSuit) TestLoginUser() {
	id := uuid.New()
	email := "yash1@gmail.com"
	hashed, _ := utils.HashPassword("Yash1234...").Get()
	user := &models.User{
		Id:       id[:],
		Email:    email,
		Password: *hashed,
	}
	repo := repositories.NewUserRepo(repositories.NewRepoContext(s.db, s.suite.GetLogger()))
	s.True(repo.Create(user).IsOk())

	s.Run("Login user", func() {
		cmdCtx := NewCommandContext(s.suite, NewMockRequestCtx(), s.suite.GetLogger())

		cmd1 := NewLoginUserCmd("yash1@gmail.com", "Yash1234...")
		result1 := cmd1.Execute(cmdCtx)

		s.True(result1.IsOk())
		loginData := result1.MustGet()
		s.NotEmpty(loginData.Token)

		cmd2 := NewLoginUserCmd("yash1@gmail.com", "....Yash1234")
		result2 := cmd2.Execute(cmdCtx)

		s.True(result2.IsError())

		cmd3 := NewLoginUserCmd("yash11@gmail.com", "Yash1234...")
		result3 := cmd3.Execute(cmdCtx)

		s.True(result3.IsError())

	})
}
