# zig

- ターミナルを zig で実装した話を聞いて興味が出てきた

## Links
- https://www.publickey1.jp/blog/24/ghostty_1012.html
- https://ziglang.org/documentation/0.13.0/#Hello-World

## Install
バージョン管理できないから brew で入れないほうがいいんだろうけど面倒なので。

```console
$ brew install zig
$ zig version
0.13.0
```

## Build App
```console
$ zig build-exe hello.zig # this create `hello`, `hello.o`
$ ./hello
Hello, world!
```
