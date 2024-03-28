package model

// A LARP.
type Game struct {
	Name         string
	Description  string
	Organization Organization
	System       System
}

func MakeGame(name string, description string, organization Organization, system System) Game {
	return Game{
		Name:         name,
		Description:  description,
		Organization: organization,
		System:       system,
	}
}

type Organization struct {
	Name        string
	Description string
}

func MakeOrganization(name string, description string) Organization {
	return Organization{
		Name:        name,
		Description: description,
	}
}

type System struct {
	Name        string
	Description string
}

func MakeSystem(name string, description string) System {
	return System{
		Name:        name,
		Description: description,
	}
}
