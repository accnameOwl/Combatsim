package player

// Statistic struct
type Statistic struct {
	value uint16
	limit uint16
}

// Value func
func (s *Statistic) Value() uint16 { return s.value }

// Limit func
func (s *Statistic) Limit() uint16 { return s.limit }

// IncreaseValueBy ...
func (s *Statistic) IncreaseValueBy(i uint16) { s.value += i }

// IncreaseLimitBy ...
func (s *Statistic) IncreaseLimitBy(i uint16) { s.limit += i }
