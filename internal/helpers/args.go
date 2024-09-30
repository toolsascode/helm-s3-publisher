package helpers

func MergeArgs(newArgs []string, args ...string) []string {

	var listArgs []string

	listArgs = append(listArgs, newArgs...)
	listArgs = append(listArgs, args...)

	return listArgs
}
