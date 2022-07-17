# go-helper

## 概要

Go 周りでいろいろ面倒なことをサポートするやつ（CLI）

## 機能一覧

現時点では下記のみ

### Go プロジェクトの生成

- ディレクトリ名を指定して Go プロジェクトを作成する。
  - デフォルトでは必要最低限の記述が行われた main ファイルを生成する
  - yml ファイルにコンフィグを記述し、追加で go.mod の作成とかも可能
- Go を使ってクイックで何かを検証したい時とかに使うやつ

## 使い方

### install して使う（推奨）

```
❯ make install
```

### build して使う

```
// リポジトリトップで下記を実行することで <リポジトリ>/bin 配下に go-helper バイナリが作成される
❯ make
```

### プロジェクトの作成

- プレーンなプロジェクトの作成

  ```
  ■ プロジェクト作成
  ❯ gohelper init -n "myproject"
  Start creating a project.
  Project Name: myproject
  Finished.

  ■ 確認
  ❯ ls myproject
  myproject.go
  ```

- 詳細設定を行う

  ```
  ■ yml ファイルにて設定を定義（下記は例）
  ❯ cat gohelper_tmpl.yml
  gomod:
    repository: github.com/tken # go.mod で指定するリポジトリ名
    goVersion: 1.18             # go.mod で指定する Go のバージョン
  withCmd: true                 # main パッケージを <リポジトリ>/cmd 配下に作成する
  httpServer: false             # HTTP サーバの最小構成で Go のファイルを作成する


  ■ yml 定義を指定してプロジェクトを作成
  ❯ gohelper init -c "gohelper_tmpl.yml" -n "myproject"
  Start creating a project.
  Project Name: myproject
  Read configration: gohelper_tmpl.yml
  [config] repository: github.com/tken
  [config] go version: 1.18
  [config] cmd dir: true
  [config] http server: false
  Finished.

  ■ 確認
  ❯ tree myproject
  myproject
  ├── cmd
  │   └── myproject
  │       └── myproject.go
  └── go.mod

  1 directory, 2 files
  ```

### yml 定義のテンプレートを取得

```
■ 生成
❯ gohelper init make-template

■ ファイル名 `gohelper_tmpl.yml` で作成される
❯ cat gohelper_tmpl.yml
gomod:
  name: github.com/<username>
  goVersion: 1.18
withCmd: false
httpServer: false
```

- （おまけ）出力先を指定して作成

  ```
  ❯ ls
  templates

  ❯ gohelper init make-template --dir "templates"

  ❯ cat templates/gohelper_tmpl.yml
  gomod:
    name: github.com/<username>
    goVersion: 1.18
  withCmd: false
  httpServer: false
  ```
