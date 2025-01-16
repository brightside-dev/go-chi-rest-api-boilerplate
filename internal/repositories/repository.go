package repositories

import "context"

type Repository[T any] interface {
	GetSearchableFields() map[string]bool
	ScanRow(ctx context.Context, row interface{}) (*T, error)

	Insert(ctx context.Context, entity *T) error
	Update(ctx context.Context, entity *T) error
	Delete(ctx context.Context, id int) error
	FindOneById(ctx context.Context, id int) (*T, error)
	FindAll(ctx context.Context, limit int, offset int) ([]*T, error)
	FindBy(ctx context.Context, field string, value interface{}, offset int) (*T, error)
}

type FieldType int

const (
	IntType FieldType = iota
	StringType
	BoolType
)

type FieldMeta struct {
	Allowed bool
	Type    interface{}
}
