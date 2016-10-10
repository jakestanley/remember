package main;

import "image"
import "github.com/hajimehoshi/ebiten"

const PAN_SPEED = 2;

type Camera struct {
    pX float64;
    pY float64;
}

func NewCamera() *Camera {
    c := Camera{15.0,15.0};
    return &c;
}

func (c *Camera) Move(dv image.Point) {

    x := dv.X;
    y := dv.Y;

    cps := PAN_SPEED * delta; // camera pan speed
    c.pX = c.pX + (float64(x) * cps);
    c.pY = c.pY + (float64(y) * cps);
}

func (c *Camera) ApplyTransformation(geometry *ebiten.GeoM) {
    geometry.Translate(c.pX, c.pY);
}