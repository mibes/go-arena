// Assembly to get into package runtime without using exported symbols.

// +build amd64 amd64p32 arm arm64 386 ppc64 ppc64le

#include "textflag.h"

#ifdef GOARCH_arm
#define JMP B
#endif
#ifdef GOARCH_ppc64
#define JMP BR
#endif
#ifdef GOARCH_ppc64le
#define JMP BR
#endif

// func memclrNoHeapPointers(ptr unsafe.Pointer, n uintptr)
TEXT ·memclrNoHeapPointers(SB), NOSPLIT, $0-16
    JMP runtime·memclrNoHeapPointers(SB)
