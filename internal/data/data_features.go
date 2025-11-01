package data


// Hardcoded features to add to database

func GetHardcodedFeatures() []AddFeatureParams {
	var hardcodedFeatures []AddFeatureParams

	// Visual
	subtitles := CreateHardcodedFeature("Subtitles", "Display text for dialogue and important sound effects", "Visual")
	colorblind := CreateHardcodedFeature("Colorblind Modes", "Provide filters or distinct color palettes to help players with color vision deficiencies", "Visual")
	textformat := CreateHardcodedFeature("Text Formatting", "Allow players to adjust size or style of text on screen", "Visual")
	visualcues := CreateHardcodedFeature("Visual Cues", "Provide visual indicators for in-game sounds", "Visual")
	highcontrast := CreateHardcodedFeature("High Contrast", "Enhance visibility by increasing the contrast between foreground and background elements", "Visual")
	screenmagnifier := CreateHardcodedFeature("Screen Magnifier", "Allow players to zoom in or out on screen elements", "Visual")

	// Audio
	volumecontrols := CreateHardcodedFeature("Volume Controls", "Allow players to adjust the volume for independent sound elements such as dialogue, sound effects, or music", "Audio")
	audiocues := CreateHardcodedFeature("Audio Cues", "Provide sound effects for important gameplay events", "Audio")
	texttospeech := CreateHardcodedFeature("Text to Speech", "Read on-screen text aloud", "Audio")

	// Motor
	remapcontrols := CreateHardcodedFeature("Remappable Controls", "Allow players to change game inputs such as buttons or keys, or to toggle actions on or off instead of holding a button down", "Motor")
	difficulty := CreateHardcodedFeature("Adjustable Difficulty", "Offer a range of difficulty settings, including assist modes that might provide more resources or slow down gameplay", "Motor")
	alternativeinput := CreateHardcodedFeature("Alternative Input", "Support for additional game controllers or input schemes", "Motor")
	gamespeed := CreateHardcodedFeature("Game Speed", "Provide an option to change the speed at which the game runs", "Motor")

	// Cognitive
	clearmenu := CreateHardcodedFeature("Clear Interfaces", "Menus and in-game information is clear and easy to navigate", "Cognitive")
	motionblur := CreateHardcodedFeature("Motion Blur and Screen Shake", "Options to disable or reduce visual effects that can cause motion sickness", "Cognitive")

	hardcodedFeatures = append(hardcodedFeatures, subtitles, colorblind, textformat, visualcues, highcontrast, screenmagnifier, volumecontrols, audiocues, texttospeech, remapcontrols, difficulty, alternativeinput, gamespeed, clearmenu, motionblur)
	return hardcodedFeatures
}

func CreateHardcodedFeature(name, description, category string) AddFeatureParams {
	return AddFeatureParams{Name: name, Description: ToPgtypeText(description), Category: ToPgtypeText(category)}
}