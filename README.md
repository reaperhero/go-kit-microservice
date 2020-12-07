# go kit microservice
```
curl http://localhost:8080/get
{"date":"07/12/2020"}

curl http://localhost:8080/status
{"status":"ok"}

curl -XPOST -d '{"date":"32/12/2020"}' http://localhost:8080/validate
{"valid":false,"err":"parsing time \"32/12/2020\": day out of range"}

curl -XPOST -d '{"date":"12/12/2021"}' http://localhost:8080/validate
{"valid":true}
```


