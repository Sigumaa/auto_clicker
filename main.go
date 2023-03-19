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
	KEY_SHIFT            = 0x10
	KEY_PRESSED          = 0x8000
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

func main() {
	// シフトキーが押されているときだけ
	// クリックを行う
	clicks := 100
	interval := 100
	for {
		if isShiftPressed() {
			log.Println("Click!")
			for i := 0; i < clicks; i++ {
				clickMouse()
				time.Sleep(time.Duration(interval) * time.Millisecond)
			}
		}
	}

}

func isShiftPressed() bool {
	// 0x8000: キーが押されている
	// 0x0000: キーが押されていない
	ret, _, _ := procGetAsyncKeyState.Call(uintptr(KEY_SHIFT))
	if ret == KEY_PRESSED {
		return true
	}
	return false
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
