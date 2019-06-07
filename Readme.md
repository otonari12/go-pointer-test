## 概要
関数への値渡し+ポインタ渡しのベンチマークをとる

## ベンチマークを取る
go test -bench . -benchmem ./...

## メモリ割り当て先のチェック
go build -gcflags -m ./...

