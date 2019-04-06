# goa2-sample

以下の記事用のリポジトリ。

### features
- goa v2
- mysql
- swagger-ui
- viron

## 環境構築
golangおよびdockerは既にインストール済みであることを想定

リポジトリのクローン

```
$ git clone git@github.com:tonouchi510/goa2-sample.git
```

ライブラリのインストール
```
$ make install
```

## 開発
- designディレクトリ以下に新しいAPIデザインを追加
- `$ make goagen` `make regen`でコード自動生成
- 新しく作られたコントローラの雛形を元にビジネスロジックを追加
- `$ make run`でビルド&実行

自動生成コードの移動や修正を行なっているので、Makefileやscriptを参照してください。

## 関連サーバ起動
mysql、vironはdockerコンテナで起動する

```
$ make docker-build
$ make docker-up
```

コンテナの停止
```
$ make docker-rm
```

### swagger-ui
<img src="https://github.com/tonouchi510/goa2-sample/blob/master/readme/swagger-ui-1.png" >

<img src="https://github.com/tonouchi510/goa2-sample/blob/master/readme/swagger-ui-2.png" >

### viron
<img src="https://github.com/tonouchi510/goa2-sample/blob/master/readme/viron_dashboard.png" >

<img src="https://github.com/tonouchi510/goa2-sample/blob/master/readme/viron_adminscreen.png" >
