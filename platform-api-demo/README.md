##GO Open Platform API Demo

####Getting started
`go run platform_api_demo.go`

####Settings
configure these variables value in the go file will be executed;
like your appKey/appSecret/code...
```
func GetOpenId() {
    appKey := ""
    appSecret := ""
    code := ""
}
```

####Examples
#####GetOpenId
```
func main() {
    market.GetOpenId()
}
```