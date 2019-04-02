# sockmap

```
vagrant up
make bpf
```

## Debug

The generated object file can be viewed using `llvm-objdump`

```
llvm-objdump -S ./bpf/bpf_tty.o
```
