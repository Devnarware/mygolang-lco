package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	msg := [3] string {
		0: primary,
		1: secondary,
		2: tertiary,
	}
	cost := [3] int {
		0: len(primary),
		1: len(secondary) + len(primary),
		2: len(tertiary) + len(secondary) + len(primary),
	}

	return msg, cost
}
