To create swagger: 
```go install github.com/swaggo/swag/cmd/swag@1.16```
```swag init```

To deploy on server:
docker --context tp_remote compose -f docker-compose.prod.yml up -d