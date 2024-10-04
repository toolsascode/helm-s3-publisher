package helpers

func MergeArgs(newArgs []string, args ...string) []string {

	var listArgs []string

	listArgs = append(listArgs, newArgs...)
	listArgs = append(listArgs, args...)
	listArgs = DeleteEmpty(listArgs)

	return listArgs
}

func DeleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
