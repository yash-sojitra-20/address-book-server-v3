package repositories

import (
	"address-book-server-v3/internal/common/fault"
	"address-book-server-v3/internal/models"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/samber/mo"
	"gorm.io/gorm"
)

type address = models.Address
type anyStringMap = map[string]interface{}

type addressData struct {
	Addresses []address
	Total     int64
}

type AddressRepo interface {
	Create(address *address) mo.Result[*address]
	FindByUser(userID uuid.UUID) mo.Result[*[]address]
	FindByID(id uuid.UUID, userID uuid.UUID) mo.Result[*address]
	Update(address *address) mo.Result[*address]
	Delete(address *address) mo.Result[*address]
	FindAllForExport(fields []string, userID uuid.UUID) mo.Result[*[]anyStringMap]
	FindFiltered(userId uuid.UUID, listAddressQuery *models.FilterAddrQuery) mo.Result[*addressData]
}

type addressRepo struct {
	*RepoContext
}

func NewAddressRepo(ctx *RepoContext) *addressRepo {
	return &addressRepo{
		ctx,
	}
}

func (repo *addressRepo) Create(a *address) mo.Result[*address] {
	if err := repo.db.Create(a).Error; err != nil {
		return mo.Err[*address](fault.DBError(err))
	}

	return mo.Ok(a)
}

func (repo *addressRepo) FindByUser(userID uuid.UUID) mo.Result[*[]address] {
	var a []address

	if err := repo.db.Where("user_id = ?", userID[:]).Find(&a).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return mo.Err[*[]address](fault.RecordNotFound(map[string]any{
				"user_id": userID,
			}, err))
		}

		return mo.Err[*[]address](fault.DBError(err))
	}

	return mo.Ok(&a)
}

func (repo *addressRepo) FindByID(id uuid.UUID, userID uuid.UUID) mo.Result[*address] {
	var a address

	if err := repo.db.Where("id = ? AND user_id = ?", id[:], userID[:]).First(&a).Error; err != nil {

		if err == gorm.ErrRecordNotFound {

			return mo.Err[*address](fault.RecordNotFound(map[string]any{
				"id": id,
			}, err))
		}

		return mo.Err[*address](fault.DBError(err))
	}

	return mo.Ok(&a)
}

func (repo *addressRepo) Update(a *address) mo.Result[*address] {

	if err := repo.db.Save(a).Error; err != nil {
		return mo.Err[*address](fault.DBError(err))
	}

	return mo.Ok(a)
}

func (repo *addressRepo) Delete(a *address) mo.Result[*string] {
	if err := repo.db.Delete(a).Error; err != nil {
		return mo.Err[*string](fault.DBError(err))
	}

	return mo.Ok(lo.ToPtr("Address deleted successfully"))
}

func (repo *addressRepo) FindAllForExport(fields []string, userID uuid.UUID) mo.Result[*[]anyStringMap] {
	var results []anyStringMap

	if err := repo.db.Model(&address{}).Select(fields).Where("user_id = ?", userID[:]).Find(&results).Error; err != nil {
		return mo.Err[*[]anyStringMap](fault.DBError(err))
	}

	return mo.Ok(&results)
}

func (repo *addressRepo) FindFiltered(userId uuid.UUID, listAddressQuery *models.FilterAddrQuery) mo.Result[*addressData] {
	offset := (listAddressQuery.Page - 1) * listAddressQuery.Limit

	query := repo.db.Model(&address{}).Where("user_id = ?", userId[:])

	// SEARCH (across multiple fields)
	if listAddressQuery.Search != "" {
		like := "%" + listAddressQuery.Search + "%"
		query = query.Where(`
			first_name ILIKE ? OR 
			last_name ILIKE ? OR 
			email ILIKE ? OR
			phone ILIKE ? OR
			city ILIKE ? OR
			state ILIKE ? OR
			country ILIKE ?`,
			like, like, like, like, like, like, like,
		)
	}

	// FILTERS
	if listAddressQuery.City != "" {
		query = query.Where("city ILIKE ?", listAddressQuery.City)
	}
	if listAddressQuery.State != "" {
		query = query.Where("state ILIKE ?", listAddressQuery.State)
	}
	if listAddressQuery.Country != "" {
		query = query.Where("country ILIKE ?", listAddressQuery.Country)
	}

	// fmt.Println(query)

	var total int64
	query.Count(&total) // get total records

	// fmt.Println(total)

	// PAGINATION
	var addresses []models.Address
	if err := query.Limit(listAddressQuery.Limit).Offset(offset).Order("created_at DESC").Find(&addresses).Error; err != nil {
		return mo.Err[*addressData](fault.DBError(err))
	}

	// fmt.Println("inside repo:",err)

	return mo.Ok(&addressData{Addresses: addresses, Total: total})
}
