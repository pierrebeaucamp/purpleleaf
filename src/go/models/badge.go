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
	case "ag5":
		return Badge{
			Image:   "/assets/img/" + name + ".png",
			Tooltip: "5% agrar",
			Type:    "agrar",
		}
	case "edu100":
		return Badge{
			Image:   "/assets/img/" + name + ".png",
			Tooltip: "100% education",
			Type:    "education",
		}
	case "trade50":
		return Badge{
			Image:   "/assets/img/" + name + ".png",
			Tooltip: "50% trade",
			Type:    "trade",
		}
	}
	return Badge{}
}
