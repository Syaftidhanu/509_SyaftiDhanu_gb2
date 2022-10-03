package main

func main() {
	rs := []interface{}{1, "Arell", true, 2, "ananda", true}

	rm := map[string]interface{}{
		"Name":   "Arell",
		"Status": true,
		"Age":    23,
	}

	_, _ = rs, rm
}
