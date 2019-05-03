# Zipaddr

Zipaddr はシングルバイナリで動作する郵便番号検索サーバーです。

## 注意事項

- このプログラムは [日本郵便株式会社が提供する郵便番号データ](https://www.post.japanpost.jp/zipcode/dl/kogaki-zip.html) を利用しています。

## インストール方法

```sh
$ go get github.com/okeyaki/zipaddr
```

## 使用方法

まず最初に、郵便番号データを更新しておく必要があります。

```sh
$ zipaddr update --data-dir /var/lib/zipaddr 
```

郵便番号データの更新が完了したら、サーバーを起動してください。

```sh
$ zipaddr serve --data-dir /var/lib/zipaddr &
```

サーバーが起動したら、郵便番号検索 API が利用できるようになります。

```sh
$ curl http://localhost:1234/addrs/0600010
```

## コマンド

### 郵便番号データを更新する

`update` コマンドは郵便番号データを更新します。

```sh
$ zipaddr update --data-dir /var/lib/zipaddr
```

サーバーが起動している（`serve` を実行している）場合は、サーバーの再起動が必要です。

```sh
$ pkill zipaddr

$ zipaddr serve --data-dir /var/lib/zipaddr
```

### サーバーを起動する

```sh
$ zipaddr serve --data-dir /var/lib/zipaddr
```

特に指定のない場合、サーバーは `0.0.0.0:3000` で起動します。

起動するアドレスを変更したい場合は、`--listen` オプションを指定してください。

```sh
$ zipaddr serve --data-dir /var/lib/zipaddr --listen 0.0.0.0:1234
```

## API

### 郵便番号を検索する

```sh
$ curl http://localhost:1234/addrs/0600010
```
