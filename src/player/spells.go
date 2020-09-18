package player

// Spell ...
type Spell struct {
	name             string
	id               uint8
	damageMultiplier float32
	School
}

// School ...
//
// school of spelldamage
type School int

// enum ... type School
const (
	Physical School = iota
	Void
	Holy
	Frost
	Fire
	Nature
	Chaos
)

// SpellList ...
var SpellList map[int]Spell

// InitSpells ...
func InitSpells() {
	SpellList[1] = Spell{"Melee", 1, 1, Physical}
	SpellList[2] = Spell{"Shatter", 2, 1.2, Void}
	SpellList[3] = Spell{"Smite", 3, 1.2, Holy}
	SpellList[4] = Spell{"Frostbolt", 4, 1.2, Frost}
	SpellList[5] = Spell{"Fireball", 5, 1.2, Fire}
	SpellList[6] = Spell{"Lightningbolt", 6, 1.2, Nature}
	SpellList[7] = Spell{"Chaos bolt", 7, 1.2, Chaos}
}

// GiveSpellByID ...
func (p *Player) GiveSpellByID(arg int) {
	if len(SpellList) >= arg {
		p.Spells = append(p.Spells, SpellList[arg])
	}
}

// GiveSpellsByID ...
func (p *Player) GiveSpellsByID(arg []int) {
	var spellsToGive []Spell
	for _, v := range arg {
		spellsToGive = append(spellsToGive, SpellList[v])
	}
	p.Spells = append(p.Spells, spellsToGive...)
}
