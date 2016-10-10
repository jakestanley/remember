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

)

// render translation
var camera *Camera;

// loop counter
var count int = 0;

// date/time
var day int = 1;
var prevTime time.Time;
var delta float64; // time in nanos
var deltas []float64;

// prevents holding down keys
var keyDownW bool = false;
var keyDownE bool = false;

// world
var player *Player;

func initialise() {

    // initialise arrays, etc
    initDirectionVectors();

    camera = NewCamera();
    initTiles();
    prevTime = time.Now();
    delta = 1;
    player = NewPlayer();
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

    // TODO update camera

    return nil;
}

func input() {

    // move up // TODO move the player character and make the camera move accordingly
    if(ebiten.IsKeyPressed(ebiten.Key(ebiten.KeyW))){
        player.Move(W_MOVE_UP);
    }

    // move down
    if(ebiten.IsKeyPressed(ebiten.Key(ebiten.KeyS))){
        player.Move(W_MOVE_DOWN);
    }

    // move left
    if(ebiten.IsKeyPressed(ebiten.Key(ebiten.KeyA))){
        player.Move(W_MOVE_LEFT);
    }

    // move right
    if(ebiten.IsKeyPressed(ebiten.Key(ebiten.KeyD))){
        player.Move(W_MOVE_RIGHT);
    }

    // camera up // TODO move the player character and make the camera move accordingly
    if(ebiten.IsKeyPressed(ebiten.Key(ebiten.KeyUp))){
        camera.Move(CAM_PAN_UP);
    }

    // camera down
    if(ebiten.IsKeyPressed(ebiten.Key(ebiten.KeyDown))){
        camera.Move(CAM_PAN_UP); // TODO change this to 1,-1 bounded vector, make GetCamSpeed()
    }

    // camera left
    if(ebiten.IsKeyPressed(ebiten.Key(ebiten.KeyLeft))){
        camera.Move(CAM_PAN_LEFT);
    }

    // camera right
    if(ebiten.IsKeyPressed(ebiten.Key(ebiten.KeyRight))){
        camera.Move(CAM_PAN_RIGHT);
    }
    

}

func draw(screen *ebiten.Image) {
    
    drawTiles(screen, camera); // TODO merge into drawMap?

    // renderOffsetX += delta;
    // renderOffsetY += delta;

    drawMap();
    drawEntities(screen);
    drawUi(screen);
    drawModals(); // dim if necessary

}

func drawMap() {

}

func drawEntities(screen *ebiten.Image) {

    // screen *ebiten.Image, camera *Camera

    imgx := player.sprite;
    options := &ebiten.DrawImageOptions{};
    options.GeoM.Translate(player.x, player.y);
    // options.ImageParts = tileArranger;

    // apply view transformation
    // camera.ApplyTransformation(&options.GeoM);

    // apply view transformation
    camera.ApplyTransformation(&options.GeoM);

    screen.DrawImage(imgx, options);

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
