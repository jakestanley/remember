package main;

import "fmt"
import "strconv"
import "image"
import "errors"
import "github.com/hajimehoshi/ebiten"

const PLAYER_SPEED = 1.4;
const PLAYER_SPEED_INC = 0.0125;
const PLAYER_WIDTH = 16;

type Player struct {
    x, y float64;
    sprite *ebiten.Image; // TODO sprite sheet
    rect image.Rectangle;
    v *Velocity;
}

func NewPlayer(v *Velocity) *Player { // TODO character sheet loader

    sprite, err := loadImage("_resources/dolf_1x1_character_sprite.png"); // TODO remove hard-coding
    if(err != nil){
        // TODO something
    }

    rect := image.Rect(0, 0, PLAYER_WIDTH - 1, PLAYER_WIDTH - 1);

    initX := float64((SCREEN_WIDTH - PLAYER_WIDTH) / 2);
    initY := float64((SCREEN_HEIGHT - PLAYER_WIDTH) / 2);

    p := Player{initX, initY, sprite, rect, v};

    return &p;
}

func (p *Player) Move(dv image.Point) error {

    ps := p.v.GetDeltaSpeed();

    x := dv.X;
    y := dv.Y;

    // check that we're in the valid range    
    err := validateMoveDirections(x, y); // TODO probably remove this now or change the test
    if(err != nil){
        // TODO some error
    }

    p.x = p.x + (float64(x) * ps); // i hope this isn't too expensive
    p.y = p.y + (float64(y) * ps); // feels inefficient. TODO improve

    fmt.Println("cur speed: " + strconv.FormatFloat(ps, 'f', 2, 32));

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