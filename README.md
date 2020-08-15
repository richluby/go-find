# GO-FIND

Serves as a simple replacement for `find`, targeted for use on Windows. While
there are packages, powershell, and a host of options, sometimes I just want `find`.

Uses a list of filters under the hood. Cheap filters are executed before expensive filters.
Only the last specified value for a particular filter will be accepted.

# Difference from UNIX `find`

- file paths MUST come after the filters due to the use of the `flag` package  
- fewer filters  
- uses an actual regex instead of name globbing for `-name` and `-iname`

# Example

```
.\go-find.exe -help
Usage of go-find.exe:
  -iname string
        the case-insensitive file name to match
  -mtime string
        the modified file time to match (file time must be after this value)
  -name string
        the file name to match
  -type string
        the file type to match [d,f]

./go-find.exe -name .*.go 
```