package sequence

type PK int64

type Sequence interface {
    NextVal() PK
    CurrVal() PK
}

type sequence struct {
    id int16
}

func (s *sequence) get() int16 {
    return s.id
}

func (s *sequence) next() *sequence {
    s.id = (s.id + 1) & maxSequence
    return s
}

func (s *sequence) clear() *sequence {
    s.id = 0
    return s
}
