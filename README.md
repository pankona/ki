# ki

This is tree like application written in Go

# Install

```bash
$ go get -u github.com/pankona/ki/cmd/ki
```

# Usage

Available options:

```
Usage of ki:
  -a	specify to include hidden directory [default: false]
  -c int
    	specify concurrent num [default: 0]
  -d	specify to include only directories [default: false]
  -p	specify to enable plane rendering [default: false]
  -with-profile
    	specify to enable profiling [default: false]
```

Example:

```bash
# Specify directory (s) to show directory tree
$ ki ./hoge
./
|-- fuga/
|-- piyo.txt
`-- zzz/
    |-- piyo.txt
    `-- qqq.txt
```

Use with peco:

```bash
# list directories and change directory via peco
$ cd $(ki -d -p . | peco)
```

# License

MIT

# Author

Yosuke Akatsuka (a.k.a pankona)
