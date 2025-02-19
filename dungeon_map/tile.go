package dungeonmap

type Tile struct {
	sprite uint8
	IsWall bool
}

func (t Tile) GetSprite() string {
    return string(t.sprite)
}
