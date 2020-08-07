# lclコマンドの使い方

* lcl
  * [base](https://github.com/ken109/lcl/wiki/Base)
    * [start](https://github.com/ken109/lcl/wiki/Base#%E3%83%99%E3%83%BC%E3%82%B9%E3%82%B3%E3%83%B3%E3%83%86%E3%83%8A%E8%B5%B7%E5%8B%95)
    * [stop](https://github.com/ken109/lcl/wiki/Base#%E3%83%99%E3%83%BC%E3%82%B9%E3%82%B3%E3%83%B3%E3%83%86%E3%83%8A%E7%B5%82%E4%BA%86)
  * start
    * wp
    * la
    * dj
  * stop

# インストール
1. Dockerをインストールする
   ```bash
   brew cask install docker
   brew cask install docker-toolbox
   ```

2. 以下のアイコンからDockerを起動する
> <img width="150" alt="docker-icon" src="https://storage.googleapis.com/ken109-image/docker-icon.png">

3. 以下をクリックしてDashboardを開く
> <img width="150" alt="docker-open-dashboard" src="https://storage.googleapis.com/ken109-image/docker-open-dashboard.png">

3. Dockerの`File Sharing`に`/usr/local/etc`を追加する
> <img width="500" alt="docker-file-sharing" src="https://storage.googleapis.com/ken109-image/docker-file-sharing.png">

4. インストール
   ```bash
   brew tap ken109/lcl
   brew install lcl
   ```

5. DNSサーバーの先頭に127.0.0.1を追加する
