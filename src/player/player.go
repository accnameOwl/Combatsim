package player

// Player struct
type Player struct {
	Level     uint16
	Exp       Experience
	Stamina   Statistic
	Mana      Statistic
	Strength  Statistic
	Intellect Statistic
	Agility   Statistic

	Spells []Spell
}

// Levelup values for different stats
const (
	LevelUpIncreaseStamina   = 200
	LevelUpIncreaseMana      = 10
	LevelUpIncreaseStrength  = 10
	LevelUpIncreaseIntellect = 10
	LevelUpIncreaseAgility   = 10
)

// IncrementLevelBy func
func (p *Player) IncrementLevelBy(i uint16) { p.Level += i }

// OnGainExperience ...
func (p *Player) OnGainExperience(i float32) {
	levelsIncreased := uint16(p.Exp.IncreaseBy(i))
	if levelsIncreased > 0 {
		p.IncrementLevelBy(levelsIncreased)
		// Statistic Value
		for levelsIncreased > 0 {
			p.IncreaseStatisticStd(LevelUpIncreaseStamina, LevelUpIncreaseMana, LevelUpIncreaseStrength, LevelUpIncreaseIntellect, LevelUpIncreaseAgility)
			levelsIncreased--
		}
	}
}

// IncreaseStatisticStd ...
func (p *Player) IncreaseStatisticStd(sta uint16, mana uint16, str uint16, int uint16, agil uint16) {
	p.Stamina.IncreaseValueBy(LevelUpIncreaseStamina)
	p.Mana.IncreaseValueBy(LevelUpIncreaseMana)
	p.Strength.IncreaseValueBy(LevelUpIncreaseStrength)
	p.Intellect.IncreaseValueBy(LevelUpIncreaseIntellect)
	p.Agility.IncreaseValueBy(LevelUpIncreaseAgility)
	// Statistic Limit
	p.Stamina.IncreaseLimitBy(LevelUpIncreaseStamina)
	p.Mana.IncreaseLimitBy(LevelUpIncreaseMana)
	p.Strength.IncreaseLimitBy(LevelUpIncreaseStrength)
	p.Intellect.IncreaseLimitBy(LevelUpIncreaseIntellect)
	p.Agility.IncreaseLimitBy(LevelUpIncreaseAgility)
}
