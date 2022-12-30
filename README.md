// Catch to mockgen error 
error: "mockgen missing go.sum entry for module providing package github .com/golang/mock/mockgen/model"
```
mockgen -build_flags=--mod=mod -package mockdb -destination db/mock/store.go github.com/alrobwilloliver/bank/db/sqlc Store
```