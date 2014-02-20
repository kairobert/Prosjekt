package elevCtrl

import (
	"fmt"
	"elevdriver"
)

type State_t int
const(
	IDLE State_t = iota
	DOORS_OPEN
	MOVING_DOWN
	MOVING_UP
	EMG
	OBST
	OBST_EMG
)

type Event_t int
const(
	start_down Event_t = iota
	start_up
	exec_order 
	timeout
	stop
	obst
)

/* 	FSM Actions */
func (elev *Elevator)action_start_down(){
	elevdriver.MotorDown(elev.motorChan)
	elev.state = MOVING_DOWN
	elev.lastDir = elevdriver.DOWN
	fmt.Println("fsm: MOVING_DOWN")
}

func (elev *Elevator)action_start_up(){
	elevdriver.MotorUp(elev.motorChan)
	elev.state = MOVING_UP
	elev.lastDir = elevdriver.UP
	fmt.Println("fsm: MOVING_UP")
}

func (elev *Elevator)action_exec_order(){
	elevdriver.MotorStop(elev.motorChan)
	elevdriver.OpenDoor()
	//start_timer()	
	//order_executed()
	elev.state = DOORS_OPEN 
	elevdriver.setLight(elev.lastFloor, elev.lastDirection)
	fmt.Println("fsm: DOORS_OPEN")
}

func (elev *Elevator)action_done(){
	elevdriver.CloseDoor()
	// stop_timer()	
	elev.state = IDLE
	fmt.Println("fsm: IDLE")
}

func (elev *Elevator)action_stop(){
	elevdriver.MotorStop(elev.motorChan)
	elev.state = EMG
}

func (elev *Elevator)action_pause(){
	elev.state = OBST
}

func action_dummy(){
	fmt.Println("fsm: dummy!")
}

/* Finite State Machine */
func (elev *Elevator)fsm_init(){
	elev.fsm_table = [][]func(){
/*STATES:	\	EVENTS:		//start_down			//start_up				//exec_order			//timeout			
/* IDLE			*/	[]func(){elev.action_start_down, elev.action_start_up, 	elev.action_exec_order,	action_dummy},	
/* DOORS_OPEN	*/	[]func(){action_dummy, 			action_dummy, 			action_dummy,			elev.action_done},	
/* MOVING_UP	*/	[]func(){action_dummy, 			action_dummy, 			elev.action_stop,		action_dummy},
/* MOVING_DOWN	*/	[]func(){action_dummy, 			action_dummy,			elev.action_stop,		action_dummy}, 
/* EMG				[]func(){action_dummy, 			action_dummy, 			action_dummy,			action_dummy},*/
/* OBST				[]func(){action_dummy, 			action_dummy, 			action_dummy,			action_dummy},*/
	}
}

/* FSM help functions */
func (elev *Elevator)fsm_update(event Event_t){
	elev.fsm_table[elev.state][event]()
}

func (elev *Elevator)should_stop(floor int) bool{
	//communicate with orders and check if current floor got pending orders in correct direction
	stop := true	
	return stop
}

func (elev *Elevator)get_nearest_order() elevdriver.Direction_t{
	//communicate with orders and get next direction
	return elevdriver.UP
}


func (elev *Elevator)fsm_handle_events(){
	select{

	case <- elev.stopButtonChan:
		elev.fsm_update(stop)

	case <- elev.obsChan:
		elev.fsm_update(obst)

	case floor:=<- elev.floorChan:
		if floor != -1{
			elev.lastFloor = floor
			if elev.should_stop(floor){
				elev.fsm_update(exec_order)
			}
		}
	
	case <- elev.newOrder:
		nextDir := elev.get_nearest_order()
		if elev.state == IDLE{
			switch nextDir{
			case elevdriver.UP:
				elev.fsm_update(start_down)	
	
			case elevdriver.DOWN:
				elev.fsm_update(start_up)

			case elevdriver.NONE:
				elev.fsm_update(exec_order)
			}
		}
	}
}
