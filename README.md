# todo_app_by_go

## アプリについて
- goで作るtodoアプリです
## 参考
- https://qiita.com/hanadaUG/items/0fd5ddc0407c4ab55d0f
- https://qiita.com/hyo_07/items/59c093dda143325b1859
- dbがmysqlなやつ
https://blog.kannart.co.jp/programming/2026/

## 作業記録
- goのインストール
```
brew install go
go version
```
- ginのインストール
```
go get github.com/gin-gonic/gin
```
- ブラウザにhello worldを表示させる
コミットにて
- DBの接続
各種関数を記述する
- 実行
```
$ go run main.go
main.go:8:5: cannot find package "github.com/jinzhu/gorm" in any of:
main.go:10:5: cannot find package "github.com/mattn/go-sqlite3" in any of:
```
パッケージが見つからない模様。。
ginを入れた時みたいにやらなきゃダメ？
```
go get github.com/jinzhu/gorm
go get github.com/jinzhu/gorm
```
できた！
- 各種ソースリード
コメントを入れたものをコミットに入れる
