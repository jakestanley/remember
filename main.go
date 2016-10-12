package main

import (
        // "os"
        "errors"
        // "log"
        // "image"
        // "fmt"
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

// date/time
var day int = 1;

// world
var player *Player;

func initialise() {

    // initialise arrays, etc
    initDirectionVectors();

    camera = NewCamera();
    // initialise delta
    initDelta();

    initTiles();
    player = NewPlayer();

}

func update() error {

    // update delta
    updateDelta();

    // quit
    if(ebiten.IsKeyPressed(ebiten.Key(ebiten.KeyQ))){
        return errors.New("USER_QUIT");
    }

    // do other input
    input();


    return nil;
}

func input() {

    // move up // TODO move the player character and make the camera move accordingly
    if(ebiten.IsKeyPressed(ebiten.Key(ebiten.KeyW))){
        player.Move(W_MOVE_UP);
        camera.Move(CAM_PAN_UP);
    }

    // move down
    if(ebiten.IsKeyPressed(ebiten.Key(ebiten.KeyS))){
        player.Move(W_MOVE_DOWN);
        camera.Move(CAM_PAN_DOWN);
    }

    // move left
    if(ebiten.IsKeyPressed(ebiten.Key(ebiten.KeyA))){
        player.Move(W_MOVE_LEFT);
        camera.Move(CAM_PAN_LEFT);
    }

    // move right
    if(ebiten.IsKeyPressed(ebiten.Key(ebiten.KeyD))){
        player.Move(W_MOVE_RIGHT);
        camera.Move(CAM_PAN_RIGHT);
    }

    // camera up
    // if(ebiten.IsKeyPressed(ebiten.Key(ebiten.KeyUp))){
        
    // }

    // // camera down
    // if(ebiten.IsKeyPressed(ebiten.Key(ebiten.KeyDown))){
        
    // }

    // // camera left
    // if(ebiten.IsKeyPressed(ebiten.Key(ebiten.KeyLeft))){
        
    // }

    // // camera right
    // if(ebiten.IsKeyPressed(ebiten.Key(ebiten.KeyRight))){
        
    // }
    

}

func draw(screen *ebiten.Image) {
    
    drawTiles(screen, camera); // TODO merge into drawMap?
    drawMap();
    drawEntities(screen);
    drawUi(screen);
    drawModals(); // dim if necessary

}

func drawMap() {

}

func drawEntities(screen *ebiten.Image) {

    // screen *ebiten.Image, camera *Camera

    imgx := player.sprite; // TODO for each, and have many sheets
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

    // get text height
    height := common.ArcadeFont.TextHeight("H");

    // // figured out text size and position and draw it on the screen
    width := common.ArcadeFont.TextWidth(deltaString);
    x := (SCREEN_WIDTH - width) / 2; // center
    y := SCREEN_HEIGHT - height - SCREEN_MARGIN;

    // player position debug
    posStr := "Player position: " + strconv.FormatFloat(player.x, 'f', 1, 64);
    posStr += ", " + strconv.FormatFloat(player.y, 'f', 1, 64);

    common.ArcadeFont.DrawText(screen, deltaString, x, y, FONT_SCALE, color.White);
    common.ArcadeFont.DrawText(screen, fpsString, SCREEN_MARGIN, y, FONT_SCALE, color.White);

    posStrY := y - height - SCREEN_MARGIN;
    common.ArcadeFont.DrawText(screen, posStr, SCREEN_MARGIN, posStrY, FONT_SCALE, color.White);

}

func drawModals() {
    // TODO window tiles. doesn't need to be fancy
    // for production build, we can merge everything into one big sprite sheet. use coordinates
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
