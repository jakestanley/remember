package main;

import "testing"

var tests [4]TileTestIO;

type TileTestIO struct {
    inputX int;
    inputY int;
    expectedMinX int;
    expectedMinY int;
    expectedMaxX int;
    expectedMaxY int;
}

func initTileTestIO(){
    tests[0] = TileTestIO{0, 0, 0, 0, 15, 15};
    tests[1] = TileTestIO{1, 0, 16, 0, 31, 15};
    tests[2] = TileTestIO{0, 1, 0, 16, 15, 31};
    tests[3] = TileTestIO{1, 1, 16, 16, 31, 31};
}

func TestTiles(t * testing.T) {
    if(DEFAULT_TILE_SIZE != 16){
        t.Error("Expected const DEFAULT_TILE_SIZE to equal 16, but it equals ", DEFAULT_TILE_SIZE);
    }
}

func TestNewTile(t *testing.T) {

    initTileTestIO();

    failures := 0;

    for i := 0; i < len(tests); i++ {

        inputX := tests[i].inputX;
        inputY := tests[i].inputY;
        expectedMinX := tests[i].expectedMinX;
        expectedMinY := tests[i].expectedMinY;
        expectedMaxX := tests[i].expectedMaxX;
        expectedMaxY := tests[i].expectedMaxY;

        tile := NewTileCustom(inputX, inputY, DEFAULT_TILE_SIZE);
        fail := (tile.Min.X != expectedMinX) || (tile.Min.Y != expectedMinY) || (tile.Max.X != expectedMaxX) || (tile.Max.Y != expectedMaxY);

        if(fail == true){
            failures = failures + 1;
        }
    }

    if(failures > 0){
        t.Error("TestNewTile failed. Incorrect results: ", failures);
    }
}