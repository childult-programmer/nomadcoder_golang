package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errWordExists = errors.New("That words already exist in the dictionary")
	errCantUpdate = errors.New("Cant update non-existing word")
	errCantDelete = errors.New("The word that you want to delete does not exist in the dictionary")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the Dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	// if the word is not in dictionary, add the word to the dictionary
	if err == errNotFound {
		d[word] = def
	} else if err == nil {
		return errWordExists
	}
	return nil
}

// Update the definition of the word
func (d Dictionary) Update(word, def string) error {
	// The word that you want to update must be in the dictionary
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete word in dictionary
func (d Dictionary) Delete(word string) error {
	// The word that you want to delete must be in the dictionary
	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word)
	case errNotFound:
		return errCantDelete
	}
	return nil
}
