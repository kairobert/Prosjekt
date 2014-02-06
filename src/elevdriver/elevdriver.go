package elevdriver

import (
    "fmt"
    "log"
    "os"
    "os/signal"
    "time"
)

type Direction int 

const N_FLOORS 4

const SPEED0 = 2048
const SPEED1 = 4024

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

var buttonChan chan Button
var sensorChan chan int
var motorChan chan Direction
var stopButtonChan chan bool
var obsChan chan bool


func setLight(floor int, dir Direction){
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
        printf("Error: Illegal floor or direction")
}

func clearLight(floor int, dir Direction){
    switch{  
    case floor == 1 && dir == NONE:
        Clear_bit(LIGHT_COMMAND1)
    case floor == 1 && dir == UP:
        Clear_bit(LIGHT_UP1)
 	case floor == 2 && dir == NONE:
        Clear_bit(LIGHT_COMMAND2)
    case floor == 2 && dir == UP:
        Clear_bit(LIGHT_UP2)
    case floor == 2 && dir == DOWN:
        Clear_bit(LIGHT_DOWN2)      
    case floor == 3 && dir == NONE:
        Clear_bit(LIGHT_COMMAND3)
    case floor == 3 && dir == UP:
        Clear_bit(LIGHT_UP3)
    case floor == 3 && dir == DOWN:
        Clear_bit(LIGHT_DOWN4)   
    case floor == 4 && dir == NONE:
        Clear_bit(LIGHT_COMMAND4)
    case floor == 4 && dir == DOWN:
        Clear_bit(LIGHT_DOWN4)
    default:
        printf("elevdriver: Error! Illegal floor or direction!")
}

func clearAllLights(){
    for int floor := 1; floor<N_FLOORS ; floor++{
        for int dir := 1; dir<N_DIR; dir+{
            clear_light(floor, dir)
        }
    }
}

func motorCtrl(motorChan channel Direction){
    lastDir := NONE
    for {
        switch newDir := <-motorChan{
        case UP:
            Clear_bit(MOTORDIR)
            Write_analog(MOTOR,SPEED1)
        case DOWN:
            Set_bit(MOTORDIR)
            Write_analog(MOTOR,SPEED1)
        case NONE:
            if lastDir == DOWN{
                Clear_bit(MOTORDIR)
                Write_analog(MOTOR,SPEED0)
            } 
            else if lastDir == UP{
                Set_bit(MOTORDIR)
                Write_analog(MOTOR,SPEED0)
            } else{
            fmt.Println("elevdriver: ERROR, illegal lastDir")
        default:
            Write_analog(MOTOR,SPEED0)
            fmt.Println("elevdriver: ERROR, illegal motor direction")
        }
        lastDir = newDir
    }
}

func listenButtons(ButtonChan buttonChan){
    var buttonMap = map[int]button{
        FLOOR_COMMAND1: {1, NONE},
        FLOOR_COMMAND2: {2, NONE},
        FLOOR_COMMAND3: {3, NONE},
        FLOOR_COMMAND4: {4, NONE},
        FLOOR_UP1:      {1, UP},
        FLOOR_UP2:      {2, UP},
        FLOOR_UP3:      {3, UP},
        FLOOR_DOWN2:    {2, DOWN},
        FLOOR_DOWN3:    {3, DOWN},
        FLOOR_DOWN4:    {4, DOWN},
    {

    buttonList := make(map[int]bool)
    for key, _ := range buttonMap {
        buttonList[key] = Read_bit(key)
    }    
    
     for key, button := range buttonMap {
         newValue := Read_bit(key)
         if newValue && !buttonList[key] {
            newButton := button
            go func() {
                buttonChan <- newButton
            }()
         }
      buttonList[key] = newValue
      }
}

func listenSensors(FloorChan sensorChan){
    var floorMap = map[int]int{
        SENSOR1: 1,
        SENSOR2: 2,
        SENSOR3: 3,
        SENSOR4: 4,
    }
    
    atFloor := false
    
    floorList := make(map[int]bool)
    for key, _ := range floorMap {
        floorList[key] = Read_bit(key)
    }
    
    for {
        time.Sleep(1E7)
        atFloor = false
        for key, floor := range floorMap {
            if Read_bit(key) {
                select {
                    case FloorChan <- floor:
                    default:
                }
                atFloor = true
            }
        }
            if !atFloor {
                select {
                    case FloorChan <- -1:
                    default:
                }
            }   
}

func InitElev(){

    if !IoInit(){
        fmt.Println("elevdriver: Driver init()... OK!")
	} else {
	    fmt.Println("elevdriver: Driver init()... FAILED!")
	}
	
	clearAllLights();
	
	buttonChan = make(chan Button)
    sensorChan = make(chan int)
    motorChan = make(chan Direction)

    emgStopChan = make(chan bool)
    obstructionChan = make(chan bool)
	
	go listenButtons(buttonChan)
	go listenSensors(sensorChan)
	
	go func() {
    // capture ctrl+c and stop elevator
        c := make(chan os.Signal)
        signal.Notify(c, os.Interrupt)
        s := <-c
        log.Printf("Got: %v, terminating program..", s)
        MotorStop()
        clearLights()
        os.Exit(1)
    }()
}

func OpenDoor(){
    Set_bit(DOOR_OPEN)
}

func CloseDoor(){
    Set_bit(DOOR_CLOSE)
}

func MotorUp() {
        motorChan <- UP
}

func MotorDown() {
        motorChan <- DOWN
}

func MotorStop() {
        motorChan <- NONE
}

func GetStopButton(){
    return Read_bit(STOP)
}

func SetStopButton(){
    Set_bit(STOP)
}

func ClearStopButton(){
    Clear_bit(STOP)
}




