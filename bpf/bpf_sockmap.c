// disable randomised task struct (Linux 4.13)
#define randomized_struct_fields_start  struct {
#define randomized_struct_fields_end    };

#include <uapi/linux/bpf.h>
#include <linux/version.h>

#include "bpf_helpers.h"
#include "bpf_map.h"

struct bpf_map_def SEC("maps/sockmap") sock_map = {
	.type = BPF_MAP_TYPE_SOCKMAP,
	.key_size = sizeof(int),
	.value_size = sizeof(unsigned int),
	.max_entries = 2,
  .pinning = 0,
	.namespace = "",
};

#define DEBUG 1
#ifndef DEBUG
/* Only use this for debug output. Notice output from bpf_trace_printk()
 * end-up in /sys/kernel/debug/tracing/trace_pipe
 */
#define bpf_debug(fmt, ...)                                                    \
	({                                                                     \
		char ____fmt[] = fmt;                                          \
		bpf_trace_printk(____fmt, sizeof(____fmt), ##__VA_ARGS__);     \
	})
#else
#define bpf_debug(fmt, ...)                                                    \
	{                                                                      \
	}                                                                      \
	while (0)
#endif

SEC("sk/skb/parser/sockmap")
int _prog_parser(struct __sk_buff *skb)
{
	return skb->len;
}

SEC("sk/skb/verdict/sockmap")
int _prog_verdict(struct __sk_buff *skb)
{
  uint32_t idx = 0;
	return bpf_sk_redirect_map(skb, &sock_map, idx, 0);
}

char _license[] SEC("license") = "GPL";
u32 _version SEC("version") = LINUX_VERSION_CODE;
