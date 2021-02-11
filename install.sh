#!/bin/bash
echo
echo "> building go binary..."
# build gpshr
go build

echo "> copying files..."
# make directory
mkdir -p ~/.gpshr
chmod +x ~/.gpshr
chmod +x ~/.gpshr/gpshr

# move the binary
cp gpshr ~/.gpshr
rm gpshr
cp -r sounds ~/.gpshr

echo "> setting sh/bash/zsh path..."
# bash PATH
echo "export PATH="~/.gpshr"":'"$PATH"' >>~/.profile
source ~/.profile

echo "> setting fish path..."
# fish PATH
mkdir -p ~/.config/fish/conf.d
touch ~/.config/fish/conf.d/gpshr.fish
echo "set PATH ~/.gpshr" : '"$PATH"' >> ~/.config/fish/conf.d/gpshr.fish
source ~/.config/fish/conf.d/gpshr.fish

echo
GREEN='\033[0;32m'
GREY='\033[0;36m'
NC='\033[0m' # No Color
echo -e "${GREEN}! INSTALLATION DONE${NC}"
echo