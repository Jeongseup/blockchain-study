# Tendermint study

- 텐더민트 인덱싱
Transactions are indexed by TxResult.Events and blocks are indexed by Response(Begin|End)Block.Events
또한, TxResult.Hash() 값을 primary key로 갖는다.

그럼 이 이벤트는 어디서 발생하냐? -> ABCI에서 [스펙정의](https://github.com/tendermint/spec/blob/master/spec/abci/abci.md#events)

Operators can configure indexing via the [tx_index] section. The indexer field takes a series of supported indexers. If null is included, indexing will be turned off regardless of okavther values provided
--> 그럼 이거 데이터를 애초에 로컬 디비에 꽂아둘 수 있는지 테스트해보고 싶은데..?

``` toml
# config.toml

#######################################################
###   Transaction Indexer Configuration Options     ###
#######################################################
[tx-index]

# The backend database list to back the indexer.
# If list contains "null" or "", meaning no indexer service will be used.
#
# The application will set which txs to index. In some cases a node operator will be able
# to decide which txs to index based on configuration set in the application.
#
# Options:
#   1) "null"
#   2) "kv" (default) - the simplest possible indexer, backed by key-value storage (defaults to levelDB; see DBBackend).
#   3) "psql" - the indexer services backed by PostgreSQL.
# When "kv" or "psql" is chosen "tx.height" and "tx.hash" will always be indexed.
indexer = ["kv"]

# The PostgreSQL connection configuration, the connection format:
#   postgresql://<user>:<password>@<host>:<port>/<db>?<opts>
psql-conn = ""
```
