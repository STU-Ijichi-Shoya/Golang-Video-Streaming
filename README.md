# Golang-Video-Streaming
Go言語で実装した動画配信サイトです．
アーキテクチャや，バックエンドを学ぶため，０から実装しています．  
[フロントエンド実装はこちらから確認可能です．（サーバはまだ完成してないのでGithubPagesで）](https://stu-idichi-syoya.github.io/Golang-Video-Streaming/)

## バックエンド
まず，動画配信サイトでは，配信する動画をMPEGーDASHやHLSなどの規格で細かく分割し，各ビットレートごとに用意，マニュフェストファイルを配信するといった
感じですが，
このシステムでは，とりあえず，ローカルの規定のディレクトリに動画をぶちこんで，そのURIをそのまま流しています．
いずれ，MPEG-DASHやHLSで配信したいと考えています．
ただ，今回の目的として，バックエンドを体系的に学ぶことなので，先になると思います．
以下にシステムの概要を示します．

### アーキテクチャ
    レイヤードアーキテクチャを採用しています．
    今回の実装ではフレームワークにGinを用いていますが，  
    Echoによる実装やGolangの標準パッケージのnet/httpに簡単に書き換えることが可能です．

### DBMS
    DBMSとして，MySQLを用いています．

## フロントエンド
    純粋なHTML5とJavaScriptを用いています．
    
    [フロントエンド実装はこちらから確認可能です．（サーバはまだ完成してないのでGithubPagesで）](https://stu-idichi-syoya.github.io/Golang-Video-Streaming/)
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
