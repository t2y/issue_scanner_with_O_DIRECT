# learn read with O_DIRECT

## learn how to read a file with O_DIRECT

```bash
$ make
go build -o main main.go

$ ./main --help
Usage of ./main:
  -direct    set O_DIRECT
```

## reference

* [open(2)](http://man7.org/linux/man-pages/man2/open.2.html)
* [read(2)](http://man7.org/linux/man-pages/man2/read.2.html)
* [Why does O_DIRECT require I/O to be 512-byte aligned?](https://www.quora.com/Why-does-O_DIRECT-require-I-O-to-be-512-byte-aligned)

