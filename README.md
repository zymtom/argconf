# argconf
Golang package for utilizing a config file aswell as being able to input args
To download you simply do
```
go get github.com/zymtom/argconf
```

Usage:

```
paramMap := map[string][]string{
    "file":[]string{"string", "", "File you want to be printed"},
    "pause":[]string{"bool", "false", "Pause before exiting"},
    "lines":[]string{"int", "2", "Number of lines to print"},
}
values, err := argconf.HandleParams(paramMap)
```

This will let you provide your variables both through config file and through arguments, as demonstrated in the demo. Config file needs to be provided via commandline though.

i.e
```
./package -config=file.conf
```
which will then read the config file in a format like this
```
pause=true
lines=2
```