package test

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"image/png"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

func GetMousePosition() {
	x, y := robotgo.GetMousePos()
	fmt.Printf("x: %v, y: %v\n", x, y)
	time.Sleep(2 * time.Second)
}

func CaptureScreen() {
	w, h := robotgo.GetScreenSize()
	bitmap := robotgo.CaptureScreen(0, 0, w, h)
	defer robotgo.FreeBitmap(bitmap)

	robotgo.SaveBitmap(bitmap, "test.png")
}

func WriteToFile() {
	err := ioutil.WriteFile("output.txt", []byte("colors\n"), 0666)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//for y := 0; y < h; y++ {
	//	for x := 0; x < w; x++ {
	//		pixel := robotgo.GetPxColor(x, y)
	//		fmt.Printf("color: %v\n", pixel)
	//		if _, err := file.WriteString(fmt.Sprintf("pixel: %v", pixel)); err != nil {
	//			log.Fatal(err)
	//		}
	//	}
	//}
}

func RgbToHex(r, g, b int) {
	fmt.Printf("r: %v\n", r)
	fmt.Printf("g: %v\n", g)
	fmt.Printf("b: %v\n", b)

	fmt.Printf("unit8 r: %v\n", uint8(r))
	fmt.Printf("unit8 g: %v\n", uint8(g))
	fmt.Printf("unit8 b: %v\n", uint8(b))

	hex := robotgo.RgbToHex(uint8(r), uint8(g), uint8(b))
	fmt.Printf("hex: %v\n", hex)
}

func ReadFromImage() {
	f, err := os.Open("pic_01.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, err := png.Decode(f)
	fmt.Println(img)
}

func GetRandomSeconds() {
	t := time.Now().UnixNano()
	rs := rand.NewSource(t)
	r := rand.New(rs)

	fmt.Println(r.Int())
	fmt.Println(r.Intn(3))
}

func Duration() {
	count := 1
	t := time.Second * time.Duration(count)
	fmt.Println(t)
}

func GetKeyPressed() {
	running := true
	i := 0

	robotgo.EventHook(hook.KeyDown, []string { "enter" }, func(e hook.Event) {
		fmt.Println("Ending training program...")
		running = false
		robotgo.EventEnd()
	})

	eventStart := robotgo.EventStart()
	robotgo.EventProcess(eventStart)

	for running {
		fmt.Printf("count: %v\n", i)
		i++
		time.Sleep(time.Second * 1)
	}
}