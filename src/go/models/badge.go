package models

type Badge struct {
	Image   string
	Tooltip string
	Type    string
}

func GetBadge(name string) Badge {
	switch name {
	case "health20":
		return Badge{
			Image:   "/assets/img/" + name + ".png",
			Tooltip: "20% health",
			Type:    "Health",
		}
	}

	return Badge{}
}
