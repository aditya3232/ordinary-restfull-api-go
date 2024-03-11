package service

import (
	"context"
	"database/sql"
	"ordinary-restfull-api-go/helper"
	"ordinary-restfull-api-go/model/domain"
	"ordinary-restfull-api-go/model/web"
	"ordinary-restfull-api-go/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
}

/*
	- service update & delete ada logic pengecekan id
	- kalau service findById tidak usah
*/

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	dataCategory := domain.Category{
		Name: request.Name,
	}

	createCategory := service.CategoryRepository.Create(ctx, tx, dataCategory)

	return helper.ToCategoryResponse(createCategory)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	checkedIdCategory, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	dataCategory := domain.Category{
		Id:   checkedIdCategory.Id,
		Name: request.Name,
	}

	updateCategory := service.CategoryRepository.Update(ctx, tx, dataCategory)

	return helper.ToCategoryResponse(updateCategory)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	checkedIdCategory, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	service.CategoryRepository.Delete(ctx, tx, checkedIdCategory.Id)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	findByIdCategory, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	return helper.ToCategoryResponse(findByIdCategory)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}
