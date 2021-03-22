# Golang-video-broadcast
Go言語で実装した動画配信サイトです．
アーキテクチャや，バックエンドを学ぶため，０から実装しています．
## バックエンド

###アーキテクチャ
    レイヤードアーキテクチャを採用しています．
    今回の実装ではフレームワークにGinを用いていますが，  
    Echoによる実装やGolangの標準パッケージのnet/httpに簡単に書き換えることが可能です．

### DBMS
    DBMSとして，MySQLを用いています．

## フロントエンド
    bootstrapを用いています．
    純粋なHTML5とJavaScriptを用いています．
## ディレクトリ構成
```bash
.
├── README.md
├── cmd #mainプログラム
│   └── main.go
├── config ## DBの設定など
├── domain ## APIのためのユースケースのインターフェースの実装です．
│   ├── repository ### 各ユースケースに応じたDB操作のインターフェースが入っています．
│   │   └── videoRepository.go
│   ├── specification ### `仕様`ファイルです．リポジトリとユースケースが複雑に絡んだときに使用します．
│   └── video.go
├── infrastructure ## repositoryのインターフェースの実装が入っています．
│   └── video.go
├── interfaces ## handlerが入っています．インターフェースとそれを実装したものが入っています．
│   └── handler
│       ├── video.go
│       └── web.go
└── usecase ## ユースケースのインターフェースです．
    └── video.go

```

## 実行方法
```go
go build cmd/main.go
./main
```
