package model

import "fmt"

type Model struct {
	status int

}

func (m Model) NewModel(status int) (Model, error) {

	if m.status < 0 {
		return m, fmt.Errorf("invalid value %d", m.status)
	}

	fmt.Println(m.status)
	return m, nil
}

func (m *Model) Run(s string) {
	m.status += 10
	fmt.Println(s, m.status)
}

func (m *Model) Jump(s string) {
	m.status -= 5
	fmt.Println(s,m.status)
}

func (m *Model) Sleep(s string) {
	m.status += 100
	fmt.Println(s,m.status)
}

func(m *Model) Walk(s string) {
	m.status -= 1
	fmt.Println(s,m.status)
}

func (m *Model) Fly(s string) {
	m.status -= 30
	fmt.Println(s,m.status)
}