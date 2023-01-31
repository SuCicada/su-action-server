```bash
go install -u github.com/jteeuwen/go-bindata/...

go-bindata -o template/notification.go -prefix ./template/ -pkg template  ./template/
```
