package main;

import "time"

var prevTime time.Time;
var delta float64;

func initDelta() {
    prevTime = time.Now();
    delta = 1;
}

func updateDelta() {
    now := time.Now();
    timediff := now.Sub(prevTime);
    delta = (float64(timediff.Nanoseconds()) / SECOND) * 60;
    prevTime = now;
}

// TODO add a function that prevents functions from using an old delta value