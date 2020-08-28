## Palette network

### steps:
1. get quorum souce code and build
```bash
cd ~
git clone git@github.com:polynetwork/quorum.git
cd quorum 
make all
```

2. make istanbul-tools
```bash
cd ~/quorum/istanbul-tools
make
```

3. set environment
```bash
cd ~
vi ~/.bash_profile

# add quorum
export QUORUM=$HOME/quorum/build
export PATH=$QUORUM/bin:$PATH
:wq!

source ~/.bash_profile
```

4. start network
```bash
cd ~/quorum
chmod +x start5.sh
./start5.sh
```

5. test
modify the keystore director in simple_transfer/main.go

```bash
cd ~/quorum
go run simple_transfer/main.go
```