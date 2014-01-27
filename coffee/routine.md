type Routine interface {
    resume(VmInside, cycles)
}

type VmInside interface {
    send(chan, some integer) // on input chan
    sendPage(chan, a Page) // on input chan
    listen(chan) // only on output chan
    reset(chan)

    select() (chan, value)
}

Page 0 is noop page

type VmOutside interface {
    resume(cycles)
    bindChanIn(chanId, chan)
    bindChanOut(chanId, chan)
}

type Vm {
    VmInside
    VmOutside
    routine Routine
}

a resume() must exit for several reasons:
    halt: the routine is ready to die
    panic: the routine did some thing wrong and is ready to die
    cycleout: the routine needs more cycles to run
    select: the routine gives up running and listen on message channels


send(chan, some integer)
send(chan, some page)
listen(chan)
listen(chan)
chan = select()
if isRecvChan(chan) {
    handle recv data
} else if isSendChan(chan) {
    maybe send some more
}
