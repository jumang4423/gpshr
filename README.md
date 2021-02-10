![README LOGO](_img/bak.png)

# gpshr
y not adding ugly sounds when u git push

## dependancies

    - golang

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

- ```gpshr -install foo``` 

    install sound sctipts into foo directory

- ```gpshr -uninstall foo``` 

    uninstall sound script from foo

- ```gpshr -import foo``` 

    import sound from foo<!-- 


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