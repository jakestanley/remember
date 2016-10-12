package main;

import "image"
import "github.com/hajimehoshi/ebiten"

// default camera speed values
const PAN_SPEED = 1.4;
const INC_SPEED = 0.125;

type Camera struct {
    pX float64;
    pY float64;
    v *Velocity;
}

func NewCameraV(v *Velocity) *Camera {
    c := Camera{0, 0, v};
    return &c;
}

func NewCamera() *Camera {
    v := NewVelocity(PAN_SPEED, INC_SPEED); // velocity
    c := Camera{0, 0, v}; // TODO remove magic values
    return &c;
}

func (c *Camera) Move(dv image.Point) {

    cps := c.v.GetDeltaSpeed();

    x := dv.X;
    y := dv.Y;

    c.pX = c.pX + (float64(x) * cps);
    c.pY = c.pY + (float64(y) * cps);
}

func (c *Camera) ApplyTransformation(geometry *ebiten.GeoM) {
    geometry.Translate(c.pX, c.pY);
}