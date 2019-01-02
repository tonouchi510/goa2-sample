# goa2-sample

### 使用技術
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
- ```$ make goagen```でコード自動生成
- 新しく作られたコントローラの雛形を元にビジネスロジックを追加
- ```$ make run```でビルド&実行

## 関連サーバ起動
mysql、swagger-ui、vironはそれぞれ別個のdockerコンテナで起動する

```
$ make docker-build
$ make docker-up
```

コンテナの停止
```
$ make docker-rm
```

### swagger-ui
<img src="https://github.com/tonouchi510/goa2-sample/blob/master/readme/swagger-ui.png" >

### viron
<img src="https://github.com/tonouchi510/goa2-sample/blob/master/readme/viron_dashboard.png" >

<img src="https://github.com/tonouchi510/goa2-sample/blob/master/readme/viron_adminscreen.png" >
