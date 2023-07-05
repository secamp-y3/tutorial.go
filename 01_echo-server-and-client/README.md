# 事前課題 1: Echo server and client

## 実行方法

### ローカルで実行する場合
以下の手順で，サーバとクライアントをそれぞれ別のターミナルウィンドウ上で起動する．

1. サーバの起動
    ```sh
    go run cmd/server/main.go --port 8080
    ```
2. クラインとの起動
    ```sh
    go run cmd/client/main.go --server localhost:8080
    ```

### Dockerを使用する場合
0. 初回のみ，コンテナをビルドする．
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
