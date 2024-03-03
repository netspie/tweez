package basic

type MemoRepository[T Entity] struct {
	items map[string]T
}

func (r *MemoRepository[T]) Get(id string) (T, bool) {
	item, res := r.items[id]
	return item, res
}

func (r *MemoRepository[T]) Add(item T) bool {
	id := item.Id()
	item, res := r.items[id]
	if res {
		return false
	}

	r.items[id] = item
	return true
}

func (r *MemoRepository[T]) Update(item T) bool {
	id := item.Id()
	item, res := r.items[id]
	if !res {
		return false
	}

	r.items[id] = item
	return true
}

func (r *MemoRepository[T]) Delete(item T) bool {
	id := item.Id()
	item, res := r.items[id]
	if !res {
		return false
	}

	r.items[id] = item
	return true
}
