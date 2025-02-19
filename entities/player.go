package entities

type Player struct {
    Pos EntityPosition
    sprite uint8
}

func NewPlayer(x int, y int) Player {
    return Player{
        Pos: EntityPosition{X: x, Y: y},
        sprite: 'P',
    }
}

func (p Player) GetSprite() string {
    return "\033[31m" + string(p.sprite) + "\033[0m"
}
