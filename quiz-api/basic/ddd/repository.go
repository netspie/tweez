package ddd

type Repository[T Entity] interface {
	Get(id string) (T, bool)
	GetMany() ([]T, bool)
	Add(item T) bool
	Update(item T) bool
	Delete(item T) bool
}
