package win

import "syscall"

var (
	user32DLL = syscall.NewLazyDLL("user32.dll")

	findWindow                 = user32DLL.NewProc("FindWindowW")
	getWindowText              = user32DLL.NewProc("GetWindowTextW")
	getWindowThreadProcessID   = user32DLL.NewProc("GetWindowThreadProcessId")
	getWindowRect              = user32DLL.NewProc("GetWindowRect")
	setWindowPos               = user32DLL.NewProc("SetWindowPos")
	setLayeredWindowAttributes = user32DLL.NewProc("SetLayeredWindowAttributes")
	setWindowLong              = user32DLL.NewProc("SetWindowLongW")
	getWindowLong              = user32DLL.NewProc("GetWindowLongW")

	postMessage  = user32DLL.NewProc("PostMessageW")
	setCursorPos = user32DLL.NewProc("SetCursorPos")
	mouseEvent   = user32DLL.NewProc("mouse_event")
)

var (
	setForegroundWin    = user32DLL.NewProc("SetForegroundWindow")
	getWindowLongPtr    = user32DLL.NewProc("GetWindowLongPtrW")
	showWindow          = user32DLL.NewProc("ShowWindow")
	setForegroundWindow = user32DLL.NewProc("SetForegroundWindow")
)
