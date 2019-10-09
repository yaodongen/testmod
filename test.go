package testmod

import "fmt"
import "util"

func Hi(name string) string {
    util.Show()
    return fmt.Sprintf("Hi, %s! v1.1.1", name)
}
