package utilities

// CheckStringInSlice checks if string s is in given array
func CheckStringInSlice(s string, arr []string) bool{
	for _, i := range arr{
		if i == s{
			return true
		}
	}
	return false
}
