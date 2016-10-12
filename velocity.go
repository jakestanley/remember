package main;

type Velocity struct {
    max float64;
    cur float64;
    inc float64;
    acc bool;
}

func NewVelocity(max, inc float64) *Velocity {
    v := Velocity{max, 0.0, inc, false};
    return &v;
}

func (v *Velocity) Accelerate() {
    v.acc = true;
}

func (v *Velocity) UpdateSpeed(delta float64) {
    new := 0.0;
    di := v.inc * delta; // delta * inc(/dec)
    if(v.acc){
        new = v.cur + di;
        if(new > v.max){
            new = v.max; // TODO use reference?
        }
    } else {
        new = 0.0;
        // new = v.cur - di;
        // if(new < 0.0){
        //     new = 0.0; // TODO const(?)
        // }
    }
    v.acc = false;
    v.cur = new;
    
}

func (v *Velocity) GetDeltaSpeed() float64 {
    return v.cur;
}