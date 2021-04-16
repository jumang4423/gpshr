<h1 align="center"> <a href="#english">english</a> |<a href="#japanese">日本語</a></h1>

![README LOGO](_img/bak.png)
# gpshr | git pusher

y not adding ugly sounds when u git push

<h1 align="left" id="english"> 🇺🇸english<h1>

## dependancies for build

    - golang
    - afplay (for darwin like OSX) or aplay (for linux OS)

## development with a docker
```bash
# build go image
sudo docker-compose build
# launch image on background process
sudo docker-compose up -d
# exec go command using docker envinroment
sudo docker-compose exec gpshr go run *.go
```

## installation

run this command below:

```bash
chmod +x install.sh
./install.sh
```
default, this script add path **both bash/zsh & fish**

## usage

- ```gpshr``` 

    see help

- ```gpshr -(un)install [directory] -hooks [hooks]``` 

    select [hooks] hooks (push, commit and checkout) then (un)install sound sctipts into [directory]


- ```gpshr -import [music file]``` 

    import sound from [music file]


## doesnt work? 
try import them into each shell settings:
### bash | zsh

``` ~/.profile ```

```bash
export PATH="~/.gpshr" : "$PATH" 
```

### fish
    
``` ~/.config/fish/conf.d/gpshr.fish```

```bash
set PATH ~/.gpshr : "$PATH" 
```


<h1 align="left" id="japanese"> 🇯🇵日本語<h1>

## ビルドのための依存パッケージ

    - golang
    - afplay (darwinおよびOSXを使ってる人はこれ) or aplay (linuxを使ってる人はこれ)

## dockerをつかった開発環境
```bash
# イメージをビルド
sudo docker-compose build
# バックグラウンドでイメージを立ち上げる
sudo docker-compose up -d
# ドッカー環境でgoを使ってソフトを動かす
sudo docker-compose exec gpshr go run *.go
```

## インストール方法

クローンしたフォルダーで以下のコマンドを叩いてください:

```bash
chmod +x install.sh
./install.sh
```
**bash, zsh & fish**にgpshrのパスが通されます
実際にスクリプトをインストールして、音を鳴らしてみましょう！！！楽しみ！！！

## 使いかた

- ```gpshr``` 

    コマンドのヘルプが表示されます

- ```gpshr -(un)install [ディレクトリ] -hooks [フックス]``` 

    [ディレクトリ]で、[フックス]（push, commitまたはcheckoutを入れてください）された時に音がなる様に設定します。

- ```gpshr -import [音楽ファイル]``` 

    ~/.gpshr/sounds　に[音楽ファイル]をインポートします


## 動きません！
パスが通ってないのかもしれません、使用されているシェルに従って以下のスクリプトを設置してください:
### bash | zsh

``` ~/.profile ```

```bash
export PATH="~/.gpshr" : "$PATH" 
```

### fish
    
``` ~/.config/fish/conf.d/gpshr.fish```

```bash
set PATH ~/.gpshr : "$PATH" 
```