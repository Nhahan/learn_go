package dict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errWordExists = errors.New("That word is already exists")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word] // convention way of return map value
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to ths dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	//if err == errNotFound {
	//	d[word] = def
	//} else if err == nil {
	//	return errWordExists
	//}
	return nil
}
