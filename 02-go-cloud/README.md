Getting started with [go-cloud](https://github.com/google/go-cloud/)
 for a simple go app

With our buckets created locally, in S3 and GCS and our credentials for both set up,
we'll build the program first:

``` shell
$ go build -o blobber
```

Then, we will send `gopher.png` (in the same directory as this README) to the local bucket:

``` shell
$ ./blobber -cloud file gopher.png /gophers/first.png
```

Then, we will retrieve `/gophers/first.png` (in the same directory as this README) from the local bucket and open it in Preview (OSX):

``` shell
$ ./blobber -cloud file -d /gophers/first.png /dev/stdout | open -a Preview.app -f
```
