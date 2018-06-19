// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Call the GitHub API to get a list of repository contributors.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Create a type where we can decode contributor json values.
// It needs the fields "login" and "contributions".

type Contributor struct {
	Login         string
	Contributions int
}

func main() {

	// Get an access token from the environment.
	tkn := os.Getenv("GITHUB_TOKEN")
	if tkn == "" {
		log.Print("Token not found. You must set it in your environment like")
		log.Print("export GITHUB_TOKEN=000a0aaaa0000a00000000aaa00000000a000000")
		log.Print("You can generate a token at https://github.com/settings/tokens")
		os.Exit(1)
	}

	// Create a request for the contributors api endpoint.
	url := "http://127.0.0.1:59528/repos/golang/go/contributors"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Add the access token in the "Authorization" header.
	// The value should be in the form "token 000aa0a0..."
	req.Header.Set("Authorization", "token "+tkn)

	// Create an http.Client and make the request.
	var cl http.Client
	resp, err := cl.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// Defer closing the response body.
	defer resp.Body.Close()

	// Ensure we get a 200 OK status back.
	fmt.Println(resp.Status)

	// Decode the results into a []contributor.
	var cons []Contributor
	if err := json.NewDecoder(resp.Body).Decode(&cons); err != nil {
		log.Fatal(err)
	}

	// Loop through the []contributor and print the values.
	for i, con := range cons {
		fmt.Println(i, con.Login, con.Contributions)
	}
}
