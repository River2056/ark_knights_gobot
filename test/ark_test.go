package test

import "testing"

func TestGetMousePosition(t *testing.T) {
	for {
		GetMousePosition()
	}
}

func TestCaptureScreen(t *testing.T) {
	CaptureScreen()
}

func TestWriteToFile(t *testing.T) {
	WriteToFile()
}

func TestRgbToHex(t *testing.T) {
	RgbToHex(0, 138, 204)
	RgbToHex(193, 70, 0)
	RgbToHex(255, 150, 2)
}

func TestReadFromImage(t *testing.T) {
	ReadFromImage()
}