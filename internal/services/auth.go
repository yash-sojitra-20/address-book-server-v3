package services

import (
	"address-book-server-v3/internal/common/fault"
	"address-book-server-v3/internal/common/utils"
	"address-book-server-v3/internal/models"
	"address-book-server-v3/internal/repositories"

	"bitbucket.org/vayana/walt-go/command"
	"bitbucket.org/vayana/walt-gorm.go/transaction"
	"github.com/google/uuid"
	"github.com/samber/mo"
	"gorm.io/gorm"
)

type registerUserCmd struct {
	email    string
	password string
}

func NewRegisterUserCmd(email string, password string) *registerUserCmd {
	return &registerUserCmd{
		email:    email,
		password: password,
	}
}

func (cmd *registerUserCmd) Execute(c command.CmdContext) mo.Result[*models.UserCmdOutputData] {
	ctx := c.(CommandContext)

	operation := func(db *gorm.DB) mo.Result[*models.UserCmdOutputData] {
		logger := ctx.GetLogger()

		repoCtx := repositories.NewRepoContext(db, logger)
		userRepo := repositories.NewUserRepo(repoCtx)

		exist, err := userRepo.ExistsByEmail(cmd.email).Get()
		if err != nil {
			return mo.Err[*models.UserCmdOutputData](err)
		}
		if *exist {
			err := fault.UserExistWithEmailAlready(cmd.email, nil)
			return mo.Err[*models.UserCmdOutputData](err)
		}

		hashedPass, err := utils.HashPassword(cmd.password).Get()
		if err != nil {
			return mo.Err[*models.UserCmdOutputData](err)
		}

		user := &models.User{
			Email:    cmd.email,
			Password: *hashedPass,
		}

		createdUser, err := userRepo.Create(user).Get()
		if err != nil {
			return mo.Err[*models.UserCmdOutputData](err)
		}
		return mo.Ok(models.NewUserCmdOutputData(createdUser))
	}

	return transaction.DoInTransaction(ctx.GetDb(), operation)
}

type loginUserCmd struct {
	email    string
	password string
}

func NewLoginUserCmd(email string, password string) *loginUserCmd {
	return &loginUserCmd{
		email:    email,
		password: password,
	}
}

func (cmd *loginUserCmd) Execute(c command.CmdContext) mo.Result[*string] {
	ctx := c.(CommandContext)

	operation := func(db *gorm.DB) mo.Result[*string] {
		logger := ctx.GetLogger()

		repoCtx := repositories.NewRepoContext(db, logger)
		repo := repositories.NewUserRepo(repoCtx)

		user, err := repo.FindByEmail(cmd.email).Get()
		if err != nil {
			return mo.Err[*string](err)
		}

		isPassMatch, err := utils.ComparePassword(user.Password, cmd.password).Get()
		if err != nil {
			return mo.Err[*string](fault.InvalidPassword(nil))
		}

		var token *string
		if *isPassMatch {
			token, err = utils.GenerateToken(ctx.GetConfig().GetSecretKey(), uuid.UUID(user.Id), user.Email).Get()
		}
		if err != nil {
			return mo.Err[*string](err)
		}

		return mo.Ok(token)
	}
	
	return transaction.DoInTransaction(ctx.GetDb(), operation)
}
