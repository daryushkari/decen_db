package utilities


// CheckStringInSlice checks if string s is in given array
func CheckStringInSlice(s string, arr []string) bool {
	for _, i := range arr {
		if i == s {
			return true
		}
	}
	return false
}

// Remove a value from slice of strings
func RemoveFromSlice(value string, slice []string) []string {
	for i, v := range slice{
		if v == value{
			slice[i] = slice[len(slice) - 1]
			slice = slice[:len(slice) - 1]
		}
	}
	return slice
}