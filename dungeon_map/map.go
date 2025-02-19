package dungeonmap

import "fmt"

type DungeonMap struct {
	Rows  int
	Cols  int
	Tiles [][]Tile
}

func NewMap(rows, cols int) *DungeonMap {
    dungeonMap := DungeonMap {
        Rows: rows,
        Cols: cols,
        Tiles: [][]Tile{},
    }
    for i := 0; i < rows; i++ {
        row := []Tile{}
        for j := 0; j < cols; j++ {
            var tile Tile
            if i == 0 || j == 0 || i == rows-1 || j == cols-1 {
                tile.Sprite = '0'
                tile.IsWall = true
            } else {
                tile.Sprite = ' '
                tile.IsWall = false
            }
            row = append(row, tile)
        }
        dungeonMap.Tiles = append(dungeonMap.Tiles, row)
    }

    return &dungeonMap
}

func (d *DungeonMap) PrintMap() {
    fmt.Printf("\x1B[H")
    for i := 0; i < d.Rows; i++ {
        for j := 0; j < d.Cols; j++ {
            tile := d.Tiles[i][j]
            fmt.Printf("%c ", tile.Sprite)
        }
        fmt.Println()
    }
}
