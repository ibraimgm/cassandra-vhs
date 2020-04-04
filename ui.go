package main

import (
	"fmt"
	"sort"
)

func ask(prompt string) string {
	fmt.Printf("%s: ", prompt)

	var value string
	fmt.Scanln(&value)

	return value
}

func askIn(prompt string, values ...string) (string, bool) {
	response := ask(prompt)

	for _, str := range values {
		if response == str {
			return response, true
		}
	}

	return response, false
}

func menu(items map[string]string) string {
	keys := make([]string, 0, len(items))

	for k := range items {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for {
		fmt.Printf("Available options:\n\n")

		for _, key := range keys {
			fmt.Printf("%3s - %s\n", key, items[key])
		}
		fmt.Printf("\n")

		if resp, ok := askIn("Choose an option", keys...); ok {
			return resp
		}
		fmt.Printf("=== ERROR: invalid option! ===\n\n")
	}
}
