package repositories

import (
	"address-book-server-v3/internal/common/fault"
	"address-book-server-v3/internal/models"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/samber/mo"
	"gorm.io/gorm"
)

type user = models.User

type UserRepo interface {
	Create(user *models.User) mo.Result[*user]
	FindByEmail(email string) mo.Result[*user]
	ExistsByEmail(email string) mo.Result[*bool]
	ExistsByID(userID uuid.UUID) mo.Result[*bool]
}

type userRepo struct {
	*RepoContext
}

func NewUserRepo(ctx *RepoContext) UserRepo {
	return &userRepo{
		ctx,
	}
}

func (repo *userRepo) Create(u *models.User) mo.Result[*user] {
	if err := repo.db.Create(&u).Error; err != nil {
		return mo.Err[*user](fault.DBError(err))
	}

	return mo.Ok(u)
}

func (repo *userRepo) FindByEmail(email string) mo.Result[*user] {
	var u user
	if err := repo.db.Where("email = ?", email).First(&u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return mo.Err[*user](fault.RecordNotFound(map[string]any{
				"email": email,
			}, err))
		}

		return mo.Err[*user](fault.DBError(err))
	}

	return mo.Ok(&u)
}

func (repo *userRepo) ExistsByEmail(email string) mo.Result[*bool] {
	var count int64

	if err := repo.db.Model(&user{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return mo.Err[*bool](fault.DBError(err))
	}

	// if count <= 0 {
	// 	return mo.Err[*bool](fault.RecordNotFound(map[string]any{
	// 		"email": email,
	// 	}, nil))
	// }

	return mo.Ok(lo.ToPtr(count > 0))
}

func (repo *userRepo) ExistsByID(userID uuid.UUID) mo.Result[*bool] {
	var count int64

	if err := repo.db.Model(&user{}).Where("id = ?", userID[:]).Count(&count).Error; err != nil {
		return mo.Err[*bool](fault.DBError(err))
	}

	if count <= 0 {
		return mo.Err[*bool](fault.RecordNotFound(map[string]any{
			"user_id": userID,
		}, nil))
	}

	return mo.Ok(lo.ToPtr(count > 0))
}