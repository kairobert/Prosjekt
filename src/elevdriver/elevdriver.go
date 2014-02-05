package elevdriver

import (
    "fmt"
    "log"
    "os"
    "os/signal"
    "time"
)

type Direction int 

const (
    UP Direction //= iota
    DOWN
    NONE
    N_DIR
)

type Button struct{
    dir Direction       
    floor int
}

buttonChan = make(chan Button)
sensorChan = make(chan int)
motorChan = make(chan Direction)

emgStopChan = make(chan bool)
obstructionChan = make(chan bool)

func listen_buttons(){
}

func listen_sensors({
}

func set_light(floor int, dir Direction){
    switch{  
    case floor == 1 && dir == NONE:
            Set_bit(LIGHT_COMMAND1)
    case floor == 1 && dir == UP:
            Set_bit(LIGHT_UP1)

 	case floor == 2 && dir == NONE:
        Set_bit(LIGHT_COMMAND2)
    case floor == 2 && dir == UP:
        Set_bit(LIGHT_UP2)
    case floor == 2 && dir == DOWN:
        Set_bit(LIGHT_DOWN2)
        
    case floor == 3 && dir == NONE:
        Set_bit(LIGHT_COMMAND3)
    case floor == 3 && dir == UP:
        Set_bit(LIGHT_UP3)
    case floor == 3 && dir == DOWN:
        Set_bit(LIGHT_DOWN4)
        
    case floor == 4 && dir == NONE:
        Set_bit(LIGHT_COMMAND4)
    case floor == 4 && dir == DOWN:
        Set_bit(LIGHT_DOWN4)
    default:
        printf("Error: Illegal floor or dir")
}

func clear_light(){
}

func clear_lights(){
    for
        clear_light(i)
}

func motor_ctrl(motorChan channel){ // fix channel type
}

func InitElev(){
	
	go listen_buttons(buttonChan)
	go listen_sensors(sensorChan)
	
	go func() {
    // capture ctrl+c and stop elevator
        c := make(chan os.Signal)
        signal.Notify(c, os.Interrupt)
        s := <-c
        log.Printf("Got: %v, terminating program..", s)
        //Set motorspead = 0
        clearLights()
        os.Exit(1)
    }()
}
