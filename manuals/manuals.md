# GoLang/Echoのマニュアル

## プロジェクトの生成

```bash
# ディレクトリの作成
mkdir backend
cd backend/

# モジュールの初期化
go mod init github.com/kojikawazu/backend
```

## モジュールの個別インストール

```bash
go get github.com/labstack/echo/v4
go get github.com/go-resty/resty/v2
go get github.com/joho/godotenv
go get github.com/google/uuid
go get github.com/stretchr/testify/mock
```

## モジュールの整理

```bash
go mod tidy
```

## モジュールの実行

```bash
go run main.go
```

## モジュールのテスト

```bash
go test -count=1 ./...
```
