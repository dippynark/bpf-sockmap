# bpf-sockmap

```
vagrant up
vagrant ssh
cd /vagrant
make build
make run
```

## Debug

The generated object file can be viewed using `llvm-objdump`

```
llvm-objdump -S ./bpf/bpf_sockmap.o
```
