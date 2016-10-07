package main;

import "github.com/hajimehoshi/ebiten"

type Camera struct {
    pX float64;
    pY float64;
}

func NewCamera() *Camera {
    c := Camera{15.0,15.0};
    return &c;
}

func (c *Camera) Move(dX, dY float64) {
    c.pX = c.pX + dX;
    c.pY = c.pY + dY;
}

func (c *Camera) ApplyTransformation(geometry *ebiten.GeoM) {
    geometry.Translate(c.pX, c.pY);
}