package entities

type Enemy struct {
    Pos EntityPosition
    sprite uint8
}

func (e Enemy) GetSprite() string {
    return string(e.sprite)
}
