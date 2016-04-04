package argconf
import (
    "flag"
    "strconv"
    "strings"
    "os"
    "bufio"
    "errors"
)
func HandleParams(params map[string][]string)(map[string]interface{}, error){
    args := map[string]interface{}{}
    args["config"] = flag.String("config", "", "Config file to read from, may be used as an alternative to writing cli over and over")
    for k, v := range params {
        if(len(v) != 3){
             return nil, errors.New("Not enough arguments for flag: "+k)
        }
        types := v[0]
        defaultvalue := v[1]
        text := v[2]
        
        if types == "string" || types == "str" {
            args[k] = flag.String(k, defaultvalue, text)  
        }else if types == "int" || types == "integer" {
            if i, err := strconv.Atoi(defaultvalue); err == nil {
                args[k] = flag.Int(k, i, text)
                //fmt.Println(reflect.TypeOf(meme))  
            }else{
                return nil, errors.New("Invalid default value for flag: "+k + "| "+err.Error())
            }
        }else if types == "bool" || types == "boolean" {
            if strings.ToLower(defaultvalue) == "true" {
                args[k] = flag.Bool(k, true, text) 
            }else if strings.ToLower(defaultvalue) == "false"{
                args[k] = flag.Bool(k, false, text)
            }else{
                return nil, errors.New("Invalid default value for flag: "+k)
            }
            
        }           
        
    }
    flag.Parse()
    config := make(map[string]string)
    if k, v := args["config"].(*string); v {
        if *k != ""{
            file, err := os.Open(*k)
            if err != nil {
                return nil, errors.New(err.Error())
            }
            defer file.Close()
            scanner := bufio.NewScanner(file)
            for scanner.Scan() {
                ex := strings.Split(scanner.Text(), "=")
                for _, v := range ex[1:] {
                    config[ex[0]] = config[ex[0]]+v
                }
            }
            if err := scanner.Err(); err != nil {
                return nil, errors.New(err.Error())
            }
        }
    }
    flags := map[string]interface{}{}
    for k, v := range config {
        if strings.ToLower(v) == "true" {
            flags[k] = true
        }else if strings.ToLower(v) == "false"{
            flags[k] = false
        }else if i, err := strconv.Atoi(v); err == nil {
            flags[k] = i
        }else{
            flags[k] = v
        }
            
    }
    for k, v := range params {
        if _, vc := flags[k]; vc  {
            if str, ok := args[k].(*string); ok {
                if v[1] != *str {
                    flags[k] = *str
                }
            }else if str, ok := args[k].(*int); ok {
                incInterface := stringToValidType(v[1])
                if incInterface.(int) != *str {
                    flags[k] = *str
                }
            }else if str, ok := args[k].(*bool); ok {
                incInterface := stringToValidType(v[1])
                if incInterface.(bool) != *str {
                    flags[k] = *str
                }
            }
        }else{
            if str, ok := args[k].(*string); ok {
                flags[k] = *str
            }else if str, ok := args[k].(*int); ok {
                flags[k] = *str
            }else if str, ok := args[k].(*bool); ok {
                flags[k] = *str
            }
        }
    }
    return flags, nil
}
func stringToValidType(str string)(interface{}){
    if strings.ToLower(str) == "true" {
        return true
    }else if strings.ToLower(str) == "false"{
        return false
    }else if i, err := strconv.Atoi(str); err == nil {
        return i
    }else{
        return str
    }
}