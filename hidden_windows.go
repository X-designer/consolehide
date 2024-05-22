//go:build windows
// +build windows

package consolehide

import "syscall"

const (
	win_sw_hide = 0
	win_sw_show = 5
)

func HideConsole(hide bool) {
	getConsoleWindow := syscall.NewLazyDLL("kernel32.dll").NewProc("GetConsoleWindow")
	showWindow := syscall.NewLazyDLL("user32.dll").NewProc("ShowWindow")
	if getConsoleWindow.Find() == nil && showWindow.Find() == nil {
		hwnd, _, _ := getConsoleWindow.Call()
		if hwnd != 0 {
			if hide {
				showWindow.Call(hwnd, win_sw_hide)
			} else {
				showWindow.Call(hwnd, win_sw_show)
			}
		}
	}
}
