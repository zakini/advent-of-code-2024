package utils

type Vector2 struct {
	X int
	Y int
}

func (a Vector2) Add(b Vector2) Vector2 {
	return Vector2{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

func (a Vector2) Subtract(b Vector2) Vector2 {
	return Vector2{
		X: a.X - b.X,
		Y: a.Y - b.Y,
	}
}

func (a Vector2) Clone() Vector2 {
	return Vector2{a.X, a.Y}
}
