# セキュリティ・キャンプ全国大会 開発トラック Y3 事前学習
このレポジトリは，セキュリティ・キャンプ全国大会 開発トラック Y-3「故障を乗り越えて動くシステムのための分散合意」事前学習の内容とサンプルコードをまとめたものです．

すべてのサンプルコードは Go 1.20 で動作することを確認しています．

## 課題 1: Echo server and client
TCP通信を用いて，以下の動作を実現するサーバ・クライアントアプリケーションを作成せよ．

#### Server
1. クライアントからTCP通信で送られてきた文字列を受け取る．
2. 受け取った文字列をそのままクライアントへ返信する (echo back)．

#### Client
1. 標準入力から文字列を受け取る．
2. サーバへTCP通信で入力文字列を送信する．
3. サーバからの応答を標準出力を通じて画面に表示する．


## 課題 2: RPC
`net/rpc`パッケージを使用し，課題1のシステムをRPCで実装し直してみよ．

### 補足
- RPCとは「遠隔手続き呼出し (Remote Procedure Call)」と呼ばれる技術であり，一般に特定の処理などをネットワークを通じて別の端末上で実行することを指す．
- ここではgRPCではなく，Go言語が標準ライブラリとして保有するRPCライブラリを使用することを想定している．
- 課題1のプログラムをディレクトリごとコピーして，適宜修正を加える形式で実装すると良い．
- 参考: [net/rpc](https://pkg.go.dev/net/rpc)

## 課題 3: Replication
TBU
