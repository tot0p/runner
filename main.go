package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tot0p/env"
	"runner/vutlr"
)

type Model struct {
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		default:
			return m, nil
		}
	}
	return m, nil
}

func (m Model) View() string {
	return "Hello, Bubble Tea!"
}

func init() {
	err := env.Load()
	if err != nil {
		panic(err)
	}

}

func main() {
	api := vutlr.New()
	api.SetAPIKey(env.Get("API_KEY"))

	p := tea.NewProgram(Model{})
	if _, err := p.Run(); err != nil {
		panic(err)
	}

	/*
		js, err := os.ReadFile("vm.json")
		if err != nil {
			panic(err)
		}
		i, err := api.CreateInstance(string(js))
		if err != nil {
			panic(err)
		}
		fmt.Println(i.ID)

		lst := api.ListInstances()
		fmt.Println(lst.Meta.Total)

	*/

	/*
		id := lst.Instances[0].ID

		fmt.Println(id)

		i, err := api.GetInstance(id)
		if err != nil {
			panic(err)
		}

		fmt.Println(i.ID)
		fmt.Println(i.Os)

	*/
	/*
		for i.ServerStatus != "ok" {
			i, err = api.GetInstance(i.ID)
			if err != nil {
				panic(err)
			}
			fmt.Println(i.ServerStatus)
		}

		if i.Label == "ApiRunner" {
			fmt.Println("Instance created")
			err := api.DeleteInstance(i.ID)
			if err != nil {
				panic(err)
			}
			fmt.Println("Instance deleted")
		}

	*/
}
