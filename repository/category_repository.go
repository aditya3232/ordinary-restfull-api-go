package repository

import (
	"context"
	"database/sql"
	"ordinary-restfull-api-go/model/domain"
)

/*
	- selalu buat interface, (kontrak)
	- digunakan untuk mempermudah, agar bisa tau detail dari kontrak seperti apa, menggunakan struct apa
*/

type CategoryRepository interface {
	Create(ctx context.Context, tx sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx sql.Tx, category domain.Category)
	FindById(ctx context.Context, tx sql.Tx, categoryId int) domain.Category
	FindAll(ctx context.Context, tx sql.Tx) []domain.Category
}
