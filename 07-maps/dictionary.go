package maps

type Dictionary map[string]string
type DictionaryErr string

var (
	ErrNoItemDefined  = DictionaryErr("no item defined")
	ErrItemExists     = DictionaryErr("item already defined")
	ErrNoItemToUpdate = DictionaryErr("no item to update")
	ErrNoItemToDelete = DictionaryErr("no item to delete")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	val, ok := d[key]
	if !ok {
		return "", ErrNoItemDefined
	}
	return val, nil
}

func (d Dictionary) Add(key, val string) error {
	_, ok := d[key]
	if !ok {
		d[key] = val
		return nil
	}
	return ErrItemExists
}

func (d Dictionary) Update(key, val string) error {
	_, ok := d[key]
	if ok {
		d[key] = val
		return nil
	}
	return ErrNoItemToUpdate
}

func (d Dictionary) Delete(key string) error {
	_, ok := d[key]
	if ok {
		delete(d, key)
		return nil
	}
	return ErrNoItemToDelete
}
