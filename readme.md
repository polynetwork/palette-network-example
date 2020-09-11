## create palette network with scratch

    this document described how to create a palette network with scratch.

# 1.admin account
```bash
geth --datadir=admin account new
```

Public address of the key:   0xf3A9d42C01635A585f1721463842F8936075105F 
Path of the secret key file: admin/keystore/UTC--2020-09-11T02-29-49.024005000Z--f3a9d42c01635a585f1721463842f8936075105f 

# 2.mkdir node dir
```bash
mkdir node0 node1 node2 node3 node4
```

# 3.gen nodekeys, static-node.json, genesis.json
```bash
mkdir setup
cd setup
../istanbul-tools/build/bin/istanbul setup --num 5 --nodes --quorum --save --verbose
```

```xml
validators
{
	"Address": "0x88210e568ffd4c1a1d89c8f8fdd454cdde1ed614",
	"Nodekey": "77a5f6abf01528962fc3583ae008b5cff601c24abd72984d91ecbf73a2c5eeaa",
	"NodeInfo": "enode://a99091a27dca6bd4424b320b05475a6957da834f604f21a40b0289d81c06f3c625c89c8c9c7090b2bcd7efcfac9e44cefad29a73439f97eb3718aad061c6573d@0.0.0.0:30303?discport=0"
}
{
	"Address": "0xca4f5700022fb42da9440c6d7600ca2ed862cf98",
	"Nodekey": "0ac17450be0f73ee95b73e51cfdc4a3b66b2cdb50cc7d8dc2ebd1fdc8b79a202",
	"NodeInfo": "enode://3388bdf9486a3cb254b78e28aba9756f8c5f0a9a32f74c8cf5846d2c346b24e9d9e3149714036aacf3128e2d2ded309d5807e4a6d8508a9c7b54aa00b48073e6@0.0.0.0:30303?discport=0"
}
{
	"Address": "0xdc11f8102f4e2991f7b4887f6efa7521bdddf24e",
	"Nodekey": "1b64433e0a47b71a6abe0d0fe068efc2969f8f75c33876f161c2cdf68f3588c2",
	"NodeInfo": "enode://c43fcb7f4da368d6d0a1210337e254c8d11ac968e97acb93e3922fc581c35c200746ea01900e872e8da80f52344b266abf89cb2ca198b7cf22be7b6c1c2e92a8@0.0.0.0:30303?discport=0"
}
{
	"Address": "0x715d8b50ad9acfdde186c247b8e2930d122aa901",
	"Nodekey": "ef762392d83b0fb2c94e7f17e9ec0514a7009cc0beb522d407008439a83cef60",
	"NodeInfo": "enode://d705634fa93059525f08eda19f17cba858dd03b635f00746b354fb1c524996cc63619453e88c1e11009fc487e336896ea8288bb2f99a165f50821a474fc08ebe@0.0.0.0:30303?discport=0"
}
{
	"Address": "0x2113bbc1cf2c09e5b166857250282a2677e00800",
	"Nodekey": "beff30ce180cefa78ba841c81863f22ad3b258ae6053d647fed9922f359d6fe6",
	"NodeInfo": "enode://aaa86f39b4359a00973b6a903f2dd207226b9ff4da66243292e2db543d9b198d89b8e008f299bc10b7d1391813b80829765af089faf2327041ac3a593124d53e@0.0.0.0:30303?discport=0"
}



static-nodes.json
[
	"enode://a99091a27dca6bd4424b320b05475a6957da834f604f21a40b0289d81c06f3c625c89c8c9c7090b2bcd7efcfac9e44cefad29a73439f97eb3718aad061c6573d@0.0.0.0:30303?discport=0",
	"enode://3388bdf9486a3cb254b78e28aba9756f8c5f0a9a32f74c8cf5846d2c346b24e9d9e3149714036aacf3128e2d2ded309d5807e4a6d8508a9c7b54aa00b48073e6@0.0.0.0:30303?discport=0",
	"enode://c43fcb7f4da368d6d0a1210337e254c8d11ac968e97acb93e3922fc581c35c200746ea01900e872e8da80f52344b266abf89cb2ca198b7cf22be7b6c1c2e92a8@0.0.0.0:30303?discport=0",
	"enode://d705634fa93059525f08eda19f17cba858dd03b635f00746b354fb1c524996cc63619453e88c1e11009fc487e336896ea8288bb2f99a165f50821a474fc08ebe@0.0.0.0:30303?discport=0",
	"enode://aaa86f39b4359a00973b6a903f2dd207226b9ff4da66243292e2db543d9b198d89b8e008f299bc10b7d1391813b80829765af089faf2327041ac3a593124d53e@0.0.0.0:30303?discport=0"
]


genesis.json
{
    "config": {
        "chainId": 10,
        "homesteadBlock": 0,
        "eip150Block": 0,
        "eip150Hash": "0x0000000000000000000000000000000000000000000000000000000000000000",
        "eip155Block": 0,
        "eip158Block": 0,
        "byzantiumBlock": 0,
        "constantinopleBlock": 0,
        "istanbul": {
            "epoch": 30000,
            "policy": 0,
            "ceil2Nby3Block": 0
        },
        "txnSizeLimit": 64,
        "maxCodeSize": 0,
        "isQuorum": true
    },
    "nonce": "0x0",
    "timestamp": "0x5f5ae674",
    "extraData": "0x0000000000000000000000000000000000000000000000000000000000000000f8aff8699488210e568ffd4c1a1d89c8f8fdd454cdde1ed61494ca4f5700022fb42da9440c6d7600ca2ed862cf9894dc11f8102f4e2991f7b4887f6efa7521bdddf24e94715d8b50ad9acfdde186c247b8e2930d122aa901942113bbc1cf2c09e5b166857250282a2677e00800b8410000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c0",
    "gasLimit": "0xe0000000",
    "difficulty": "0x1",
    "mixHash": "0x63746963616c2062797a616e74696e65206661756c7420746f6c6572616e6365",
    "coinbase": "0x0000000000000000000000000000000000000000",
    "alloc": {
        "2113bbc1cf2c09e5b166857250282a2677e00800": {
            "balance": "0x446c3b15f9926687d2c40534fdb564000000000000"
        },
        "715d8b50ad9acfdde186c247b8e2930d122aa901": {
            "balance": "0x446c3b15f9926687d2c40534fdb564000000000000"
        },
        "88210e568ffd4c1a1d89c8f8fdd454cdde1ed614": {
            "balance": "0x446c3b15f9926687d2c40534fdb564000000000000"
        },
        "ca4f5700022fb42da9440c6d7600ca2ed862cf98": {
            "balance": "0x446c3b15f9926687d2c40534fdb564000000000000"
        },
        "dc11f8102f4e2991f7b4887f6efa7521bdddf24e": {
            "balance": "0x446c3b15f9926687d2c40534fdb564000000000000"
        }
    },
    "number": "0x0",
    "gasUsed": "0x0",
    "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000"
}
```

