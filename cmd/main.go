package main

import (
	"bufio"
	dungeonmap "dungeon-crawler-go/dungeon_map"
	"dungeon-crawler-go/entities"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("\x1B[2J\x1B[HWelcome to my dungeon crawler game!")
	dungeonMap := dungeonmap.NewMap(10, 20)

	isRunning := true
	inputBuffer := make(chan string)
	frameDuration := time.Second / 5 
	frameCount := 0
	fpsTime := time.Now()
	previousTime := time.Now()
	enemies := []entities.Enemy{}
	player := entities.NewPlayer(1, 1)

    time.Sleep(time.Second * 2)

	go getInput(inputBuffer)

	for isRunning {
		currentTime := time.Now()
		dt := currentTime.Sub(previousTime)

		select {
		case input := <-inputBuffer:
            switch input {
            case "w\n":
                player.Pos.Y--
            case "a\n":
                player.Pos.X--
            case "s\n":
                player.Pos.Y++
            case "d\n":
                player.Pos.X++
            }
		default:
			if dt < frameDuration {
				time.Sleep(frameDuration - dt)
				currentTime = time.Now()
				dt = currentTime.Sub(previousTime)
			}
			frameCount++
            fmt.Printf("\x1B 7")
			if currentTime.Sub(fpsTime) >= time.Second {
				fps := float64(frameCount) / currentTime.Sub(fpsTime).Seconds()
				fmt.Printf("\x1B[2;%dHFPS: %.2f\n", (dungeonMap.Cols*2)+5, fps)
				frameCount = 0
				fpsTime = currentTime
			}
			fmt.Printf("\x1B[2;0H")
			Render(dungeonMap, enemies, player)
            fmt.Printf("\x1B 8")

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

func Render(dungeonMap *dungeonmap.DungeonMap, enemies []entities.Enemy, player entities.Player) {
	mapBuffer := make([][]string, dungeonMap.Rows)
	for i := range mapBuffer {
		mapBuffer[i] = make([]string, dungeonMap.Cols)
	}

	for i, row := range dungeonMap.Tiles {
		for j, tile := range row {
			mapBuffer[i][j] = tile.GetSprite()
		}
	}
	for _, enemy := range enemies {
		mapBuffer[enemy.Pos.Y][enemy.Pos.X] = enemy.GetSprite()
	}
	mapBuffer[player.Pos.Y][player.Pos.X] = player.GetSprite()

	for _, row := range mapBuffer {
		for _, sprite := range row {
			fmt.Printf("%s ", sprite)
		}
		fmt.Println()
	}
}
