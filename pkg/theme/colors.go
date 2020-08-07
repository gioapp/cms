package theme

func NewColors() (c map[string]string) {
	c = map[string]string{
		"black":       "ff000000",
		"light-black": "ff222222",
		"blue":        "ff859ca9",
		"charcoal":    "ff34373f",
		"dark-blue":   "ff0b3a53",
		"dark-gray":   "ff656565",
		"silver":      "ff999999",
		"green":       "ff0cb892",
		"lite-blue":   "fff0f6fa",
		"orange":      "fff39021",
		"purple":      "ffcf30cf",
		"red":         "ffcf3030",
		"white":       "ffffffff",
		"dark-white":  "fffbfbfb",
		"yellow":      "ffcfcf30",
	}

	c["Black"] = c["black"]
	c["White"] = c["white"]
	c["Silver"] = c["silver"]
	c["Charcoal"] = c["charcoal"]
	c["Primary"] = c["blue"]
	c["Secondary"] = c["green"]
	c["Success"] = c["green"]
	c["Danger"] = c["red"]
	c["Warning"] = c["yellow"]
	c["Info"] = c["lite-blue"]
	c["Check"] = c["orange"]
	c["Hint"] = c["light-gray"]
	c["InvText"] = c["light"]
	c["ButtonText"] = c["light"]
	c["ButtonBg"] = c["blue-lite-blue"]
	c["NavBg"] = c["dark-blue"]
	c["PanelText"] = c["light"]
	c["PanelBg"] = c["dark-white"]
	c["DocText"] = c["dark"]
	c["DocBg"] = c["light"]
	c["ButtonTextDim"] = c["dark-white"]
	c["ButtonBgDim"] = "ff30809a"
	c["PanelTextDim"] = c["00000000"]
	c["PanelBgDim"] = c["dark-white"]
	c["DocTextDim"] = c["00000000"]
	c["DocBgDim"] = c["00000000"]
	c["Transparent"] = c["00000000"]
	c["Fatal"] = "ff880000"
	return c
}
