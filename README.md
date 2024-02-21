
### Proto
create folder `go_proto`

1. Create proto:
- product.proto
- pagination.proto

2. Compile
Pastikan sudah menginstall proto compiler di masing-masing
kemudian jalankan
`$ protoc *.proto --go_out=.`

3. Inisialisasi project go
```
- $ go mod init go-proto
- $ go mod tidy
```

4. Create main.go
jalankan dengan perintah 
`$ go run main.go`




