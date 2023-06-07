package main

import (
	"fmt"
	"log"
	"syscall"
	"time"
	"unsafe"
)

const (
	MOUSEEVENTF_LEFTDOWN = 0x0002
	MOUSEEVENTF_LEFTUP   = 0x0004
	// 0x8000: キーが押されている
	// 0x0000: キーが押されていない
	KEY_PRESSED = 0x8000

	KEY_SHIFT = 0x10
	KEY_CTRL  = 0x11
	KEY_ALT   = 0x12
	KEY_A     = 0x41
	KEY_B     = 0x42
	KEY_C     = 0x43
	KEY_D     = 0x44
	KEY_E     = 0x45
	KEY_F     = 0x46
	KEY_G     = 0x47
	KEY_H     = 0x48
	KEY_I     = 0x49
	KEY_J     = 0x4A
	KEY_K     = 0x4B
	KEY_L     = 0x4C
	KEY_M     = 0x4D
	KEY_N     = 0x4E
	KEY_O     = 0x4F
	KEY_P     = 0x50
	KEY_Q     = 0x51
	KEY_R     = 0x52
	KEY_S     = 0x53
	KEY_T     = 0x54
	KEY_U     = 0x55
	KEY_V     = 0x56
	KEY_W     = 0x57
	KEY_X     = 0x58
	KEY_Y     = 0x59
	KEY_Z     = 0x5A
)

var (
	user32               = syscall.NewLazyDLL("user32.dll")
	procGetCursorPos     = user32.NewProc("GetCursorPos")
	procMouseEvent       = user32.NewProc("mouse_event")
	procGetAsyncKeyState = user32.NewProc("GetAsyncKeyState")
)

type POINT struct {
	X, Y int32
}

const (
	// clicks は連続でクリックする回数
	clicks = 1

	// interval はクリック間隔(ms)
	interval = 100

	// KEY はクリックするキー
	KEY = KEY_A
)

func main() {
	for {
		if isKeyPressed(KEY) {
			log.Println("Click!")
			for i := 0; i < clicks; i++ {
				clickMouse()
				time.Sleep(time.Duration(interval) * time.Millisecond)
			}
		}
	}
}

func isKeyPressed(KEY int) bool {
	ret, _, _ := procGetAsyncKeyState.Call(uintptr(KEY))
	return ret == KEY_PRESSED
}

func getCursorPos() POINT {
	var pt POINT
	ret, _, _ := procGetCursorPos.Call(uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		fmt.Println("Error getting cursor position")
	}
	return pt
}

func clickMouse() {
	pos := getCursorPos()
	procMouseEvent.Call(
		uintptr(MOUSEEVENTF_LEFTDOWN),
		uintptr(pos.X),
		uintptr(pos.Y),
		0,
		0,
	)
	procMouseEvent.Call(
		uintptr(MOUSEEVENTF_LEFTUP),
		uintptr(pos.X),
		uintptr(pos.Y),
		0,
		0,
	)
}
