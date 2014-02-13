package elevdriver

import (
    "fmt"
    "log"
    "os"
    "os/signal"
    "time"
)

type Direction int 

const SPEED0 = 2048
const SPEED1 = 4024

const N_FLOORS = 4
const (
    UP Direction = iota
    DOWN
    NONE
)

type Button struct{
	floor int    
	dir Direction           
}

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
        fmt.Println("Error: Illegal floor or direction")
	}
}

func ClearLight(floor int, dir Direction){
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
        fmt.Println("elevdriver: Error! Illegal floor or direction!")
		fmt.Println("dir: ", dir, ", floor: ",floor)
	}
}

func ClearAllLights(){
        ClearLight(1, UP)
        ClearLight(2, UP)
        ClearLight(3, UP)
        ClearLight(2, DOWN)
        ClearLight(3, DOWN)
        ClearLight(4, DOWN)
        ClearLight(1, NONE)
        ClearLight(2, NONE)
        ClearLight(3, NONE)
        ClearLight(4, NONE)
		CloseDoor()
        ClearStopButton()
}

func motorCtrl(motorChan chan Direction){
    lastDir := NONE
	newDir := NONE

    for {
        newDir=<-motorChan
		fmt.Println("motorCtrl recv newDir=", newDir)
		switch newDir{
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
            if lastDir == UP{
                Set_bit(MOTORDIR)
                Write_analog(MOTOR,SPEED0)
            } else{
            fmt.Println("elevdriver: ERROR, illegal lastDir")
			}
		default:
            Write_analog(MOTOR,SPEED0)
            fmt.Println("elevdriver: ERROR, illegal motor direction")
        }
        lastDir = newDir
    }
}

func listenButtons(buttonChan chan Button){
    var buttonMap = map[int]Button{
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
    }

   	buttonList := make(map[int]bool)
    for key, _ := range buttonMap {
        buttonList[key] = Read_bit(key)
    }    
    
     for key, button := range buttonMap {
         newValue := Read_bit(key)
         if newValue && !buttonList[key] {
            newButton := button
            go func() {		//why not select???
                buttonChan <- newButton
            }()
         }
      buttonList[key] = newValue
      }
}

func listenSensors(sensorChan chan int){
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
        time.Sleep(25 * time.Millisecond)
        atFloor = false
        for key, floor := range floorMap {
            if Read_bit(key) {
                select {		//why not go?
                    case sensorChan <- floor:
                    default:
                }
                atFloor = true
            }
        }
        if !atFloor {
	        select {
            case sensorChan <- -1:
            default:
            }
        }
	}   
}

func InitElev(	
		buttonChan chan Button,
		sensorChan chan int,
		motorChan chan Direction,
		stopButtonChan chan bool,
		obsChan chan bool){

	if !IoInit(){
        fmt.Println("elevdriver: Driver init()... OK!")
	} else {
	    fmt.Println("elevdriver: Driver init()... FAILED!")
	}
	
	ClearAllLights();
	
	go listenButtons(buttonChan)
	go listenSensors(sensorChan)
	go motorCtrl(motorChan)
	
	go func() {
    	// capture ctrl+c and stop elevator
        c := make(chan os.Signal)
        signal.Notify(c, os.Interrupt)
        s := <-c
        log.Printf("Got: %v, terminating program..", s)
        MotorStop(motorChan)
        ClearAllLights()
        os.Exit(1)
    }()
}

func OpenDoor(){
    	Set_bit(DOOR_OPEN)
}

func CloseDoor(){
    	Clear_bit(DOOR_OPEN)
}

func MotorUp(motorChan chan Direction) {
        motorChan <- UP
}

func MotorDown(motorChan chan Direction) {
        motorChan <- DOWN
		fmt.Println("motorChan <- DOWN")
}

func MotorStop(motorChan chan Direction) {
        motorChan <- NONE
}

func GetStopButton() bool{
    	return Read_bit(STOP)
}

func SetStopButton(){
    	Set_bit(STOP)
}

func ClearStopButton(){
    	Clear_bit(STOP)
}

func GetObstruction() bool{
	return Read_bit(OBSTRUCTION)
}
	
