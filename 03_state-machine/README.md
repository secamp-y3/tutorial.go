# 事前課題 3: State machine

## 実行方法

### ローカルで実行する場合
以下の手順で，サーバとクライアントをそれぞれ別のターミナルウィンドウ上で起動する．

1. サーバの起動
    ```sh
    go run cmd/server/main.go --port 8080
    ```
2. クライアントの起動
    ```sh
    go run cmd/client/main.go --server localhost:8080
    ```

### Dockerを使用する場合
0. 初回のみコンテナをビルドする．
    ```sh
    docker-compose build
    ```
1. コンテナを起動する．
    ```sh
    docker-compose up -d
    ```
2. クライアントを起動する．
    ```sh
    docker-compose exec client go run main.go
    ```

#### 有効なコマンド
- `echo <message>`: メッセージを送信する (事前課題2の機能を保存している)
- `add <val>`: `val` を加算する
- `sub <val>`: `val` を減算する
- `mul <val>`: `val` を乗算する
- `div <val>`: `val` を除算する (`val = 0`の場合エラーが返ってくる)

#### 補足
- Dockerを使用する場合，サーバはバックグラウンドで動作しているため明示的に起動する必要はない．
- クライアントは環境変数からサーバアドレスを読み込むため，`--server`オプションによる指定は必要ない．
- Docker内部では，以下のIPアドレスを割り当てている．
    | App    | IP            |
    |--------|---------------|
    | server | 172.26.249.11 |
    | client | 172.26.249.21 |
- 使用するネットワークが同じであるため，事前課題1および事前課題2のコンテナをシャットダウンしておく必要があります．
