
<h1 align="center"> <a href="#english">english</a> |<a href="#japanese">日本語</a></h1>

![README LOGO](_img/bak.png)
# gpshr | git pusher

y not adding ugly sounds when u git push

<h1 align="left" id="english"> 🇺🇸english<h1>

## dependancies

    - golang
    - afplay (for darwin like OSX)
    - aplay (for linux OS)

## installation

run this command below:

```
$ chmod +x install.sh
$ ./install.sh
```
default, this script add path **both bash/zsh & fish**

## usage

- ```gpshr``` 

    see help

- ```gpshr -install .``` 

    install sound sctipts into current directory

- ```gpshr -uninstall .``` 

    uninstall sound script from current directory

- ```gpshr -import foo.mp3``` 

    import sound from foo.mp3


## doesnt work? 
try import them into each shell settings:
### bash

``` ~/.profile ```

``` 
export PATH="~/.gpshr" : "$PATH" 
```

### fish
    
``` ~/.config/fish/conf.d/gpshr.fish```

``` 
set PATH ~/.gpshr : "$PATH" 
```


<h1 align="left" id="japanese"> 🇯🇵日本語<h1>

## 依存パッケージ

    - golang
    - afplay (darwinおよびOSXを使ってる人はこれ)
    - aplay (linuxを使ってる人はこれ)

## インストール方法

クローンしたフォルダーで以下のコマンドを叩いてください:

```
$ chmod +x install.sh
$ ./install.sh
```
**bash, zsh & fish**にgpshrのパスが通されます
実際にスクリプトをインストールして、音を鳴らしてみましょう！！！楽しみ！！！

## 使いかた

- ```gpshr``` 

    コマンドのヘルプが表示されます

- ```gpshr -install フォルダー``` 

    指定したフォルダーにgpshrがインストールされます
    git initされているフォルダーを指定してください!

- ```gpshr -uninstall フォルダー``` 

    指定したフォルダーからgpshrがアンインストールされます

- ```gpshr -import foo.mp3``` 

    ~/.gpshr/sounds　に音声ファイルをインポートします


## 動きません！
パスが通ってないのかもしれません、使用されているシェルに従って以下のスクリプトを設置してください:
### bash

``` ~/.profile ```

``` 
export PATH="~/.gpshr" : "$PATH" 
```

### fish
    
``` ~/.config/fish/conf.d/gpshr.fish```

``` 
set PATH ~/.gpshr : "$PATH" 
```