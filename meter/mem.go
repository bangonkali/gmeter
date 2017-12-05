package meter

// #include <unistd.h>
import "C"

func Mem() uint64 {
	return uint64(C.sysconf(C._SC_PHYS_PAGES) * C.sysconf(C._SC_PAGE_SIZE))
}