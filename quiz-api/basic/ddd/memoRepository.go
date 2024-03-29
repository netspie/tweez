package ddd

import "quiz/basic/mapx"

type MemoRepository[T Entity] struct {
	empty T
	items map[string]T
}

func NewMemoRepository[T Entity](empty T) MemoRepository[T] {
	return MemoRepository[T]{
		empty: empty,
		items: map[string]T{},
	}
}

func (r *MemoRepository[T]) Get(id string) (T, bool) {
	item, res := r.items[id]
	if !res {
		return r.empty, res
	}

	return item, res
}

func (r *MemoRepository[T]) GetMany() ([]T, bool) {
	return mapx.Map(r.items, func(key string, v T) T { return v }), true
}

func (r *MemoRepository[T]) Add(item T) bool {
	id := item.GetId()
	_, res := r.items[id]
	if res {
		return false
	}

	r.items[id] = item
	return true
}

func (r *MemoRepository[T]) Update(item T) bool {
	id := item.GetId()
	_, res := r.items[id]
	if !res {
		return false
	}

	r.items[id] = item
	return true
}

func (r *MemoRepository[T]) Delete(item T) bool {
	id := item.GetId()
	_, res := r.items[id]
	if !res {
		return false
	}

	delete(r.items, id)
	return true
}
