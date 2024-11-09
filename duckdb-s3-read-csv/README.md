# duckdb

- cli or python, r, node.js のバインディングとして使えるが、ここでは cli で生の sql を実行
- see https://github.com/enuesaa/playground-2023/tree/main/duckdb-pandas

## Install
```bash
brew install duckdb
```

## 1. Create Database file
永続化に必要
```bash
duckdb data.duckdb
```

## 2. Load AWS Extension
- see https://duckdb.org/docs/extensions/aws.html
- see https://zenn.dev/jigjp_engineer/articles/a9e0bcfa536ad2
- see https://road288.hatenablog.com/entry/2024/11/06/113954

```console
➜ duckdb data.duckdb

v1.1.3 19864453f7
Enter ".help" for usage hints.
D INSTALL aws;
D LOAD aws;
D INSTALL httpfs;
D LOAD httpfs;
D CREATE SECRET (
      TYPE S3,
      PROVIDER CREDENTIAL_CHAIN
  );
┌─────────┐
│ Success │
│ boolean │
├─────────┤
│ true    │
└─────────┘
```

## 3. Create Table
```console
D CREATE TABLE logs
  AS (
      SELECT *
      FROM read_csv(
          's3://bucket-name/testdata.csv',
          columns={
              'Timestamp': 'VARCHAR',
              'Method': 'VARCHAR',
              'Path': 'VARCHAR',
              'HTTP': 'VARCHAR',
              'Code': 'VARCHAR',
              'Size': 'VARCHAR',
              'Referrer': 'VARCHAR',
          })
  );
```

## 4. select
```console
D select * from logs;
┌──────────────────────┬─────────┬───────────┬──────────┬─────────┬─────────┬───────────┐
│      Timestamp       │ Method  │   Path    │   HTTP   │  Code   │  Size   │ Referrer  │
│       varchar        │ varchar │  varchar  │ varchar  │ varchar │ varchar │  varchar  │
├──────────────────────┼─────────┼───────────┼──────────┼─────────┼─────────┼───────────┤
│ 2024-11-09T10:20:30Z │ GET     │ /home     │ HTTP/1.1 │ 200     │ 512     │ -         │
│ 2024-11-09T10:21:00Z │ POST    │ /login    │ HTTP/1.1 │ 302     │ 256     │ /home     │
│ 2024-11-09T10:22:15Z │ GET     │ /about    │ HTTP/1.1 │ 404     │ 128     │ /home     │
│ 2024-11-09T10:23:45Z │ GET     │ /contact  │ HTTP/1.1 │ 200     │ 512     │ /about    │
│ 2024-11-09T10:24:10Z │ GET     │ /services │ HTTP/1.1 │ 200     │ 768     │ /home     │
│ 2024-11-09T10:25:20Z │ POST    │ /submit   │ HTTP/1.1 │ 201     │ 1024    │ /contact  │
│ 2024-11-09T10:26:30Z │ GET     │ /blog     │ HTTP/1.1 │ 200     │ 640     │ /home     │
│ 2024-11-09T10:27:45Z │ GET     │ /faq      │ HTTP/1.1 │ 404     │ 256     │ /home     │
│ 2024-11-09T10:28:10Z │ GET     │ /services │ HTTP/1.1 │ 200     │ 512     │ /blog     │
│ 2024-11-09T10:29:00Z │ POST    │ /login    │ HTTP/1.1 │ 302     │ 256     │ /services │
├──────────────────────┴─────────┴───────────┴──────────┴─────────┴─────────┴───────────┤
│ 10 rows                                                                     7 columns │
└───────────────────────────────────────────────────────────────────────────────────────┘
D
D select * from logs where code = 200;
┌──────────────────────┬─────────┬───────────┬──────────┬─────────┬─────────┬──────────┐
│      Timestamp       │ Method  │   Path    │   HTTP   │  Code   │  Size   │ Referrer │
│       varchar        │ varchar │  varchar  │ varchar  │ varchar │ varchar │ varchar  │
├──────────────────────┼─────────┼───────────┼──────────┼─────────┼─────────┼──────────┤
│ 2024-11-09T10:20:30Z │ GET     │ /home     │ HTTP/1.1 │ 200     │ 512     │ -        │
│ 2024-11-09T10:23:45Z │ GET     │ /contact  │ HTTP/1.1 │ 200     │ 512     │ /about   │
│ 2024-11-09T10:24:10Z │ GET     │ /services │ HTTP/1.1 │ 200     │ 768     │ /home    │
│ 2024-11-09T10:26:30Z │ GET     │ /blog     │ HTTP/1.1 │ 200     │ 640     │ /home    │
│ 2024-11-09T10:28:10Z │ GET     │ /services │ HTTP/1.1 │ 200     │ 512     │ /blog    │
└──────────────────────┴─────────┴───────────┴──────────┴─────────┴─────────┴──────────┘
```

## NOTE
testdata is dummy.
