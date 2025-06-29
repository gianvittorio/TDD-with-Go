package maps

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrNotFound     = DictionaryErr("could not find the word you were looking for")
	ErrWordExists   = DictionaryErr("cannot add word because it already exists")
	ErrDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if !exists {
		return "", ErrNotFound
	}

	return value, nil
}

func (d Dictionary) Add(word, definition string) error {
	switch _, err := d.Search(word); err {
	case ErrNotFound:
		d[word] = definition
		return nil
	case nil:
		return ErrWordExists
	default:
		return err
	}
}

func (d Dictionary) Update(word, definition string) error {
	switch _, err := d.Search(word); err {
	case ErrNotFound:
		return ErrDoesNotExist
	case nil:
		d[word] = definition
		return nil
	default:
		return err
	}
}

func (d Dictionary) Delete(word string) error {
	switch _, err := d.Search(word); err {
		case ErrNotFound:
			return ErrDoesNotExist
		case nil:
			delete(d, word)
			return nil
		default:
			return err
	}
}
