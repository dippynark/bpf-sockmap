# bpf-sockmap

bpf-sockmap uses [BPF_PROG_TYPE_SK_SKB](https://lwn.net/Articles/731133/) BPF programs to create a simple telnet echo server. It is heavily inspired by the [Cloudflare blog](https://github.com/cloudflare/cloudflare-blog/blob/master/2019-02-tcp-splice/echo-sockmap-kern.c).

## Quickstart

[Vagrant](https://www.vagrantup.com/) can be used to spin up a virtual environment to test the BPF programs. The environment depends on [VirtualBox](https://www.virtualbox.org/wiki/Downloads) but other [providers](https://www.vagrantup.com/docs/providers/) exist.

```
$ vagrant box list | grep ubuntu/bionic64 || vagrant box add ubuntu/bionic64
$ vagrant up
$ vagrant ssh
$ cd /vagrant
$ make build
$ make run
...
2019/04/03 00:53:12 listening on address: 0.0.0.0:12345
# in another terminal watch debug output
$ sudo cat /sys/kernel/debug/tracing/trace_pipe
# in yet another terminal start a telnet session
$ telnet 127.0.0.1 12345
Trying 127.0.0.1...
Connected to 127.0.0.1.
Escape character is '^]'.
Hello!
Hello!
Bye
Bye
^]q

telnet> q
Connection closed.
```

## Debug

The generated object file can be inspected using `llvm-objdump`

```
llvm-objdump -S ./bpf/bpf_sockmap.o
```
