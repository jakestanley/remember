package main;

import "image"

// move characters, objects, scenery, etc
var W_MOVE_UP image.Point;
var W_MOVE_DOWN image.Point;
var W_MOVE_LEFT image.Point;
var W_MOVE_RIGHT image.Point;

// move the camera
var CAM_PAN_UP image.Point;
var CAM_PAN_DOWN image.Point;
var CAM_PAN_LEFT image.Point;
var CAM_PAN_RIGHT image.Point;

func initDirectionVectors() {
    W_MOVE_UP       = image.Point{0.0, -1.0};
    W_MOVE_DOWN     = image.Point{0.0, 1.0};
    W_MOVE_LEFT     = image.Point{-1.0, 0.0};
    W_MOVE_RIGHT    = image.Point{1.0, 0.0};
    CAM_PAN_UP      = image.Point{0.0, 1.0};
    CAM_PAN_DOWN    = image.Point{0.0, -1.0};
    CAM_PAN_LEFT    = image.Point{1.0, 0.0};
    CAM_PAN_RIGHT   = image.Point{-1.0, 0.0};
}