# 1brcgo

[1 billion rows challenge](https://github.com/gunnarmorling/1brc) implemented in Go.

## Solutions

### Naive scanner

First solution that come to mind. Read file line by line and process lines sequentially. 

Very slow and uses insane amount of memory (32GB RAM + 15GB swap).

```sh
real	2m44.457s
user	2m17.534s
sys	0m20.253s
```

### Naive read all

Read the whole file to array of bytes and go through it sequentially

Slow, but much better than scanner (and fits into memory).

```sh

```