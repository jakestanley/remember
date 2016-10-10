package main;

import "strconv"
import "image"
import "errors"
import "github.com/hajimehoshi/ebiten"

const PLAYER_SPEED = 1.8;
const PLAYER_WIDTH = 16;

type Player struct {
    x, y float64;
    sprite *ebiten.Image; // TODO sprite sheet
    rect image.Rectangle;
}

func NewPlayer() *Player { // TODO character sheet loader

    
    sprite, err := loadImage("_resources/dolf_1x1_character_sprite.png"); // TODO provide a non hard-coded value
    if(err != nil){
        // TODO something
    }

    rect := image.Rect(0, 0, PLAYER_WIDTH - 1, PLAYER_WIDTH - 1);

    p := Player{0,0,sprite,rect};

    return &p;
}

// TODO make this take a direction // TODO test
func (p *Player) Move(dv image.Point) error { // TODO bound 1,-1 or make a DirectionVector object

    x := dv.X;
    y := dv.Y;

    // check that we're in the valid range    
    err := validateMoveDirections(x, y);
    if(err != nil){
        // TODO some error
    }

    ps := PLAYER_SPEED * delta;
    p.x = p.x + (float64(x) * ps); // i hope this isn't too expensive
    p.y = p.y + (float64(y) * ps); // feels inefficient. TODO improve

    return nil;
}

func validateMoveDirections(x, y int) error { // TODO use t variable, e.g t.Fail, t.Error, t.Log

    err1 := validateMoveDirection(x);
    err2 := validateMoveDirection(y);
    if(err1 != nil || err2 != nil){
        errStr := "move direction failed to validate: ";
        xStr := strconv.Itoa(x);
        yStr := strconv.Itoa(y);
        errStr = errStr + "[" + xStr + "," + yStr + "]";
        errors.New(errStr);
    }
    return nil;
}

func validateMoveDirection(d int) error {

    if(d > 1){
        return errors.New("d is greater than +1");
    }
    if(d < -1){
        return errors.New("d is less than -1");
    }
    return nil;
}