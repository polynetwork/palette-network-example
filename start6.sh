#!/bin/bash
killall -INT geth

sleep 5s;

rm -rf node0/node.log
rm -rf node1/node.log
rm -rf node2/node.log
rm -rf node3/node.log
rm -rf node4/node.log
rm -rf node5/node.log

sleep 1s;

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

cd ../node5
PRIVATE_CONFIG=ignore nohup geth --datadir data --nodiscover --istanbul.blockperiod 5 --syncmode full --mine --minerthreads 1 --verbosity 5 --networkid 10 --rpc --rpcaddr 0.0.0.0 --rpcport 22005 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,istanbul --emitcheckpoints --port 30305 2>>node.log &

ps -ef|grep geth