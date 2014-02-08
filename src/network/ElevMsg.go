package network


type button struct {
        floor int
        dir   Direction
}

buttonChan = make(chan button)
