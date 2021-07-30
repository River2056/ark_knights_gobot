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

func clickOnElement(r, g, b, x, y int, c chan string) {
	for {
		bitmap := robotgo.CaptureScreen(0, 0, screenWidth, screenHeight)
		defer robotgo.FreeBitmap(bitmap)

		hex := robotgo.RgbToHex(uint8(r), uint8(g), uint8(b))
		findX, findY := robotgo.FindColorCS(robotgo.CHex(hex), 0, 0, screenWidth, screenHeight)
		if findX != -1 && findY != -1 {
			c <- fmt.Sprintf("pixel found: %v, %v, %v, clicking on cords: %v, %v\n", r, g, b, x, y)
			//fmt.Printf("pixel found: %v, %v, %v, clicking on cords: %v, %v\n", r, g, b, x, y)
			robotgo.MoveClick(x, y, "left", false)
		} else {
			c <- fmt.Sprintf("element not found: %v, %v, %v\n", r, g, b)
		}
		time.Sleep(3 * time.Second)
	}
}

type Config struct {
	Commence [][]int `json:"commence"`
	OperationStart [][]int `json:"operationStart"`
	OperationEnd [][]int `json:"operationEnd"`
}

func main() {
	initMap()
	screenWidth, screenHeight = robotgo.GetScreenSize()

	c := make(chan string)
	for _, val := range pixelAndPos {
		r, g, b, x, y := val[0][0], val[0][1], val[0][2], val[1][0], val[1][1]
		go clickOnElement(r, g, b, x, y, c)
	}

	for i := 5; i >= 1; i-- {
		fmt.Printf("commencing training in %v seconds...\n", i)
		time.Sleep(time.Second * 1)
	}

	for {
		select {
		case msg := <- c:
			fmt.Print(msg)
		}
	}

}
