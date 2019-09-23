package main 

// Dictionary is a type
type Dictionary map[string]string

const (
	// ErrNotFound does things
	ErrNotFound = DictionaryError("could not find the word you were looking for")
	// ErrWordExists does things
	ErrWordExists = DictionaryError("word already exists")
	// ErrWordDoesNotExist does things
	ErrWordDoesNotExist = DictionaryError("cannot update a word that does not exist")
)

// DictionaryError is a type
type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

// Search finds and returns a defintion from the dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add adds a word to the dictionary
func (d Dictionary) Add(word string, defintion string) error {
	if _, found := d[word]; found == true {
		return ErrWordExists
	}
	d[word] = defintion
	return nil
}

// Update changes an existing word in the dictionary to have a new definition
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

// Delete will remove a word from the dictionary
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	if err != nil {
		return err
	}

	delete(d, word)

	return nil
}
