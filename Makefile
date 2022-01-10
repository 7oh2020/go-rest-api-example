APPNAME="go-rest-api-example"
OUTDIR="./dist"
DSN="mysql://$(MYSQL_DSN)"
MIGRATION_DIR="file://db/migration"

# テスト + ビルドを行います。
all: test build

# 一時ファイルの削除などを行います。
clean:
	@rm -rf $(OUTDIR)
	@mkdir $(OUTDIR)

# アプリケーションの実行に必要な依存関係をインストールします。
depend:
	@go mod tidy
	@go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# アプリケーションをテストします。
test: depend 
	@go test ./...

# アプリケーションをビルドします。
build: clean depend test
	@go build -o $(OUTDIR)/$(APPNAME)

# アプリケーションを実行します。
# freshパッケージにより、ファイルが保存される度に自動ビルドされます。
run: depend
	@go get github.com/pilu/fresh
	fresh
	
# データベースのバージョンを１つ進めます。
up: depend
	migrate -database=$(DSN) -source=$(MIGRATION_DIR) up

# データベースのバージョンを１つ戻します。
down: depend
	migrate -database=$(DSN) -source=$(MIGRATION_DIR) down

# コードの自動生成を行います。
# mockery を使用すると、インタフェースからtestifyパッケージのモックを自動生成できます。
generate:
	@go get github.com/vektra/mockery/v2/.../
	mockery --all --recursive=true
	@go mod tidy
