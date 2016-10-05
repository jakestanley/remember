package main;

import "image"
import "github.com/hajimehoshi/ebiten"
import "github.com/hajimehoshi/ebiten/ebitenutil"

const DEFAULT_FILTER = ebiten.FilterNearest;

func loadImage(path string) (*ebiten.Image, error) {
    file, err := ebitenutil.OpenFile(path)
    if err != nil {
        return nil, err
    }
    defer func() {
        _ = file.Close()
    }()
    img, _, err := image.Decode(file)
    if err != nil {
        return nil, err
    }
    img2, err := ebiten.NewImageFromImage(img, DEFAULT_FILTER)
    if err != nil {
        return nil, err
    }
    return img2, err
}