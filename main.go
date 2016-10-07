package main

import (
        // "os"
        "errors"
        // "log"
        // "image"
        // "fmt"
        "time"
        "strconv"
        "image/color"
        // "github.com/aodin/date"
        "github.com/hajimehoshi/ebiten"
        // "github.com/hajimehoshi/ebiten/ebitenutil"
        "github.com/hajimehoshi/ebiten/examples/common"
)

const (
    // game constants
    GAME_TITLE = "Remember When We Were Able To Go Outside?";
    SCREEN_WIDTH = 288;
    SCREEN_HEIGHT = 216;
    SCREEN_MARGIN = 10;
    SCALE = 2;
    FONT_SCALE = 1;

    // time
    SECOND = 1000000000;
    HALF_SECOND = 500000000;

    // world
    MOVE_SPEED = 2;

)

// render translation
var camera *Camera;

// loop counter
var count int = 0;

// date/time
var day int = 1;
var prevTime time.Time;
var prevSec time.Time;
var delta float64; // time in nanos
var deltas []float64;

// prevents holding down keys
var keyDownW bool = false;
var keyDownE bool = false;

func initialise() {
    camera = NewCamera();
    initTiles();
    prevTime = time.Now();
    prevSec = time.Now();
    delta = 1;
    // ebitenImage, imageImage, err := ebitenutil.NewImageFromFile("_resoures/64_px_square.png", 
    // ebiten.FilterNearest);

}

func update() error {
    // FIXME delta on focus loss
    // TODO flexible camera

    // quit
    if(ebiten.IsKeyPressed(ebiten.Key(ebiten.KeyQ))){
        return errors.New("USER_QUIT");
    }

    // calculate delta
    now := time.Now();
    timediff := now.Sub(prevTime);
    delta = (float64(timediff.Nanoseconds()) / SECOND) * 60;
    prevTime = now;

    input();

    // handle E key
    if (ebiten.IsKeyPressed(ebiten.Key(ebiten.KeyE))){
        if(!keyDownE){
            keyDownE = true;
            // d = d.AddDays(30); // fire action
        }
    } else {
        keyDownE = false;
    }

    return nil;
}

func input() {

    ms := MOVE_SPEED * delta;

    // move up // TODO move the player character and make the camera move accordingly
    if(ebiten.IsKeyPressed(ebiten.Key(ebiten.KeyW))){
        camera.Move(0.0, ms);
    }

    // move down
    if(ebiten.IsKeyPressed(ebiten.Key(ebiten.KeyS))){
        camera.Move(0.0, 0.0 -ms);
    }

    // move left
    if(ebiten.IsKeyPressed(ebiten.Key(ebiten.KeyA))){
        camera.Move(ms, 0.0);
    }

    // move right
    if(ebiten.IsKeyPressed(ebiten.Key(ebiten.KeyD))){
        camera.Move(0.0 - ms, 0.0);
    }

}

func draw(screen *ebiten.Image) {
    
    drawTiles(screen, camera); // TODO merge into drawMap?

    // renderOffsetX += delta;
    // renderOffsetY += delta;

    drawMap();
    drawEntities();
    drawUi(screen);
    drawModals(); // dim if necessary

}

func drawMap() {

}

func drawEntities() {

}

func drawUi(screen *ebiten.Image) {
    
    fpsStrConv := strconv.FormatFloat(ebiten.CurrentFPS(), 'f', 2, 64);
    fpsString := "fps " + fpsStrConv;

    deltaStrConv := strconv.FormatFloat(delta, 'f', 2, 32);
    deltaString := "delta " + deltaStrConv;

    // // figured out text size and position and draw it on the screen
    width := common.ArcadeFont.TextWidth(deltaString);
    x := (SCREEN_WIDTH - width) / 2; // center
    y := SCREEN_HEIGHT - common.ArcadeFont.TextHeight("H") - SCREEN_MARGIN;

    common.ArcadeFont.DrawText(screen, deltaString, x, y, FONT_SCALE, color.White);
    common.ArcadeFont.DrawText(screen, fpsString, SCREEN_MARGIN, y, FONT_SCALE, color.White);

}

func drawModals() {

}

func loop(screen *ebiten.Image) error {

    // delta should be 1 at 60fps, 2 at 30fps, etc
    
    // update and return error if one is thrown. will probably be a user quit
    err := update();
    if err != nil{
        return err;
    }

    // render
    draw(screen);

    // return no error, allowing the game to continue
    return nil;
}

func deinitialise() {
    // quit
}

func main() { // int?
    initialise();
    // if loop returns an error, quit
    if err := ebiten.Run(loop, SCREEN_WIDTH, SCREEN_HEIGHT, SCALE, GAME_TITLE); err != nil {
        deinitialise();
    }
}
