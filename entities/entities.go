package entities

type EntityPosition struct {
    X int
    Y int
}

func (p EntityPosition) IsColliding(other Position) bool {
    if p.X == other.X && p.Y == other.Y {
        return true
    }
    return false
}


type Entity interface {
    GetSprite() string
}
