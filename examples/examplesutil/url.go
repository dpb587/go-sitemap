package examplesutil

func URLFromArgOrDefault(args []string, def string) string {
	if len(args) > 1 {
		return args[1]
	}

	return def
}
