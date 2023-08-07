package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"github.com/sumeetpatil/mysecuritytool/internal/httpclient"
)

type Fuzzing struct {
	Fuzz string
}

var fuzzingCmd = &cobra.Command{
	Use:   "fuzzing",
	Short: "Fuzzing",
	Long: `Fuzzing tool where the file line text will be appended/replaced. 
You can use {{.fuzz}} as a template where you can replace the passwords. You can use {{.fuzz}} in url, body, cookie or headers of the payload.

Example:
./mysecuritytool fuzzing --file fuzz.txt --url https://your_url?name={{.fuzz}} --cookie "auth_token=test_token"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fuzzing called")

		filePath, err := cmd.Flags().GetString("file")
		if err != nil {
			log.Fatalf("error - %s", err.Error())
		}

		cookie, err := cmd.Flags().GetString("cookie")
		if err != nil {
			log.Fatalf("error - %s", err.Error())
		}

		headers, err := cmd.Flags().GetString("headers")
		if err != nil {
			log.Fatalf("error - %s", err.Error())
		}

		url, err := cmd.Flags().GetString("url")
		if err != nil {
			log.Fatalf("error - %s", err.Error())
		}

		pattern, err := cmd.Flags().GetString("regex")
		if err != nil {
			log.Fatalf("error - %s", err.Error())
		}

		body, err := cmd.Flags().GetString("body")
		if err != nil {
			log.Fatalf("error - %s", err.Error())
		}

		regex := regexp.MustCompile(pattern)

		file, err := os.Open(filePath)
		if err != nil {
			log.Fatalf("error opening file - %s", err.Error())
		}

		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()

			cookieMap := make(map[string]string)
			if cookie != "" {
				cookie = replaceData(cookie, line)
				cookieData := strings.Split(cookie, ";")
				for _, v := range cookieData {
					cookieDataSplit := strings.Split(v, "=")
					cookieMap[strings.TrimSpace(cookieDataSplit[0])] = strings.TrimSpace(cookieDataSplit[1])
				}
			}

			headersMap := make(map[string]string)
			if headers != "" {
				headers = replaceData(headers, line)
				headersData := strings.Split(cookie, ";")
				for _, v := range headersData {
					headersDataSplit := strings.Split(v, "=")
					headersMap[strings.TrimSpace(headersDataSplit[0])] = strings.TrimSpace(headersDataSplit[1])
				}
			}

			url = replaceData(url, line)
			body = replaceData(body, line)

			if pattern != "" {
				if regex.MatchString(line) {
					call(url, line, cookieMap, headersMap, body)
				}
			} else {
				call(url, line, cookieMap, headersMap, body)
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatalf("error reading file - %s", err.Error())
		}

		log.Fatalf("no success in fuzzing")
	},
}

func replaceData(data string, fileLineText string) string {
	return strings.ReplaceAll(data, "{{.fuzz}}", fileLineText)
}

func call(url string, line string, cookieMap map[string]string, headersMap map[string]string, body string) {
	url = url + line
	fmt.Println(line)
	client := httpclient.NewHttpClient(url, cookieMap, headersMap)
	httpStatus := httpclient.HttpResp{}
	if body != "" {
		httpStatus = client.Post(body)
	} else {
		httpStatus = client.Get()
	}

	if httpStatus.StatusCode == 200 {
		log.Println("Success with " + line)
		os.Exit(0)
	}
}

func init() {
	rootCmd.AddCommand(fuzzingCmd)
	fuzzingCmd.Flags().String("file", "", "File name used for fuzzing")
	fuzzingCmd.Flags().String("url", "", "Url to make call")
	fuzzingCmd.Flags().String("cookie", "", "Cookie information")
	fuzzingCmd.Flags().String("headers", "", "Header information. You can pass like Content-Type=application/json;Accept:application/json")
	fuzzingCmd.Flags().String("regex", "", "Regex in fuzzing file line")
	fuzzingCmd.Flags().String("body", "", "Body for post")
}
