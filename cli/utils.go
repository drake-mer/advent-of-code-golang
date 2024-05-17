package cli

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

var SESSION string = os.Getenv("SESSION")

func ValidDirectory(directory string) bool {
	_, err := os.Open(directory)
	if err != nil {
		fmt.Printf("Folder %s is not readable: %s\n", directory, err)
		return false
	}
	return true
}

func GetOrCreateCacheDir(year int) string {
	cache_folder, err := os.UserCacheDir()
	if err != nil {
		fmt.Printf("No cache folder for this user: %s\n", err)
	}
	advent_of_code_cache_folder := path.Join(cache_folder, "advent-of-code", fmt.Sprintf("%d", year))
	if !ValidDirectory(advent_of_code_cache_folder) {
		os.MkdirAll(advent_of_code_cache_folder, 0700)
	}
	return advent_of_code_cache_folder
}

func GetFileContent(filename string) (data []string, err error) {
	total_data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Read %d bytes\n", len(total_data))
	file_content := string(total_data[:])
	return strings.Split(file_content, "\n"), nil
}

func SaveFileContent(filepath string, filecontent []string) (err error) {
	opened_file, err := os.Create(filepath)
	if err != nil {
		fmt.Printf("error %s", err)
		return fmt.Errorf("issue to cache input data for aoc %s", err)
	}
	total_lines := len(filecontent)
	for pos, line := range filecontent {
		opened_file.WriteString(line)
		if pos < (total_lines - 1) {
			opened_file.WriteString("\n")
		}
	}
	opened_file.Close()
	return nil
}

func GetInputWithCaching(day int, year int) (data []string) {
	filename := path.Join(GetOrCreateCacheDir(year), fmt.Sprintf("input_day%02d.txt", day))
	file_content, err := GetFileContent(filename)
	if err != nil {
		fmt.Printf("could not fetch content from filesystem: %s\n", err)
		input_strings := DownloadInput(day, year)
		fmt.Printf("saving the content on %s\n", filename)
		err := SaveFileContent(filename, input_strings)
		if err != nil {
			panic("could not save file")
		}
		return input_strings
	} else {
		fmt.Printf("Content loaded from the file system: %s\n", filename)
		return file_content
	}
}

func DownloadInput(day int, year int) []string {
	fmt.Printf("==> Downloading Input <==\n")
	download_url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	request_headers := make(http.Header)
	request_headers.Add("session", SESSION)
	request, err := http.NewRequest(http.MethodGet, download_url, nil)
	if err != nil {
		fmt.Printf("client: error when creating request object")
		os.Exit(1)
	}
	session_cookie := &http.Cookie{
		Name:    "session",
		Value:   SESSION,
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
	content := string(resBody[:])
	fmt.Printf("==> Start of Puzzle Input <==\n%s%s", content, "==>  End of Puzzle Input <==\n")
	return strings.Split(content, "\n")
}

func SubmitAnswer(day int, year int) bool {
	success := false
	return success
}
