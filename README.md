# GO-FIND

Serves as a simple replacement for `find`, target for use on Windows. While
there are packages, powershell, and a host of options, sometimes I just want `find`.

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
```