package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func checkUsernameAvailability(username string) (bool, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", username)
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return true, nil // Username is available
	} else if resp.StatusCode == http.StatusOK {
		return false, nil // Username exists
	}
	return false, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
}

func main() {
	// Specify the usernames.txt file directly
	fileName := "usernames.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		username := scanner.Text()
		if username == "" {
			continue
		}
		available, err := checkUsernameAvailability(username)
		if err != nil {
			fmt.Println("Error checking", username, ":", err)
			continue
		}

		if available {
			fmt.Printf("The username '%s' is available!\n", username)
		} else {
			fmt.Printf("The username '%s' is taken.\n", username)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
