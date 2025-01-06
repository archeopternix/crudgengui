# pkg

The `pkg` package provides various utility functions and custom error types for handling file operations, string manipulations, and YAML conversions.

## Functions

### IsFirstLetterUppercase
Checks if the first letter of the input string is uppercase.
```go
func IsFirstLetterUppercase(str string) bool 
```

### CheckMkdir
Checks and creates a directory with the given path if it does not yet exist. Throws a DirectoryExistError if the directory already exists.

```Go
func CheckMkdir(path string) error 
```

### StringYAML
Returns a YAML string of the data structure obj or an error if something went wrong.

```Go
func StringYAML(obj interface{}) (string, error) 
```

### CleanString
Converts or removes all non-numeric and non-alphanumeric characters from the input string.

```Go
func CleanString(input string) string
```

### CleanID
Converts any name into a clean ID with just Unicode letters and ensures the first letter is uppercase.

```Go
func CleanID(name string) string {
```

## Custom Error Types
### DirectoryExistError
Thrown by CheckMkdir when a directory already exists.

```Go
type DirectoryExistError struct {
    Dir string;
    Err error;
}

func (r *DirectoryExistError) Error() string 
```

### FileExistError
Thrown when a file already exists and is overwritten by CopyFile.

```Go
type FileExistError struct {
    File string;
    Err  error;
}

func (r *FileExistError) Error() string
```

## Example Usage

```Go
package main

import (
    "fmt"
    "crudgengui/pkg"
)

func main() {
    fmt.Println(pkg.IsFirstLetterUppercase("Hello")); // true
    fmt.Println(pkg.IsFirstLetterUppercase("hello")); // false

    err := pkg.CheckMkdir("./testdir");
    if err != nil {
        fmt.Println(err);
    }

    err = pkg.CopyFile("source.txt", "dest.txt");
    if err != nil {
        fmt.Println(err);
    }

    yamlStr, err := pkg.StringYAML(map[string]string{"key": "value"});
    if err != nil {
        fmt.Println(err);
    }
    fmt.Println(yamlStr);

    fmt.Println(pkg.CleanString("Héllo Wörld!")); // HelloWorld
    fmt.Println(pkg.CleanID("héllo wörld")); // HelloWorld
}
```

## License
This project is licensed under the MIT License.
