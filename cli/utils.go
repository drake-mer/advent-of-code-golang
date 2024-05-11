package cli

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const SESSION = "53616c7465645f5faf675d8ae7b16e7a9bc98d2bc0eeb97a551918d7b8e6b1e61425e46d423bbe21bc27e9c7ace54ed52b14f45c298fb8e7bbc92e866df36405"


func DownloadInput (day int, year int) []string {
	download_url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	request_headers := make(http.Header)
	request_headers.Add("session", SESSION)
	request, err := http.NewRequest(http.MethodGet, download_url, nil)
	if err != nil {
		fmt.Printf("client: error when creating request object")
		os.Exit(1)
	}
	session_cookie := &http.Cookie{
		Name: "session",
		Value: SESSION,
		Expires: time.Now().Add(24 * time.Hour),
	}
	request.AddCookie(session_cookie)

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("client: error making http request:%s\n", err)
		os.Exit(1)
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
	}
	fmt.Printf("Downloaded %s\n", resBody)
	content := string(resBody[:])
	return strings.Split(content, "\n")
}