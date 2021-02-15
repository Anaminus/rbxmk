// +build windows

package winapi

import (
	"errors"
	"unsafe"

	"golang.org/x/sys/windows"
)

type (
	UINT    = uint32
	SIZE_T  = uint
	HANDLE  = windows.Handle
	HWND    = windows.HWND
	HGLOBAL uintptr
)

const NULL = 0

const GMEM_MOVEABLE = 0x0002

const (
	CF_TEXT            = 1
	CF_BITMAP          = 2
	CF_METAFILEPICT    = 3
	CF_SYLK            = 4
	CF_DIF             = 5
	CF_TIFF            = 6
	CF_OEMTEXT         = 7
	CF_DIB             = 8
	CF_PALETTE         = 9
	CF_PENDATA         = 10
	CF_RIFF            = 11
	CF_WAVE            = 12
	CF_UNICODETEXT     = 13
	CF_ENHMETAFILE     = 14
	CF_HDROP           = 15
	CF_LOCALE          = 16
	CF_DIBV5           = 17
	CF_OWNERDISPLAY    = 0x0080
	CF_DSPTEXT         = 0x0081
	CF_DSPBITMAP       = 0x0082
	CF_DSPMETAFILEPICT = 0x0083
	CF_DSPENHMETAFILE  = 0x008E
	CF_PRIVATEFIRST    = 0x0200
	CF_PRIVATELAST     = 0x02FF
	CF_GDIOBJFIRST     = 0x0300
	CF_GDIOBJLAST      = 0x03FF
)

var (
	kernel32      = windows.NewLazySystemDLL("kernel32.dll")
	globalAlloc   = kernel32.NewProc("GlobalAlloc")
	globalFree    = kernel32.NewProc("GlobalFree")
	globalLock    = kernel32.NewProc("GlobalLock")
	globalSize    = kernel32.NewProc("GlobalSize")
	globalUnlock  = kernel32.NewProc("GlobalUnlock")
	rtlMoveMemory = kernel32.NewProc("RtlMoveMemory")
)

func GlobalAlloc(uFlags UINT, dwBytes SIZE_T) (HGLOBAL, error) {
	r, _, err := globalAlloc.Call(uintptr(uFlags), uintptr(dwBytes))
	if r == NULL {
		return 0, err
	}
	return HGLOBAL(r), nil
}

func GlobalFree(hMem HGLOBAL) (HGLOBAL, error) {
	r, _, err := globalFree.Call(uintptr(hMem))
	if r != NULL {
		return HGLOBAL(r), err
	}
	return 0, nil
}

func GlobalLock(hMem HGLOBAL) (unsafe.Pointer, error) {
	r, _, err := globalLock.Call(uintptr(hMem))
	if r == NULL {
		return nil, err
	}
	return unsafe.Pointer(r), nil
}

func GlobalSize(hMem HGLOBAL) (SIZE_T, error) {
	r, _, err := globalSize.Call(uintptr(hMem))
	if r == 0 {
		return 0, err
	}
	return SIZE_T(r), nil
}

func GlobalUnlock(hMem HGLOBAL) error {
	r, _, err := globalUnlock.Call(uintptr(hMem))
	if r == 0 {
		if en := windows.Errno(0); errors.As(err, &en) && en != windows.NO_ERROR {
			return err
		}
	}
	return nil
}

func RtlMoveMemory(dst, src unsafe.Pointer, length SIZE_T) {
	rtlMoveMemory.Call(uintptr(unsafe.Pointer(dst)), uintptr(src), uintptr(length))
}

var (
	user32                   = windows.NewLazySystemDLL("user32.dll")
	closeClipboard           = user32.NewProc("CloseClipboard")
	emptyClipboard           = user32.NewProc("EmptyClipboard")
	getClipboardData         = user32.NewProc("GetClipboardData")
	openClipboard            = user32.NewProc("OpenClipboard")
	registerClipboardFormatW = user32.NewProc("RegisterClipboardFormatW")
	setClipboardData         = user32.NewProc("SetClipboardData")
)

func CloseClipboard() error {
	r, _, err := closeClipboard.Call()
	if r == 0 {
		return err
	}
	return nil
}

func EmptyClipboard() error {
	r, _, err := emptyClipboard.Call()
	if r == 0 {
		return err
	}
	return nil
}

func GetClipboardData(format UINT) (HANDLE, error) {
	r, _, err := getClipboardData.Call(uintptr(format))
	if r == NULL {
		return 0, err
	}
	return HANDLE(r), nil
}

func OpenClipboard(hWndNewOwner HWND) error {
	r, _, err := openClipboard.Call(uintptr(hWndNewOwner))
	if r == 0 {
		return err
	}
	return nil
}

func RegisterClipboardFormat(lpszFormat string) (UINT, error) {
	s, err := windows.UTF16PtrFromString(lpszFormat)
	if err != nil {
		return 0, err
	}
	r, _, err := registerClipboardFormatW.Call(uintptr(unsafe.Pointer(s)))
	if r == 0 {
		return 0, err
	}
	return UINT(r), nil
}

func SetClipboardData(uFormat UINT, hMem HANDLE) (HANDLE, error) {
	r, _, err := setClipboardData.Call(uintptr(uFormat), uintptr(hMem))
	if r == NULL {
		return 0, err
	}
	return HANDLE(r), nil
}
