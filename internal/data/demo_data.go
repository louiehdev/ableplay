package data

// Hardcoded games to add to database

func GetHardcodedGames() []AddGameParams {
	return []AddGameParams{
		CreateHardcodedGame("Baldur's Game 3", "Larian Studios", "2023", "A story-rich, party-based RPG set in the Dungeons & Dragons universe where player choices shape the narrative of fellowship and betrayal.", []string{"PC", "Xbox", "Playstation"}),
		CreateHardcodedGame("Grand Theft Auto V", "Rockstar Games", "2015", "An open-world action-adventure game where players control three criminals as they commit a series of heists in the city of Los Santos and surrounding areas.", []string{"PC", "Xbox", "Playstation"}),
		CreateHardcodedGame("Mario Kart 8 Deluxe", "Nintendo", "2017", "A racing game where players compete in go-karts or motorcycles, featuring characters from the Mario series.", []string{"Nintendo Switch"}),
		CreateHardcodedGame("Stardew Valley", "ConcernedApe", "2016", "A farming simulation RPG where you inherit a run-down farm and can choose to restore it or focus on other activities like fishing, mining, and foraging.", []string{"Mobile", "PC", "Xbox", "Playstation", "Nintendo Switch"}),
		CreateHardcodedGame("Fortnite", "Epic Games", "2017", "Loot, build, explore, and fight in a game of 100 players competing to be the last one standing.", []string{"Mobile", "PC", "Xbox", "Playstation", "Nintendo Switch"}),
	}
}

// Hardcoded features to add to database

func GetHardcodedFeatures() []AddFeatureParams {
	return []AddFeatureParams{
		CreateHardcodedFeature("Subtitles", "Display text for dialogue and important sound effects", "Visual"),
		CreateHardcodedFeature("Colorblind Modes", "Provide filters or distinct color palettes to help players with color vision deficiencies", "Visual"),
		CreateHardcodedFeature("Text Formatting", "Allow players to adjust size or style of text on screen", "Visual"),
		CreateHardcodedFeature("Visual Cues", "Provide visual indicators for in-game sounds", "Visual"),
		CreateHardcodedFeature("High Contrast", "Enhance visibility by increasing the contrast between foreground and background elements", "Visual"),
		CreateHardcodedFeature("Screen Magnifier", "Allow players to zoom in or out on screen elements", "Visual"),
		CreateHardcodedFeature("Volume Controls", "Allow players to adjust the volume for independent sound elements such as dialogue, sound effects, or music", "Audio"),
		CreateHardcodedFeature("Audio Cues", "Provide sound effects for important gameplay events", "Audio"),
		CreateHardcodedFeature("Text to Speech", "Read on-screen text aloud", "Audio"),
		CreateHardcodedFeature("Remappable Controls", "Allow players to change game inputs such as buttons or keys, or to toggle actions on or off instead of holding a button down", "Motor"),
		CreateHardcodedFeature("Adjustable Difficulty", "Offer a range of difficulty settings, including assist modes that might provide more resources or slow down gameplay", "Motor"),
		CreateHardcodedFeature("Alternative Input", "Support for additional game controllers or input schemes", "Motor"),
		CreateHardcodedFeature("Game Speed", "Provide an option to change the speed at which the game runs", "Motor"),
		CreateHardcodedFeature("Clear Interfaces", "Menus and in-game information is clear and easy to navigate", "Cognitive"),
		CreateHardcodedFeature("Motion Blur and Screen Shake", "Options to disable or reduce visual effects that can cause motion sickness", "Cognitive"),
	}
}

func CreateHardcodedGame(title, developer, releaseYear, description string, platforms []string) AddGameParams {
	return AddGameParams{Title: title, Developer: ToPgtypeText(developer), ReleaseYear: ToPgtypeInt4(releaseYear), Description: ToPgtypeText(description), Platforms: platforms}
}

func CreateHardcodedFeature(name, description, category string) AddFeatureParams {
	return AddFeatureParams{Name: name, Description: ToPgtypeText(description), Category: ToPgtypeText(category)}
}
