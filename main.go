package main

import "C"
import (
	"encoding/json"
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
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
var typeMap map[string]int

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

	commence := "commence"
	operationStart := "operationStart"
	operationEnd := "operationEnd"
	pixelAndPos = make(map[string][][]int)
	pixelAndPos[commence] = [][]int { config.Commence[0], config.Commence[1] }
	pixelAndPos[operationStart] = [][]int { config.OperationStart[0], config.OperationStart[1] }
	pixelAndPos[operationEnd] = [][]int { config.OperationEnd[0], config.OperationEnd[1] }
	for key, elem := range pixelAndPos {
		fmt.Printf("%v: %v\n", key, elem)
	}
	typeMap = make(map[string]int)
	typeMap[commence] = 0
	typeMap[operationStart] = 1
	typeMap[operationEnd] = 2
}

func checkIfCanExecute(previous, current int) bool {
	if
		(current == 0 && previous == 2) ||
		(current == 1 && previous == 0) ||
		(current == 2 && previous == 1) {
		return true
	}
	return false
}

func clickOnElement(r, g, b, x, y int, receive chan string, send chan int, routineType string, sec time.Duration) {
	selfType := typeMap[routineType]
	canExecute := false
	for {
		bitmap := robotgo.CaptureScreen(0, 0, screenWidth, screenHeight)
		defer robotgo.FreeBitmap(bitmap)

		hex := robotgo.RgbToHex(uint8(r), uint8(g), uint8(b))
		findX, findY := robotgo.FindColorCS(robotgo.CHex(hex), 0, 0, screenWidth, screenHeight)
		previousAction := <- send
		canExecute = checkIfCanExecute(previousAction, selfType)
		fmt.Printf("selfType: %v, previousAction: %v, canExecute: %v\n", selfType, previousAction, canExecute)
		if findX != -1 && findY != -1 && canExecute {
			//send <- selfType
			receive <- fmt.Sprintf("pixel found: %v, %v, %v, clicking on cords: %v, %v\n", r, g, b, x, y)
			robotgo.MoveClick(x, y, "left", false)
		} else {
			//nextAction := selfType + 1
			//send <- nextAction % 3
			receive <- fmt.Sprintf("element not found: %v, %v, %v\n", r, g, b)
		}
		//if canExecute {
		//	sec += time.Second * 1
		//}
		time.Sleep(sec)
	}
}

type Config struct {
	Commence [][]int `json:"commence"`
	OperationStart [][]int `json:"operationStart"`
	OperationEnd [][]int `json:"operationEnd"`
}

func main() {
	running := true
	initMap()
	screenWidth, screenHeight = robotgo.GetScreenSize()
	seq := 2

	receive := make(chan string)
	send := make(chan int, 3)
	//send <- 2
	count := 3

	// add event hook for easy program termination
	robotgo.EventHook(hook.KeyDown, []string { "enter" }, func(e hook.Event) {
		fmt.Println("ending training program...")
		running = false
		robotgo.EventEnd()
	})

	eventStart := robotgo.EventStart()
	robotgo.EventProcess(eventStart)

	for i := 5; i >= 1; i-- {
		fmt.Printf("commencing training in %v seconds...\n", i)
		time.Sleep(time.Second * 1)
	}

	for key, val := range pixelAndPos {
		r, g, b, x, y := val[0][0], val[0][1], val[0][2], val[1][0], val[1][1]
		t := time.Second * time.Duration(count)
		go clickOnElement(r, g, b, x, y, receive, send, key, t)
		fmt.Printf("created go routine with time.Sleep => %v second\n", count)
		//count++
	}

	for running {
		send <- seq % 3
		seq++
		time.Sleep(time.Second * 1)
		select {
		case msg := <-receive:
			fmt.Print(msg)
		}
	}
}