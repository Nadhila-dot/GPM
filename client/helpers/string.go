package nadhi

import "strings"


func CheckString(shouldInclude []string, fullString string) bool {
    for _, s := range shouldInclude {
        if !strings.Contains(fullString, s) {
            return false
        }
    }
    return true
}