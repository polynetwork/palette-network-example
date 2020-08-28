## Palette network example

### steps:
1. get palette souce code and build
```bash
cd ~
git clone git@github.com:polynetwork/palette.git
cd palette
make all
```

2. set environment
```bash
cd ~
vi ~/.bash_profile

# add palette
export PALETTE=/your/gopath/palette/build
export PATH=$PALETTE/bin:$PATH
:wq!

source ~/.bash_profile

which geth
# make sure the `geth` path is correct
```

3. start network
```bash
cd /your/gopath/palette-network-example
chmod +x start5.sh
./start5.sh
```

4. test
```bash

# first, pls wait seconds until the p2p network connection is successful
# second, modify the keystore director in simple_transfer/main.go
cd /your/gopath/palette-network-example/simple_transfer
go mod download
go run main.go

```
