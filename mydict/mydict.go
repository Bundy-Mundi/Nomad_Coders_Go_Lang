package mydict

import (
	"errors"
)

// Dictionary type
type Dictionary map[string]string

// Error types Here
var (
	errDoesNotExist = errors.New("Word Does Not Exist")
	errAlreadyExist = errors.New("The Word Already Exist")
	errCannotUpdate = errors.New("The Word Can't be Updated. Try to Add it First")
	errCannotDelete = errors.New("The Word Can't be Deleted. Try to Add it First")

)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if !exists {
		return "", errDoesNotExist
	}
	return value, errAlreadyExist
}

// Add a Word
func (d Dictionary) Add(word string, def string) error {
	_, err := d.Search(word)
	switch err {
	case errDoesNotExist:
		d[word] = def
	case errAlreadyExist:
		return errAlreadyExist
	}
	return nil
}

// Update a Word
func (d Dictionary) Update(word string, def string) error{
	_, err := d.Search(word)
	switch err {
	case errAlreadyExist:
		d[word] = def
	case errDoesNotExist:
		return errCannotUpdate
	}
	return nil
}

// Delete a Word
func (d Dictionary) Delete(word string) (string, error) {
	_, exists := d[word]
	if exists{
		delete(d, word)
	}else{
		return "", errCannotDelete
	}
	return word, nil
}

// MyAdd words 
func (d Dictionary) MyAdd(word string, def string) error{
	_, exists := d[word]
	if !exists {
		d[word] = def
	}else{
		return errAlreadyExist
	}
	return nil
}

// MyUpdate a word
func (d Dictionary) MyUpdate(word string, newDef string) (string, error) {
	v, err := d.Search(word)
	if err == errAlreadyExist{
		d[word] = newDef
	}else{
		return v, errDoesNotExist
	}
	return v, nil
}

func (d Dictionary) String() map[string]string{
	return d
}
