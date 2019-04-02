package sockmap

import (
	"bytes"
	"errors"
	"fmt"
	"unsafe"

	"github.com/dippynark/gobpf/elf"
)

/*
#include "include/bpf.h"
*/
import "C"

type Sockmap struct {
	module *elf.Module
}

func New() (*Sockmap, error) {

	buf, err := Asset("bpf_sockmap.o")
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(buf)

	m := elf.NewModuleFromReader(reader)
	if m == nil {
		return nil, errors.New("failed to create new module")
	}

	err = m.Load(map[string]elf.SectionParams{})
	if err != nil {
		return nil, fmt.Errorf("failed to load BPF programs and maps: %s", err)
	}

	err = m.AttachParserVerdictPrograms()
	if err != nil {
		return nil, fmt.Errorf("failed to attach sockmap programs: %s", err)
	}

	sockmap := Sockmap{module: m}

	return &sockmap, nil
}

func (s *Sockmap) Close() error {
	return s.module.Close()
}

func (s *Sockmap) UpdateSocketDescriptor(socketDescriptor uintptr) error {
	sockmap := s.module.Map("sockmap")
	if sockmap == nil {
		return errors.New("failed to retrieve sockmap map")
	}
	key := 0
	value := socketDescriptor
	err := s.module.UpdateElement(sockmap, unsafe.Pointer(&key), unsafe.Pointer(&value), C.BPF_ANY)
	if err != nil {
		return fmt.Errorf("failed to update element: %s", err)
	}
	return nil
}

