# sample api

```sh
go run main.go
```

```sh
$ curl localhost:8080/echojson -X POST -d '{"id": 100, "message":"test"}'
{"id":100,"message":"test"}

$ curl http://localhost:8080/random2 -s | jq .
{
  "results": [
    {
      "success": 95
    },
    {
      "success": 66
    },
    {
      "fail": 28
    },
    {
      "success": 58
    },
    {
      "fail": 47
    },
    {
      "fail": 47
    },
    {
      "success": 87
    },
    {
      "success": 88
    },
    {
      "success": 90
    },
    {
      "fail": 15
    }
  ]
}
```