Notice: 
. [x] modify ip and port in setup/static-nodes.json 
. [x] modify genesis.json, add admin public address in config 

# 4.copy setup files in nodes
```bash
cp setup/genesis.json node0
cp setup/genesis.json node1
cp setup/genesis.json node2
cp setup/genesis.json node3
cp setup/genesis.json node4

cp setup/static-nodes.json node0/data/
cp setup/static-nodes.json node1/data/
cp setup/static-nodes.json node2/data/
cp setup/static-nodes.json node3/data/
cp setup/static-nodes.json node4/data/

cp setup/0/nodekey node0/data/geth
cp setup/1/nodekey node1/data/geth
cp setup/2/nodekey node2/data/geth
cp setup/3/nodekey node3/data/geth
cp setup/4/nodekey node4/data/geth
```

# 5.init geth node
```bash
cd node0
geth --datadir data init genesis.json

cd ../node1/
geth --datadir data init genesis.json

cd ../node2/
geth --datadir data init genesis.json

cd ../node3/
geth --datadir data init genesis.json

cd ../node4/
geth --datadir data init genesis.json
```

# 6.start up all nodes
```bash
cd node0
PRIVATE_CONFIG=ignore nohup geth --datadir data --nodiscover --istanbul.blockperiod 5 --syncmode full --mine --minerthreads 1 --verbosity 5 --networkid 10 --rpc --rpcaddr 0.0.0.0 --rpcport 22000 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,istanbul --emitcheckpoints --port 30300 2>>node.log &

cd ../node1
PRIVATE_CONFIG=ignore nohup geth --datadir data --nodiscover --istanbul.blockperiod 5 --syncmode full --mine --minerthreads 1 --verbosity 5 --networkid 10 --rpc --rpcaddr 0.0.0.0 --rpcport 22001 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,istanbul --emitcheckpoints --port 30301 2>>node.log &

cd ../node2
PRIVATE_CONFIG=ignore nohup geth --datadir data --nodiscover --istanbul.blockperiod 5 --syncmode full --mine --minerthreads 1 --verbosity 5 --networkid 10 --rpc --rpcaddr 0.0.0.0 --rpcport 22002 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,istanbul --emitcheckpoints --port 30302 2>>node.log &

cd ../node3
PRIVATE_CONFIG=ignore nohup geth --datadir data --nodiscover --istanbul.blockperiod 5 --syncmode full --mine --minerthreads 1 --verbosity 5 --networkid 10 --rpc --rpcaddr 0.0.0.0 --rpcport 22003 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,istanbul --emitcheckpoints --port 30303 2>>node.log &

cd ../node4
PRIVATE_CONFIG=ignore nohup geth --datadir data --nodiscover --istanbul.blockperiod 5 --syncmode full --mine --minerthreads 1 --verbosity 5 --networkid 10 --rpc --rpcaddr 0.0.0.0 --rpcport 22004 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,istanbul --emitcheckpoints --port 30304 2>>node.log &
```