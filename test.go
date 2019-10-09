package testmod

import "fmt"

func Hi(name string, lang string) string {
    return fmt.Sprintf("Hi, %s, %s", name, lang)
}
