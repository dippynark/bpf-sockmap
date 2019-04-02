LINUX_HEADERS=$(shell uname -r | awk '{print "/usr/src/linux-headers-"$$1}')
DEST_DIR?=/dist

BPF = bpf_sockmap.o

CLANG ?= clang
CLANG_FLAGS := -D__KERNEL__ -D__ASM_SYSREG_H \
	-D__NR_CPUS__=$(shell nproc) \
	-Wno-unused-value \
	-Wno-pointer-sign \
	-Wno-compare-distinct-pointer-types \
	-Wunused \
	-Wall \
	-Werror \
	-O2 -emit-llvm \
	--sysroot ${CURDIR} \
	$(foreach path,$(LINUX_HEADERS),-I $(path)/include -I $(path)/arch/x86/include -I $(path)/arch/x86/include/generated -I $(path)/include/generated/uapi -I $(path)/arch/x86/include/uapi -I $(path)/include/uapi)

LLC ?= llc
LLC_FLAGS := -march=bpf -filetype=obj

GO_BINDATA := go-bindata
GO_BINDATA_FLAGS := -pkg sockmap -prefix "${DEST_DIR}/" -modtime 1

bpf: $(BPF)

%.o: %.c
	@mkdir -p "$(DEST_DIR)"
	${CLANG} ${CLANG_FLAGS} -c $< -o - | ${LLC} ${LLC_FLAGS} -o "${DEST_DIR}/$@"
	${GO_BINDATA} ${GO_BINDATA_FLAGS} -o "${DEST_DIR}/$(@:.o=.go)" "${DEST_DIR}/$@"

