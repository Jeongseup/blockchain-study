#!/usr/bin/env bash

set -e

shopt -s expand_aliases

alias nsd=$(pwd)/nsd
alias namecli=$(pwd)/namecli

rm -rf ~/.nsd
rm -rf ~/.namecli

nsd init test --chain-id=namechain

namecli config output json
namecli config indent true
namecli config trust-node true
namecli config chain-id namechain
namecli config keyring-backend test

jackseed='sport eternal despair major raccoon arrest zoo magic measure ensure relief multiply priority paper reveal scrap know word lawn wagon elbow trumpet unveil young'
aliceseed='fork digital rare coin permit couch business energy session arm icon profit rule peasant ability walk office open hawk churn special over excite vessel'
fooseed='laugh circle useless shift maze quantum ribbon auction check pledge toilet ripple hobby radio indoor nice segment torch crucial ecology omit jaguar aunt later'
barseed='frown attack build change parade way brand advice voice basket panda marble reduce renew paper suspect arch club ketchup toddler include deal leader float'

namecli keys add --recover jack <<<"${jackseed}"
namecli keys add --recover alice <<<"${aliceseed}"
namecli keys add --recover foo <<<"${fooseed}"
namecli keys add --recover bar <<<"${barseed}"

nsd add-genesis-account $(namecli keys show jack -a) 1000nametoken,100000000stake
nsd add-genesis-account $(namecli keys show alice -a) 1000nametoken,100000000stake
nsd add-genesis-account $(namecli keys show foo -a) 10000nametoken,1000000stake
nsd add-genesis-account $(namecli keys show bar -a) 10000nametoken,1000000stake

nsd gentx --name jack --keyring-backend test

echo "Collecting genesis txs..."
nsd collect-gentxs

echo "Validating genesis file..."
nsd validate-genesis

for keyname in jack alice foo bar; do
  namecli keys show --address $keyname
done