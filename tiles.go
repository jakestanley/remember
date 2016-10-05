package main;

import "fmt"
import "image"
import "github.com/hajimehoshi/ebiten"

var ts TileSheet;
var tileArranger *TileArranger;

// using ebiten ImageParts helps to reduce draw calls // TODO move this tile stuff into a separate go source file
type TileSheet struct {
    sheet *ebiten.Image;
    textureRects map[string]image.Rectangle;
}

func NewTileSheet(path string) TileSheet {
    ts := TileSheet{};
    sheet, err := loadImage(path);
    ts.textureRects = make(map[string]image.Rectangle);
    if(err != nil){
        // TODO do something
    }
    ts.sheet = sheet;
    return ts;
}

// tile arranger
type TileArranger struct {
    sheet *TileSheet;
    count int;
    tile []image.Rectangle;
    point []image.Point;
}

func NewTileArranger(sheet TileSheet) *TileArranger {
    ta := TileArranger{};
    ta.sheet = &sheet;
    ta.count = 0;
    return &ta;
}

func (ta *TileArranger) Len() (int) {
    return ta.count;
}

func (ta *TileArranger) Src(i int) (x0, y0, x1, y1 int) {
    t := ta.tile[i];
    return t.Min.X, t.Min.Y, t.Max.X, t.Max.Y;
}

func (ta *TileArranger) Dst(i int) (x0, y0, x1, y1 int) {
    // point at which to start drawing the tile
    p := ta.point[i]
    x0 = p.X;
    y0 = p.Y;

    // point at which the tile should end
    t := ta.tile[i];
    x1 = x0 + t.Max.X - t.Min.X;
    y1 = y0 + t.Max.Y - t.Min.Y;

    return x0, y0, x1, y1;
}

func (ta *TileArranger) Add(texture string, xy image.Point) error {
    ta.tile = append(ta.tile, ta.sheet.textureRects[texture]);
    ta.point = append(ta.point, xy);
    ta.count = ta.count + 1;
    return nil;
}

func initTiles() {
    // load ground textures
    ts = NewTileSheet("_resources/tiles.png");

    // create texture rectangles
    ts.textureRects["tx1"] = image.Rect(0, 0, 15, 15);
    ts.textureRects["tx2"] = image.Rect(16, 0, 31, 15);
    ts.textureRects["tx3"] = image.Rect(0, 16, 15, 31);
    ts.textureRects["tx4"] = image.Rect(16, 16, 31, 31);

    // create tile arranger
    tileArranger = NewTileArranger(ts);
    err := tileArranger.Add("tx1", image.Point{50,20});
    err = tileArranger.Add("tx2", image.Point{50,40});
    err = tileArranger.Add("tx3", image.Point{50,60});
    err = tileArranger.Add("tx4", image.Point{50,80});

    if(err != nil){
        fmt.Println("error occurred adding to the tile arranger");
    }

}