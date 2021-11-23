package main

import (
	"context"
	"flag"
	"math/rand"
	"syscall"
	"time"

	"github.com/micmonay/keybd_event"
	"golang.design/x/hotkey"
	"golang.design/x/mainthread"
)

var active bool = false
var dll = syscall.NewLazyDLL("user32.dll")
var procMouse = dll.NewProc("mouse_event")

var (
	min = flag.Int("min", 1000, "min in ms")
	max = flag.Int("max", 2500, "max in ms")
)

func main() {
	flag.Parse()
	mainthread.Init(fn)
}
func fn() { // Use fn as the actual main function.
	var (
		k = hotkey.KeyP
	)
	go runBot()
	// Register a desired hotkey.
	hk, err := hotkey.Register(nil, k)
	if err != nil {
		panic("hotkey registration failed")
	}

	// Start listen hotkey event whenever you feel it is ready.
	triggered := hk.Listen(context.Background())
	for range triggered {
		println("hotkey ctrl+s is triggered")
		active = !active
	}
}

//Bot that is running
func runBot() {
	//Event for A press
	kba, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}
	kba.SetKeys(keybd_event.VK_A)

	//Event for D press
	kbd, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}
	kbd.SetKeys(keybd_event.VK_D)

	for {
		r := time.Duration(int64(rand.Intn(250-100) + 100))
		r1 := time.Duration(int64(rand.Intn(*max-*min) + *min))
		if active {
			//A press
			time.Sleep(time.Millisecond * r)
			kba.Press()
			time.Sleep(r1 * time.Millisecond)
			kba.Release()
			mouse_click(MOUSEEVENTF_LEFTDOWN | MOUSEEVENTF_LEFTUP)
		}
		r = time.Duration(int64(rand.Intn(250-100) + 100))
		r1 = time.Duration(int64(rand.Intn(*max-*min) + *min))
		if active {
			//D press
			time.Sleep(time.Millisecond * r)
			kbd.Press()
			time.Sleep(r1 * time.Millisecond)
			kbd.Release()
			mouse_click(MOUSEEVENTF_LEFTDOWN | MOUSEEVENTF_LEFTUP)
		}
	}
}

const (
	MOUSEEVENTF_ABSOLUTE = 0x8000
	MOUSEEVENTF_MOVE     = 0x0001

	MOUSEEVENTF_LEFTDOWN = 0x0002
	MOUSEEVENTF_LEFTUP   = 0x0004

	MOUSEEVENTF_MIDDLEDOWN = 0x0020
	MOUSEEVENTF_MIDDLEUP   = 0x0040

	MOUSEEVENTF_RIGHTDOWN = 0x0008
	MOUSEEVENTF_RIGHTUP   = 0x0010

	MOUSEEVENTF_WHEEL  = 0x0800
	MOUSEEVENTF_HWHEEL = 0x0 * 1000

	MOUSEEVENTF_XDOWN = 0x0080
	MOUSEEVENTF_XUP   = 0x0100
)

func mouse_click(intype uint32) {
	procMouse.Call(uintptr(intype), 0, 0, 0, 0)
}
