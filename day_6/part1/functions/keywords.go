package functions

func KeyWords(str string) string {
	s := ""
	keywords := []string{"toggle", "off", "on"}
	tasks := make(map[string]func())

	// Define tasks for each keyword
	tasks["toggle"] = func() {
		s = "Toggling the switch"
	}

	tasks["off"] = func() {
		s = "Turning the device off"
	}

	tasks["on"] = func() {
		s = "Turning the device on"
	}

	// Range through the str string
	for i := 0; i < len(str); i++ {
		for _, keyword := range keywords {
			if i+len(keyword) <= len(str) && str[i:i+len(keyword)] == keyword {
				tasks[keyword]()
				i += len(keyword) - 1
				break
			}
		}
	}
	return s
}
