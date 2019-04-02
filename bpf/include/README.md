# bpf include

```
# https://lwn.net/Articles/507794/
#bpf.h: #include <uapi/linux/bpf.h>
#bpf.h: cp -a /usr/src/linux-headers-4.18.0-16/include/uapi/linux/bpf.h /vagrant/bpf/include/
bpf.h: https://github.com/torvalds/linux/blob/master/include/uapi/linux/bpf.h
# remove bpf_map_def
bpf_helpers.h: https://github.com/torvalds/linux/blob/master/tools/testing/selftests/bpf/bpf_helpers.h
bpf_map.h: https://github.com/iovisor/gobpf/blob/master/elf/include/bpf_map.h
```