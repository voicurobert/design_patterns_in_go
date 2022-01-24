package chain_of_responsability

import "fmt"

type Creature struct {
	Name            string
	Attack, Defence int
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.Attack, c.Defence)
}

func NewCreature(name string, attack int, defence int) *Creature {
	return &Creature{name, attack, defence}
}

type Modifier interface {
	Add(m Modifier)
	Handle()
}

type CreatureModifier struct {
	creature *Creature
	next     Modifier
}

func (c *CreatureModifier) Add(m Modifier) {
	if c.next != nil {
		c.next.Add(m)
	} else {
		c.next = m
	}

}

func (c *CreatureModifier) Handle() {
	if c.next != nil {
		c.next.Handle()
	}

}

func NewCreatureModifier(creature *Creature) *CreatureModifier {
	return &CreatureModifier{creature, nil}
}

type DoubleAttackModifier struct {
	CreatureModifier
}

func NewDoubleAttackModifier(c *Creature) *DoubleAttackModifier {
	return &DoubleAttackModifier{CreatureModifier{c, nil}}
}

func (d *DoubleAttackModifier) Handle() {
	fmt.Println("Doubling", d.creature.Name, "\b's attack")
	d.creature.Attack *= 2
	d.CreatureModifier.Handle()
}

type IncreaseDefenceModifier struct {
	CreatureModifier
}

func NewIncreaseDefenceModifier(c *Creature) *IncreaseDefenceModifier {
	return &IncreaseDefenceModifier{CreatureModifier{c, nil}}
}

func (d *IncreaseDefenceModifier) Handle() {
	if d.creature.Attack <= 2 {
		fmt.Println("Increasing", d.creature.Name, "\b's defence")
		d.creature.Defence++
	}

	d.CreatureModifier.Handle()
}

type NoBonusesModifier struct {
	CreatureModifier
}

func NewNoBonusesModifier(c *Creature) *NoBonusesModifier {
	return &NoBonusesModifier{CreatureModifier{c, nil}}
}

func (n *NoBonusesModifier) Handle() {
	// empty
}

func MethodChainExample() {
	goblin := NewCreature("Goblin", 1, 1)
	fmt.Println(goblin.String())

	root := NewCreatureModifier(goblin)

	root.Add(NewNoBonusesModifier(goblin))

	root.Add(NewDoubleAttackModifier(goblin))
	root.Add(NewIncreaseDefenceModifier(goblin))
	root.Add(NewDoubleAttackModifier(goblin))
	root.Handle()
	fmt.Println(goblin.String())
}
