# hanzi-count

Command-line tool which lists the frequency of Chinese characters\* within text
files.

\* Only a common subset of the CJK Unified Ideographs unicode block is counted
for simplicity and performance reasons. This includes simplified and
traditional characters, as well as Kanji. Some less common characters are
excluded.


## Example

The basic usage of `hanzi-count` looks like:

```shell
> cat file.txt
你可以编辑这里的代码！
点击这里然后开始输入。
> hanzi-count file.txt
里 2
这 2
点 1
入 1
击 1
可 1
后 1
始 1
开 1
你 1
然 1
的 1
码 1
编 1
辑 1
输 1
以 1
代 1
```

The output is sorted by frequency descending.

If you use an SRS to keep track of characters you know, you can query it to get
characters to exclude from the output. For example:

```shell
hanzi-count --excludes <(sqlite3 srs.db 'SELECT front FROM Card WHERE deckId = 1') file.txt
```


## Usage

```
USAGE:
    hanzi-count [FLAGS] <FILE>...

FLAGS:
    -h, -help, --help        Prints help information
    -excludes, --excludes    File containing characters to exclude from the output

ARGS:
    <FILE>...    Paths of the files containing characters to count
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
