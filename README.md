# My Security Tool
Security Tool with various options like fuzzing

## Fuzzing
Fuzzing is a tool which injects random data into different protocol stacks. Fuzzing is an important technique used in security testing. More information [OWASP Fuzzing](https://owasp.org/www-community/Fuzzing). Here in "My Security Tool" we read a file which contains fuzzing data. This data is read line by line and is appended to the url provided. With the new appended url we make an http call and check if the http call is a success.

# Usage 
1. Build
   ```
   go build
   ```
2. Help
   ```
   ./mysecuritytool --help 
   My Security Tool contains mutiple security tools like fuzzing

   Usage:
    mysecuritytool [command]

   Available Commands:
     completion  Generate the autocompletion script for the specified shell
     fuzzing     Fuzzing
     help        Help about any command

   Flags:
     -h, --help   help for mysecuritytool

   Use "mysecuritytool [command] --help" for more information about a command.
   ```
3. Fuzzing help
   ```
   ./mysecuritytool fuzzing --help
   Fuzzing tool where the file line text will be appended/replaced. 
   You can use {{.fuzz}} as a template where you can replace the passwords. You can use {{.fuzz}} in url, body or headers of the payload.

   Example:
   ./mysecuritytool fuzzing --file fuzz.txt --url https://your_url?name={{.fuzz}} --header "Cookie: auth_token=test_token"

   Usage:
   mysecuritytool fuzzing [flags]

   Flags:
         --body string          Body for post
         --file string          File name used for fuzzing
         --header stringArray   Headers. Pass mutiple headers like --header 'Cookie:auth_token=test' --header  'Content-Type:application/x-www-form-urlencoded'
   -h, --help                 help for fuzzing
         --regex string         Regex in fuzzing file line
         --url string           Url to make call
   ```
4. Examples of fuzzing command
   ```
   ./mysecuritytool fuzzing --file fuzz.txt --url https://URL --header "Cookie: auth_token=test_token" --header "Content-Type: application/x-www-form-urlencoded" --regex ".inc"

   ./mysecuritytool fuzzing --file fuzz.txt --url https://your_url?name={{.fuzz}} --header "Cookie: auth_token=test_token"

   ./mysecuritytool fuzzing --file fuzz.txt --url https://your_url --body "{\"name\":\"{{.fuzz}}\"}" --header "Cookie: auth_token=test_token"
   ```

### Fuzzing file references 
1. https://github.com/danielmiessler/SecLists/
2. https://github.com/Bo0oM/fuzz.txt/blob/master/fuzz.txt
3. https://github.com/praetorian-inc/Hob0Rules/blob/master/wordlists/rockyou.txt.gz
