package main

import "C"
import (
	"encoding/json"
	"fmt"
	"github.com/go-vgo/robotgo"
	"io/ioutil"
	"log"
	"os"
	"time"
)

/*
	開始行動 0 138 204 => 1596, 958
	出動 76 10 0 => 1598, 735
	結束 255 150 2 => 1477, 789
*/

var screenWidth int
var screenHeight int
var pixelAndPos map[string][][]int

func initMap() {
	// read from settings
	file, err := os.Open("conf.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	config := Config{}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("config: %v\n", config)

	pixelAndPos = make(map[string][][]int)
	pixelAndPos["commence"] = [][]int { config.Commence[0], config.Commence[1] }
	pixelAndPos["operationStart"] = [][]int { config.OperationStart[0], config.OperationStart[1] }
	pixelAndPos["operationEnd"] = [][]int { config.OperationEnd[0], config.OperationEnd[1] }
	for key, elem := range pixelAndPos {
		fmt.Printf("%v: %v\n", key, elem)
	}
}

func clickOnElement(r, g, b, x, y int) bool {
	hex := robotgo.RgbToHex(uint8(r), uint8(g), uint8(b))
	findX, findY := robotgo.FindColorCS(robotgo.CHex(hex), 0, 0, screenWidth, screenHeight)
	if findX != -1 && findY != -1 {
		fmt.Printf("pixel found: %v, %v, %v, clicking on cords: %v, %v\n", r, g, b, x, y)
		robotgo.MoveClick(x, y, "left", false)
		time.Sleep(1 * time.Second)
		return true
	}
	time.Sleep(1 * time.Second)
	return false
}

type Config struct {
	Commence [][]int `json:"commence"`
	OperationStart [][]int `json:"operationStart"`
	OperationEnd [][]int `json:"operationEnd"`
}

func runTrain() {
	for i := 5; i > 0; i-- {
		fmt.Printf("training commencing in %v seconds...\n", i)
		time.Sleep(1 * time.Second)
	}
	initMap()
	screenWidth, screenHeight = robotgo.GetScreenSize()

	for {
		bitmap := robotgo.CaptureScreen(0, 0, screenWidth, screenHeight)
		defer robotgo.FreeBitmap(bitmap)

		for _, elem := range pixelAndPos {
			hasClicked := clickOnElement(elem[0][0], elem[0][1], elem[0][2], elem[1][0], elem[1][1])
			if hasClicked {
				break
			}
		}
		fmt.Println("element not found, wait for next loop")
		time.Sleep(2 * time.Second)
	}
}

func main() {
	//for {
	//	test.GetMousePositon()
	//}

	runTrain()
}
