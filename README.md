# My Security Tool
Security Tool with various options like fuzzing

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
   Fuzzing url tool where the file line text will be appended at the end of url.

   Usage:
     mysecuritytool fuzzing [flags]

   Flags:
         --body string     Body for post
         --cookie string   Cookie information
         --file string     File name used for fuzzing
     -h, --help            help for fuzzing
         --regex string    Regex in fuzzing file line
         --url string      Url to make call
   ```
4. Example fuzzing command
   ```
   ./mysecuritytool fuzzing --file fuzz.txt --url https://URL --cookie "auth_token=test_token" --regex ".inc"
   ```

### Fuzzing file references 
1. https://github.com/danielmiessler/SecLists/
2. https://github.com/Bo0oM/fuzz.txt/blob/master/fuzz.txt