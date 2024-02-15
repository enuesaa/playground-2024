# nomad
## Commands
```bash
nomad agent -dev
nomad run app.nomad.hcl
nomad status
nomad stop app
```

## Nomad Agent を起動
Nomad Agent を dev モードで起動する。通常は Nomad Server と Nomad Client とで複数台構成するようだが、dev モードでは1台になっている
```bash
nomad agent -dev
```

## コンテナに入る
```console
$ nomad job allocs app
ID        Node ID   Task Group  Version  Desired  Status    Created     Modified
0a36e239  c8a2e3f1  web         9        run      running   7m33s ago   7m20s ago
```

```console
$ nomad exec 0a36e239 bash
root@5b3f8a1d:/# ls
alloc  boot  docker-entrypoint.d   etc	 lib	media  opt   root  sbin     srv  tmp  var
bin    dev   docker-entrypoint.sh  home  local	mnt    proc  run   secrets  sys  usr
```
