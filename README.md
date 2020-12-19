# GO build flags for Windows and Mac

# For Windows: 
  ### Building a binary for linux from windows
  - $Env:GOOS = "linux"; $Env:GOARCH = "amd64"; go build -o output_filename

# For Mac:
  ### Building a binary for windows from mac
  - GOOS=windows GOARCH=386 go build -o output_filename

If you want to see what go compiler actually does when building a go binary we can use the -x flag with build command

```
go build -x main.go
```

If you want to pass flags as variable when building the go binary we can use -ldflags

```
go build -ldflags="-X main.version=v0.0.0" main.go
```

Here, we are passing value to the variable version that is present in our project file from the command line on build

+ -w removes dwarf information
+ -s removes debug information
+ -o helps assign an output filename 
  
We can use multiple flags at build time at once by combining them in -ldflags="<build flags and their values>"

By using -w and -s flags we can reduce the size of the compiled go binary since the standard binary comes with some additional info 

Here, we have two files 
+ goflags 
+ goflags_reduced

On linux or mac, we can use the following command to see the size of the file

```
du -h <file_name> 
```

### The reduced size has no impact on the funtionality of the binary 

There are a lot more commands and build flags available for go

To see all the tools available for compiling go code use the cmd below

```
go tool compile --help
```
