## Palette network

### steps:
1. get quorum souce code and build
```bash
cd ~
git clone git@github.com:polynetwork/quorum.git
cd quorum 
make all
```

2. set environment
```bash
cd ~
vi ~/.bash_profile

# add quorum
export QUORUM=$HOME/quorum/build
export PATH=$QUORUM/bin:$PATH
:wq!

source ~/.bash_profile
```

3. start network
```bash
cd /your/gopath/palette-network-example
chmod +x start5.sh
./start5.sh
```

4. test
modify the keystore director in simple_transfer/main.go

```bash
cd /your/gopath/palette-network-example
go run simple_transfer/main.go
```
