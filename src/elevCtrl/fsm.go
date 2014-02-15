package elevCtrl
import "fmt"



type fsm_state int
const(
	IDLE fsm_state = iota
	DOORS_OPEN
	MOVING_DOWN
	MOVING_UP
)

type fsm_event int
const(
	no_event fsm_event = iota
	floorReached 
	orderUp
	orderDown
	emg
	obst
)

//state						//floorReached		//orderUp			//orderDown			//timeout	
IDLE_ev			:= []func(){action_dummy, 		action_dummy, 		action_orderUp,		action_orderDown,	action_dummy}	
DOORS_OPEN_ev	:= []func(){action_dummy, 	action_dummy, 	action_dummy,		action_orderDown,	action_dummy}	
MOVING_UP_ev	:= []func(){action_dummy, 	floorReached, 	action_orderUp,	action_dummy,		action_dummy}
MOVING_DOWN_ev	:= []func(){action_dummy, 	floorReached, 	action_orderUp,	action_dummy,		action_dummy} 
EMG_ev			:= []func(){action_dummy, 	action_dummy, 	action_dummy,		action_dummy,		action_dummy}
OBST_ev			:= []func(){action_dummy, 	action_dummy, 	action_dummy,		action_dummy,		action_dummy}


fsm_ctrl := [][]func(){
	ev_IDLE, 
	ev_DOORS_OPEN,
	ev_MOVING_UP,
	ev_MOVING_DOWN,
	ev_EMG,ev_OBST}
