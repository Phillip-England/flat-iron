package lib

import "strings"

func IsStringInSlice(target string, slice []string) bool {
    for _, value := range slice {
        if strings.EqualFold(target, value) {
            return true
        }
    }
    return false
}
