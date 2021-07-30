package test

// previous implementation
//func runTrain() {
//	for i := 5; i > 0; i-- {
//		fmt.Printf("training commencing in %v seconds...\n", i)
//		time.Sleep(1 * time.Second)
//	}
//	initMap()
//	screenWidth, screenHeight = robotgo.GetScreenSize()
//
//	for {
//		bitmap := robotgo.CaptureScreen(0, 0, screenWidth, screenHeight)
//		defer robotgo.FreeBitmap(bitmap)
//
//		for _, elem := range pixelAndPos {
//			hasClicked := clickOnElement(elem[0][0], elem[0][1], elem[0][2], elem[1][0], elem[1][1])
//			if hasClicked {
//				break
//			}
//		}
//		fmt.Println("element not found, wait for next loop")
//		time.Sleep(2 * time.Second)
//	}
//}