# cghttp
封装go的http模块成C接口

## windows动态库
```
go build -buildmode=c-shared -o http.dll
dlltool -dllname http.dll --def http.def --output-lib http.lib
```

## ~~windows静态库(无法使用)~~
```
go build -buildmode=c-archive -o http.lib
```
