package main;

// import "fmt"
import "image"
import "time"
import "math/rand"
import "github.com/hajimehoshi/ebiten"

const DEFAULT_TILE_SIZE = 16;

var ts TileSheet;
var tileArranger *TileArranger;

// creates a rect based on tile size and tile coordinates, e.g 1,1, should return the rectangle 
// 16, 16, 31, 31 // TODO test
func NewTileCustom(x, y, size int) image.Rectangle {
    x0 := x * size;
    y0 := y * size;
    x1 := x0 + size - 1;
    y1 := y0 + size - 1;
    r := image.Rect(x0, y0, x1, y1);
    return r;
}

// calls NewTile(x, y, size) with the DEFAULT_TILE_SIZE const
func NewTile(x, y int) image.Rectangle {
    r := NewTileCustom(x, y, DEFAULT_TILE_SIZE);
    return r;
}

// using ebiten ImageParts helps to reduce draw calls
// TODO make these dynamic. they're pretty static now... can't remove tiles easily
type TileSheet struct {
    sheet *ebiten.Image;
    textureRects map[string]image.Rectangle; // refs references this
    count int;
    textureRefs []string;
    textureRectsArray []image.Rectangle;
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

func (ts *TileSheet) Add(ref string, x, y int) { // TODO return error here
    rect := NewTile(x, y);
    ts.textureRects[ref] = rect; // image.Rect(0, 0, 15, 15);
    ts.textureRefs = append(ts.textureRefs, ref);
    ts.textureRectsArray = append(ts.textureRectsArray, rect);
    ts.count = ts.count + 1;
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

// TODO add a user defined string seed and default to time.Now().UnixNano() if empty or nil
func (ta *TileArranger) Randomize(widthInTiles, heightInTiles int) {
    rand.Seed(time.Now().UnixNano()); // doesn't work with the same number FIXME 
    for x := 0; x < widthInTiles; x++ {
        for y := 0; y < heightInTiles; y++ {
            adjX := x * DEFAULT_TILE_SIZE;
            adjY := y * DEFAULT_TILE_SIZE;
            texRef := ta.sheet.textureRefs[rand.Intn(ta.sheet.count)];
            ta.Add(texRef, image.Point{adjX, adjY});
        }
    }
}

func (ta *TileArranger) Add(texture string, xy image.Point) {
    ta.tile = append(ta.tile, ta.sheet.textureRects[texture]);
    ta.point = append(ta.point, xy);
    ta.count = ta.count + 1;
}

func initTiles() {
    // load ground textures
    ts = NewTileSheet("_resources/tiles.png");

    // create texture rectangles
    ts.Add("tx1", 0, 0); // image.Rect(0, 0, 15, 15);
    ts.Add("tx2", 1, 0); // image.Rect(16, 0, 31, 15);
    ts.Add("tx3", 0, 1); // image.Rect(0, 16, 15, 31);
    ts.Add("tx4", 1, 1); // image.Rect(16, 16, 31, 31);

    // create tile arranger
    tileArranger = NewTileArranger(ts);
    tileArranger.Randomize(8, 8);

    // if(err != nil){
    //     fmt.Println("error occurred adding to the tile arranger");
    // }

}

// TODO pass camera object and use it for translations
func drawTiles(screen *ebiten.Image, camera *Camera) { // TODO handle multiple tilesheets loaded into memory?
    imgx := ts.sheet;
    options := &ebiten.DrawImageOptions{};
    options.ImageParts = tileArranger;

    // apply view transformation
    camera.ApplyTransformation(&options.GeoM);

    screen.DrawImage(imgx, options);
}