package textsearch

import (
	"errors"
	"strings"
)

type TextList struct {
	T []string
}

// add a value to the slice
func (t *TextList) AddValue(value string) error {

	value = strings.ToLower(value)

	for _, val := range t.T {
		if val == value {
			return errors.New("value already exist")
		}
	}

	t.T = append(t.T, value)
	return nil
}

// remove a value from the slice
func (t *TextList) RemoveValue(value string) error {

	value = strings.ToLower(value)

	for i, val := range t.T {
		if val == value {
			t.T[i] = t.T[len(t.T)-1]
			t.T[len(t.T)-1] = ""
			t.T = t.T[:len(t.T)-1]
			return nil
		}
	}

	return errors.New("value not found")
}

// search a value in the slice
func (t *TextList) SearchValue(value string) ([]string, error) {

	value = strings.ToLower(value)

	var res []string

	for _, val := range t.T {
		if strings.Contains(val, value) {
			res = append(res, val)
			if len(res) >= 5 {
				return res, nil
			}
		}
	}

	if len(res) > 0 {
		return res, nil
	} else {
		return nil, errors.New("value not found")
	}

}
