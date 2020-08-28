#!/bin/bash

cp genesis.json node0 
cp genesis.json node1 
cp genesis.json node2 
cp genesis.json node3 
cp genesis.json node4 

cp static-nodes.json node0/data/ 
cp static-nodes.json node1/data/ 
cp static-nodes.json node2/data/ 
cp static-nodes.json node3/data/ 
cp static-nodes.json node4/data/ 

cp nodekeys/0/nodekey node0/data/geth 
cp nodekeys/1/nodekey node1/data/geth
cp nodekeys/2/nodekey node2/data/geth 
cp nodekeys/3/nodekey node3/data/geth 
cp nodekeys/4/nodekey node4/data/geth