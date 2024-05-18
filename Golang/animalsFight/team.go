package animalsfight

type Team struct {
	first  *Fighter
	second *Fighter
}

func NewTeam() *Team {
	return &Team{NewFighter("salam", 2), NewFighter("Chetori", 3)}
}

func (t *Team) Introduce() {
	t.first.Introduce()
	t.second.Introduce()
}

func (t *Team) GetPower() int {
	return t.first.power + t.second.power
}
