# lclとは

Laravel、WordPress、Djangoのローカル開発環境を容易に作成できるコマンドです。
> <img width="500" alt="example" src="https://storage.googleapis.com/ken109-image/lcl-example.gif">

# インストール

1. Dockerをインストールする
   ```bash
   brew cask install docker
   ```

2. 以下のアイコンからDockerを起動し、Dashboardを開く

> <img width="150" alt="docker-icon" src="https://storage.googleapis.com/ken109-image/docker-icon.png">
> <img width="150" alt="docker-open-dashboard" src="https://storage.googleapis.com/ken109-image/docker-open-dashboard.png">

3. Dockerの`File Sharing`に`/usr/local/etc`を追加する

> <img width="500" alt="docker-file-sharing" src="https://storage.googleapis.com/ken109-image/docker-file-sharing.png">

4. インストール
   ```bash
   brew install ken109/tap/lcl
   lcl init
   lcl base start [mysql, mysql5, dns, nginx, mongo, redis] // 指定しなかったら全て起動する
   ```

5. DNSサーバーの先頭に127.0.0.1を追加する

# lclコマンドの使い方

* lcl
    * init
    * update
    * [base](https://github.com/ken109/lcl/wiki/Base)
        * [start](https://github.com/ken109/lcl/wiki/Base#%E3%83%99%E3%83%BC%E3%82%B9%E3%82%B3%E3%83%B3%E3%83%86%E3%83%8A%E8%B5%B7%E5%8B%95)
        * [stop](https://github.com/ken109/lcl/wiki/Base#%E3%83%99%E3%83%BC%E3%82%B9%E3%82%B3%E3%83%B3%E3%83%86%E3%83%8A%E7%B5%82%E4%BA%86)
    * start
        * wp
        * la
        * dj
    * stop
