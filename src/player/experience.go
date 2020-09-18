package player

const (
	experienceRequiredRatio = 1.05
)

// Experience struct
type Experience struct {
	current  float32
	required float32
	total    float32
}

// Current func
func (e *Experience) Current() float32 { return e.current }

// Required func
func (e *Experience) Required() float32 { return e.required }

// Total func
func (e *Experience) Total() float32 { return e.total }

// SetCurrent func
func (e *Experience) SetCurrent(i float32) { e.current += i }

// SetRequired func
func (e *Experience) SetRequired(i float32) { e.required += i }

// SetTotal func
func (e *Experience) SetTotal(i float32) { e.total += i }

// ResetCurrent func
//
// order of procedures: IncreaseTotal() -> ResetValue()
func (e *Experience) ResetCurrent() {
	e.current = 0
}

// IncreaseCurrentBy func
func (e *Experience) IncreaseCurrentBy(i float32) {
	if i >= 0 {
		e.current += i
	}
}

// IncreaseRequiredBy func
func (e *Experience) IncreaseRequiredBy(i float32) {
	if i >= 0 {
		e.required += i
	}
}

// IncreaseTotalBy func
//
// order of procedures: IncreaseTotal() -> ResetValue()
func (e *Experience) IncreaseTotalBy(i float32) {
	if e.current >= 0 {
		e.total += i
	}
}

// IncreaseBy ...
//
// Returns how many times current xp exceeded required xp.
//
// in short: returns levels_increased uint8
func (e *Experience) IncreaseBy(i float32) (levelups uint8) {
	var leftover float32
	e.IncreaseCurrentBy(i)
	for e.current >= e.required {
		levelups++
		leftover = e.current - e.required
		e.SetCurrent(leftover)
		e.IncreaseRequiredBy(e.required * experienceRequiredRatio)
	}
	e.IncreaseTotalBy(i)

	return levelups
}
