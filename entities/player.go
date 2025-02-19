package entities

type Player struct {
    Pos EntityPosition
    sprite uint8
}

func (p Player) GetSprite() string {
    return string(p.sprite)
}
