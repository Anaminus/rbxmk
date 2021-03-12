// +build windows

package clipboard

import (
	"fmt"
	"unsafe"

	"github.com/anaminus/rbxmk/library/internal/winapi"
	"golang.org/x/sys/windows"
)

var mediaToFormat = map[string]uint32{
	"text/plain": winapi.CF_UNICODETEXT,
	"audio/wave": winapi.CF_WAVE,
	"audio/riff": winapi.CF_RIFF,
	"image/bmp":  winapi.CF_DIB,
	"image/tiff": winapi.CF_TIFF,
}

// Clear removes all data from the clipboard.
func Clear() error {
	if err := winapi.OpenClipboard(0); err != nil {
		return fmt.Errorf("OpenClipboard: %w", err)
	}
	defer winapi.CloseClipboard()

	if err := winapi.EmptyClipboard(); err != nil {
		return fmt.Errorf("EmptyClipboard: %w", err)
	}
	return nil
}

// Write sets data to the clipboard. If multiple formats are supported, then
// each given format is written according to the specified media type.
// Otherwise, which format is selected is implementation-defined.
//
// If no formats are given, then the clipboard is cleared with no other action.
func Write(formats []Format) (err error) {
	if err := winapi.OpenClipboard(0); err != nil {
		return fmt.Errorf("OpenClipboard: %w", err)
	}
	defer winapi.CloseClipboard()

	if err := winapi.EmptyClipboard(); err != nil {
		return fmt.Errorf("EmptyClipboard: %w", err)
	}

	for _, format := range formats {
		if len(format.Content) == 0 {
			continue
		}

		// Find format ID.
		formatID, ok := mediaToFormat[format.Name]
		if !ok {
			if formatID, err = winapi.RegisterClipboardFormat(format.Name); err != nil {
				return fmt.Errorf("RegisterClipboardFormat: %w", err)
			}
		}

		// Prepare content for copying.
		var content unsafe.Pointer
		var length uint
		switch formatID {
		case winapi.CF_UNICODETEXT:
			// Encode to UTF-16.
			u, err := windows.UTF16FromString(string(format.Content))
			if err != nil {
				return err
			}
			content = unsafe.Pointer(&u[0])
			length = uint(len(u) * 2)
		default:
			// Encode to raw bytes.
			content = unsafe.Pointer(&format.Content[0])
			length = uint(len(format.Content))
		}

		// Allocate and copy to memory for use by clipboard.
		const GMEM_MOVEABLE = 0x0002
		hMem, err := winapi.GlobalAlloc(GMEM_MOVEABLE, length)
		if err != nil {
			return fmt.Errorf("GlobalAlloc: %w", err)
		}
		p, err := winapi.GlobalLock(hMem)
		if err != nil {
			return fmt.Errorf("GlobalLock: %w", err)
		}
		winapi.RtlMoveMemory(p, content, length)
		if err := winapi.GlobalUnlock(hMem); err != nil {
			return fmt.Errorf("GlobalUnlock: %w", err)
		}
		if _, err := winapi.SetClipboardData(formatID, winapi.HANDLE(hMem)); err != nil {
			err = fmt.Errorf("SetClipboardData: %w", err)
			winapi.GlobalFree(hMem)
			return err
		}
	}
	return nil
}

// Read gets data from the clipboard. If multiple clipboard formats are
// supported, Read selects the first format that matches one of the given
// media types.
//
// Each argument is a media type (e.g. "text/plain").
//
// If an error is returned, then f will be less than 0. If no data was found,
// then the error will contain NoDataError. If no formats were given, then f
// will be less than 0, and err will be nil.
func Read(formats ...string) (f int, b []byte, err error) {
	if len(formats) == 0 {
		return -1, nil, nil
	}
	if err := winapi.OpenClipboard(0); err != nil {
		return -1, nil, fmt.Errorf("OpenClipboard: %w", err)
	}
	defer winapi.CloseClipboard()

	// Locate first given format that matches format in clipboard.
	var formatIndex int
	var formatID uint32
	var hMem winapi.HGLOBAL
	for i, format := range formats {
		var ok bool
		if formatID, ok = mediaToFormat[format]; !ok {
			if formatID, _ = winapi.RegisterClipboardFormat(format); formatID == 0 {
				continue
			}
		}
		h, _ := winapi.GetClipboardData(formatID)
		if h == 0 {
			continue
		}
		hMem = winapi.HGLOBAL(h)
		formatIndex = i
		break
	}
	if hMem == 0 {
		return -1, nil, NoDataError{}
	}

	// Copy from clipboard memory.
	p, err := winapi.GlobalLock(hMem)
	if err != nil {
		return -1, nil, fmt.Errorf("GlobalLock: %w", err)
	}
	defer winapi.GlobalUnlock(hMem)

	switch formatID {
	case winapi.CF_UNICODETEXT:
		// Decode from UTF-16.
		b = []byte(windows.UTF16PtrToString((*uint16)(p)))
	default:
		// Decode from raw bytes.
		length, err := winapi.GlobalSize(hMem)
		if err != nil {
			return -1, nil, fmt.Errorf("GlobalSize: %w", err)
		}

		//WARN
		type Slice struct {
			Data unsafe.Pointer
			Len  int
			Cap  int
		}
		var c []byte
		h := (*Slice)(unsafe.Pointer(&c))
		h.Data = unsafe.Pointer(p)
		h.Len = int(length)
		h.Cap = int(length)

		b = make([]byte, len(c))
		copy(b, c)
	}
	return formatIndex, b, nil
}
