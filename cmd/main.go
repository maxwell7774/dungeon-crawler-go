package main

import (
	"bufio"
	dungeonmap "dungeon-crawler-go/dungeon_map"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("\x1B[2JWelcome to my dungeon crawler game!")
	dungeonMap := dungeonmap.NewMap(10, 20)

	isRunning := true
	inputBuffer := make(chan string)
	frameDuration := time.Second / 2
	frameCount := 0
	fpsTime := time.Now()
	previousTime := time.Now()

	go getInput(inputBuffer)

	for isRunning {
		currentTime := time.Now()
		dt := currentTime.Sub(previousTime)

		select {
		case input := <-inputBuffer:
			fmt.Println(input)
		default:
			if dt < frameDuration {
				time.Sleep(frameDuration - dt)
				currentTime = time.Now()
				dt = currentTime.Sub(previousTime)
			}
			frameCount++
			if currentTime.Sub(fpsTime) >= time.Second {
				fps := float64(frameCount) / currentTime.Sub(fpsTime).Seconds()
				fmt.Printf("\x1B[0;%dHFPS: %.2f\n", (dungeonMap.Cols * 2)+5, fps)
				frameCount = 0
				fpsTime = currentTime
			}
            dungeonMap.PrintMap()
			previousTime = currentTime
		}
	}

	dungeonMap.PrintMap()
}

func getInput(inputBuffer chan string) {
	reader := bufio.NewReader(os.Stdin)
	for true {
		input, _ := reader.ReadString('\n')
		inputBuffer <- input
	}
}
