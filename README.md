# hanzi-count

Command-line tool which lists the frequency of Chinese characters\* within text
files.

\* Only a common subset of the CJK Unified Ideographs unicode block is counted
for simplicity and performance reasons. This includes simplified and
traditional characters, as well as Kanji. Some less common characters may be
excluded.


## Example

```shell
> cat file.txt
你可以编辑这里的代码！
点击这里然后开始输入。
> hanzi-count file.txt
代 1
以 1
你 1
入 1
击 1
可 1
后 1
始 1
开 1
点 1
然 1
的 1
码 1
编 1
辑 1
输 1
这 2
里 2
```


## Usage

```
USAGE:
    hanzi-count [FLAGS] <FILE>...

FLAGS:
    -h, -help, --help        Prints help information
    -excludes, --excludes    File containing characters to exclude from the output

ARGS:
    <FILE>...    Path of the files containing characters to count
```

## Install

Currently, pre-compiled binaries of hanzi-count aren't being distributed. You
can install it with with the [`go` command](https://golang.org/doc/install) by
running:

```shell
go install github.com/rsookram/hanzi-count/cmd/hanzi-count@latest
```


## Build

hanzi-count can be built from source by cloning this repository and using the
`go` command.

```shell
git clone https://github.com/rsookram/hanzi-count
cd hanzi-count
go build ./cmd/hanzi-count
```


## Test

The test suite can be run like typical go tests:

```shell
go test ./...
```


## License

[MIT](LICENSE)
