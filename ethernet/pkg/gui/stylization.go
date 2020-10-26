package gui

func ConfigureBackgroundColor(data []byte, color string) []byte {
	var newBytes []byte
	newBytes = append(newBytes, []byte("<font style=\"background: "+color+"\">")...)
	newBytes = append(newBytes, data...)
	newBytes = append(newBytes, []byte("</font>")...)
	return newBytes
}
