package main

import (
        // "os"
        "errors"
        // "log"
        "image"
        "fmt"
        "time"
        "strconv"
        "image/color"
        // "github.com/aodin/date"
        "github.com/hajimehoshi/ebiten"
        "github.com/hajimehoshi/ebiten/ebitenutil"
        "github.com/hajimehoshi/ebiten/examples/common"
)

const (
    GAME_TITLE = "Remember When We Were Able To Go Outside?";
    SCREEN_WIDTH = 640;
    SCREEN_HEIGHT = 400;
    SCALE = 1;

    // time
    SECOND = 1000000000;
    HALF_SECOND = 500000000;
    THIRTIETH_SECOND = 33333333;

)

// render translation
var renderOffsetX float64 = 0.0; 
var renderOffsetY float64 = 0.0;

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

// images/resources
var testpng ebiten.Image;

func initialise() {
    // initArrays();
    prevTime = time.Now();
    prevSec = time.Now();
    delta = 1;
    // ebitenImage, imageImage, err := ebitenutil.NewImageFromFile("_resoures/64_px_square.png", ebiten.FilterNearest);

}

func loadImage(path string, filter ebiten.Filter) (*ebiten.Image, error) {
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
    img2, err := ebiten.NewImageFromImage(img, filter)
    if err != nil {
        return nil, err
    }
    return img2, err
}

func update() error {
    tick();

    // TODO flexible camera

    // calculate delta
    now := time.Now();
    timediff := now.Sub(prevTime);
    delta = (float64(timediff.Nanoseconds()) / SECOND) * 60;
    prevTime = now;

    if(ebiten.IsKeyPressed(ebiten.Key(ebiten.KeyQ))){
        return errors.New("USER_QUIT");
    }

    // handle W key
    if (ebiten.IsKeyPressed(ebiten.Key(ebiten.KeyW))){
        if(!keyDownW){
            keyDownW = true;
            // d = d.AddDays(1); // fire action
            // incrementHour(1);
        }
    } else {
        keyDownW = false;
    }

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

// increment tick and update delta time
func tick() {

}

func draw(screen *ebiten.Image) {

// trying to load and use the image pointer in the same block
    imgx, err := loadImage("_resources/64_px_square.png", ebiten.FilterNearest);
    if(err != nil){
        fmt.Println("failed to load image");
        return;
    }
    
    options := &ebiten.DrawImageOptions{}
    options.GeoM.Translate(renderOffsetX, renderOffsetY);
    renderOffsetX += delta;
    renderOffsetY += delta;

    screen.DrawImage(imgx, options);

    drawMap();
    drawEntities();
    drawUi();
    drawModals(); // dim if necessary
    
    fpsStrConv := strconv.FormatFloat(ebiten.CurrentFPS(), 'f', -1, 64);
    fpsString := "fps " + fpsStrConv;

    deltaStrConv := strconv.FormatFloat(delta, 'f', -1, 32);
    deltaString := "delta " + deltaStrConv;

    // err := screen.DrawImage(testpng, nil);
    // if(err != nil){
    //     log.Fatal();
    // }

    // // figured out text size and position and draw it on the screen
    width := common.ArcadeFont.TextWidth(deltaString) * SCALE;
    x := (SCREEN_WIDTH - width) / 2;
    common.ArcadeFont.DrawText(screen, deltaString, x, 20, SCALE, color.White);

    common.ArcadeFont.DrawText(screen, fpsString, 20, 20, SCALE, color.White);

}

func drawMap() {

}

func drawEntities() {

}

func drawUi() {

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
