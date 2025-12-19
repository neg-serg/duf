package main

import (
	"fmt"

	"github.com/muesli/termenv"
)

// Theme defines a color theme used for printing tables.
type Theme struct {
	colorRed     termenv.Color
	colorYellow  termenv.Color
	colorGreen   termenv.Color
	colorBlue    termenv.Color
	colorGray    termenv.Color
	colorMagenta termenv.Color
	colorCyan    termenv.Color
	colorWhite   termenv.Color

	colorBgRed    termenv.Color
	colorBgYellow termenv.Color
	colorBgGreen  termenv.Color
}

func defaultThemeName() string {
	if !termenv.HasDarkBackground() {
		return "light"
	}
	return "dark"
}

func loadTheme(theme string) (Theme, error) {
	themes := make(map[string]Theme)

	themes["dark"] = Theme{
		colorRed:      env.Color("#E88388"),
		colorYellow:   env.Color("#DBAB79"),
		colorGreen:    env.Color("#A8CC8C"),
		colorBlue:     env.Color("#71BEF2"),
		colorGray:     env.Color("#B9BFCA"),
		colorMagenta:  env.Color("#D290E4"),
		colorCyan:     env.Color("#66C2CD"),
		colorWhite:    env.Color("#FFFFFF"),
		colorBgRed:    env.Color("#2d1b1b"),
		colorBgYellow: env.Color("#2d2d1b"),
		colorBgGreen:  env.Color("#1b2d1b"),
	}

	themes["light"] = Theme{
		colorRed:      env.Color("#D70000"),
		colorYellow:   env.Color("#FFAF00"),
		colorGreen:    env.Color("#005F00"),
		colorBlue:     env.Color("#000087"),
		colorGray:     env.Color("#303030"),
		colorMagenta:  env.Color("#AF00FF"),
		colorCyan:     env.Color("#0087FF"),
		colorWhite:    env.Color("#000000"),
		colorBgRed:    env.Color("#ffdede"),
		colorBgYellow: env.Color("#fff4d0"),
		colorBgGreen:  env.Color("#e6ffe6"),
	}

	themes["ansi"] = Theme{
		colorRed:      env.Color("9"),
		colorYellow:   env.Color("11"),
		colorGreen:    env.Color("10"),
		colorBlue:     env.Color("12"),
		colorGray:     env.Color("7"),
		colorMagenta:  env.Color("13"),
		colorCyan:     env.Color("8"),
		colorWhite:    env.Color("15"),
		colorBgRed:    env.Color("1"),
		colorBgYellow: env.Color("3"),
		colorBgGreen:  env.Color("2"),
	}

	// Cold blue-tinted theme based on custom kitty color scheme
	themes["neg"] = Theme{
		colorRed:      env.Color("#8A2F58"), // muted magenta-red
		colorYellow:   env.Color("#914E89"), // purple-ish yellow
		colorGreen:    env.Color("#287373"), // teal green
		colorBlue:     env.Color("#005faf"), // kitty color25 - bright blue
		colorGray:     env.Color("#6C7E96"), // blue-gray
		colorMagenta:  env.Color("#5E468C"), // muted purple
		colorCyan:     env.Color("#31658C"), // cold cyan
		colorWhite:    env.Color("#95a7bc"), // kitty color250
		colorBgRed:    env.Color("#0d1824"), // dark blue bg
		colorBgYellow: env.Color("#10233a"), // darker blue bg
		colorBgGreen:  env.Color("#071526"), // darkest blue bg
	}

	if _, ok := themes[theme]; !ok {
		return Theme{}, fmt.Errorf("unknown theme: %s", theme)
	}

	return themes[theme], nil
}
