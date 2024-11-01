package pkg

import "fmt"

const (
	alph = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func Find(c uint8) (int, error) {
	for i := range alph {
		if c == alph[i] {
			return i, nil
		}
	}
	return -1, fmt.Errorf("symbol not found")
}

func Increment(oldString string) (string, error) {
	newString := []byte(oldString)
	i := len(oldString) - 1
	n := len(alph)
	for i > -1 {
		pos, err := Find(newString[i])
		if err != nil {
			return "", err
		}
		if pos != n-1 {
			newString[i] = alph[pos+1]
			i--
			break
		} else {
			for i > 0 && pos == n-1 {
				newString[i] = alph[0]
				i--
				pos, err = Find(newString[i])
				if err != nil {
					return "", err
				}
			}
			if pos == n-1 {
				return "", fmt.Errorf("Overflow")
			}
			newString[i] = alph[pos+1]
			break
		}
	}
	return string(newString), nil
}
