package main
import (
    "github.com/zymtom/argconf"
    "log"
    "fmt"
    "os"
    "bufio"
)
func main(){
    paramMap := map[string][]string{
        "file":[]string{"string", "", "File you want to be printed"},
        "pause":[]string{"bool", "false", "Pause before exiting"},
        "lines":[]string{"int", "2", "Number of lines to print"},
    }
    values, err := argconf.HandleParams(paramMap)
    if values["file"].(string) != "" {
        if err != nil {
            log.Fatal(err)
        }
        file, err := os.Open(values["file"].(string))
        if err != nil {
            log.Fatal(err)
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)
        x := 0
        for scanner.Scan() {
            x++
            if x > values["lines"].(int) {
                break
            }
            fmt.Println(scanner.Text())
        }

        if err := scanner.Err(); err != nil {
            log.Fatal(err)
        }
    }
    if values["pause"].(bool) {
        var input string
        fmt.Scanln(&input)
        return
    }
}