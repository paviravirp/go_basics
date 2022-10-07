package input

import "fmt"

func Input(name string) string {
	message := fmt.Sprintf("Hello %v", name);
	return message;
}