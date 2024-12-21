package main

import (
	"fmt"
	"strings"
	"flag"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	gloss "github.com/charmbracelet/lipgloss"
)
type AppState int

const (
	AppScreen AppState = iota
	ErrorScreen
)

const(
	minWidth=100
	minHeight=40
)

type MainModel struct {
	currentTab int
	width      int
	height     int
	tab1       Tab1Model
	tab2       Tab2Model
	styles     Styles
	currentScreen AppState
}

var tabNames = []string{"Watch Anime", "About"}


func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width-7
		m.height = msg.Height
		if m.width < minWidth || m.height < minHeight{
			m.currentScreen = ErrorScreen
		}else{
			m.currentScreen = AppScreen
		}
	case tea.KeyMsg:
		switch m.currentScreen{
			case AppScreen:
				switch msg.String() {
					case "tab":
						m.currentTab = (m.currentTab + 1) % len(tabNames)
					case "ctrl+tab":
						m.currentTab = (m.currentTab - 1 + len(tabNames)) % len(tabNames)
					case "esc":
						return m, tea.Quit
				}
				if m.currentTab == 0 {
					switch{
						case key.Matches(msg, keys.Enter):
							if m.tab1.focus == inputFocus {
								searchTerm := m.tab1.inputM.Value()
								if searchTerm == ""{
									return m, nil
								}
								m.tab1.loading = true
								m.tab1.focus = tableFocus
								m.tab1.table.Focus()
								m.tab1.styles.inputBorder = m.tab1.styles.inputBorder.BorderForeground(gloss.Color(m.tab1.styles.inactiveColor))
								m.tab1.styles.list1Border = m.tab1.styles.list1Border.BorderForeground(gloss.Color(m.tab1.styles.inactiveColor))
								m.tab1.styles.list2Border = m.tab1.styles.list2Border.BorderForeground(gloss.Color(m.tab1.styles.inactiveColor))
								m.tab1.styles.tableBorder = m.tab1.styles.tableBorder.BorderForeground(gloss.Color(m.tab1.styles.activeColor))
								return m, tea.Batch(m.tab1.fetchAnimeData(searchTerm), m.tab1.spinner.Tick)
							} 					
					}	
			
					var cmd tea.Cmd
					m.tab1, cmd = m.tab1.Update(msg)
					return m, cmd
				}else if m.currentTab == 1{
					var cmd tea.Cmd
					m.tab2, cmd = m.tab2.Update(msg)
					return m, cmd
				}
		}
	}
	return m, nil
}