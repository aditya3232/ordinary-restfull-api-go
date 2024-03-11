package service

import (
	"context"
	"ordinary-restfull-api-go/model/web"
)

/*
	- service nanti akan memanggil repository untuk akses database
	- service adalah bisnis logic, yang menghubungkan repository, logic, request & response
	- logic biasanya berisi, trasaction database, add request to data model table, panggil repository, dan beberapa logic seperti pengecekan data, dll.
	- transaction ada di level service, karena di service bisa saja memanggil beberapa repository
*/

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}
