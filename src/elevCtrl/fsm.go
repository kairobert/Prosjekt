package elevCtrl

import (
	//"fmt"
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

func (elev *Elevator)action_start_down(){
	elevdriver.MotorDown(elev.motorChan)
	elev.state = MOVING_UP
}

func (elev *Elevator)action_start_up(){
	elevdriver.MotorUp(elev.motorChan)
	elev.state = MOVING_DOWN
}

func (elev *Elevator)action_exec_order(){
	elevdriver.MotorStop(elev.motorChan)
	elevdriver.OpenDoor()
	//start_timer()	
	elev.state = DOORS_OPEN 
}

func (elev *Elevator)action_done(){
	elevdriver.CloseDoor()
	// stop_timer()	
	elev.state = IDLE
}

func (elev *Elevator)action_stop(){
	elevdriver.MotorStop(elev.motorChan)
	elev.state = EMG
}

func (elev *Elevator)action_pause(){
	elev.state = OBST
}

func action_dummy(){
	//do nothing
}

func (elev *Elevator)fsm_init(){
	//state						//start_down			//start_up				//exec_order			//timeout			//stop			//obst	
	IDLE_ev			:= []func(){elev.action_start_down, elev.action_start_up, 	elev.action_exec_order,	action_dummy,		action_dummy,	action_dummy}	
	DOORS_OPEN_ev	:= []func(){action_dummy, 			action_dummy, 			action_dummy,			elev.action_done,	action_dummy,	action_dummy}	
	MOVING_UP_ev	:= []func(){action_dummy, 			elev.action_exec_order, elev.action_start_up,	action_dummy,		action_dummy,	action_dummy}
	MOVING_DOWN_ev	:= []func(){action_dummy, 			elev.action_exec_order,	elev.action_start_up,	action_dummy,		action_dummy,	action_dummy} 
	EMG_ev			:= []func(){action_dummy, 			action_dummy, 			action_dummy,			action_dummy,		action_dummy,	action_dummy}
	OBST_ev			:= []func(){action_dummy, 			action_dummy, 			action_dummy,			action_dummy,		action_dummy,	action_dummy}

	table := [][]func(){
		IDLE_ev, 
		DOORS_OPEN_ev,
		MOVING_UP_ev,
		MOVING_DOWN_ev,
		EMG_ev,
		OBST_ev,
	}
	elev.fsm_table = table
}

func (elev *Elevator)fsm_update(event Event_t){
	elev.fsm_table[elev.state][event]()
}

