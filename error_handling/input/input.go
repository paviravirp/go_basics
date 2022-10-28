package input

import "fmt"
import "errors"

func Input(name string) (string, error) {

	if name == "" {
		return "", errors.New("name cannot be empty")
	}

	message := fmt.Sprintf("Hello %v", name);
	return message, nil;
}