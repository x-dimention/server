package manage

type manage struct {
	registerSelfField
}

func (m *manage) new() {
	m.loadHamster()
}

func (m *manage) Run() {
	detection()
	m.StartServe()
}
