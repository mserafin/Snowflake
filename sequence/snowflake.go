package sequence

import (
    "sync"
)

const nodeBits int8 = 10
const sequenceBits int8 = 12

const maxNode int16 = -1 ^ (-1 << nodeBits)
const maxSequence int16 = -1 ^ (-1 << sequenceBits)

var (
    value PK
    mu    = &sync.Mutex{}
    tid   = &epoch{}
    sid   = &sequence{}
)

type Snowflake struct {
    Node Node
}

func (s *Snowflake) NextVal() PK {
    mu.Lock()

    if tid.next().isEqual() {
        if sid.next().get() == 0 {
            tid.waitNotEqual()
            sid.clear()
        }
    } else {
        sid.clear()
    }

    r := tid.get() << (nodeBits + sequenceBits)
    r |= int64(s.Node.get()) << sequenceBits
    r |= int64(sid.get())

    value = PK(r)
    tid.toCopy()

    mu.Unlock()
    return value
}

func (Snowflake) CurrVal() PK {
    if value <= 0 {
        return PK(-1)
    }
    return value
}
