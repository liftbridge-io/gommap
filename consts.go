// ** This file is automatically generated from consts.c. **
package gommap

const (
    PROT_NONE     = 0x0
    PROT_READ     = 0x1
    PROT_WRITE    = 0x2
    PROT_EXEC     = 0x4
)

const (
    MAP_SHARED    = 0x1
    MAP_PRIVATE   = 0x2
    MAP_FIXED     = 0x10
    MAP_ANONYMOUS = 0x20
    MAP_GROWSDOWN = 0x100
    MAP_LOCKED    = 0x2000
    MAP_NONBLOCK  = 0x10000
    MAP_NORESERVE = 0x4000
    MAP_POPULATE  = 0x8000
)

const (
    MS_SYNC       = 0x4
    MS_ASYNC      = 0x1
    MS_INVALIDATE = 0x2
)

const (
    MADV_NORMAL   = 0x0
    MADV_RANDOM   = 0x1
    MADV_SEQUENTIAL = 0x2
    MADV_WILLNEED = 0x3
    MADV_DONTNEED = 0x4
    MADV_REMOVE   = 0x9
    MADV_DONTFORK = 0xa
    MADV_DOFORK   = 0xb
)
