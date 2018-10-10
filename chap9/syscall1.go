package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

type userDll struct {
	dll *syscall.DLL
}

var (
	kernel32_dll           = loadDLL("kernel32.dll")
	procgetModuleFileName  = kernel32_dll.findProc("GetModuleFileNameA")
	procgetModuleFileNameW = kernel32_dll.findProc("GetModuleFileNameW")
)

func loadDLL(name string) *userDll {
	dll, err := syscall.LoadDLL(name)
	if err != nil {
		panic(err)
	}
	return &userDll{dll}
}

func (c *userDll) findProc(name string) *syscall.Proc {
	proc, err := c.dll.FindProc(name)
	if err != nil {
			panic(err)
	}
	return proc
}

func main() {
	const max_path int = 520

	fmt.Println("ASCII")
	var bufA [max_path]byte
	r1, r2, lastErr := procgetModuleFileName.Call(0, uintptr(unsafe.Pointer(&bufA)), (uintptr)(max_path))

	fmt.Printf("r1 %v, r2 %v, lastErr %v\n", r1, r2, lastErr)
	lastErr = syscall.GetLastError()
	if lastErr != nil {
		fmt.Println("windows error: ", lastErr)
	} else {
		fmt.Printf("[%s]\n", bufA[:r1])
	}

	fmt.Println("\nUnicode:")
	var bufW [max_path]byte

	r1, _, _ = procgetModuleFileNameW.Call(0, uintptr(unsafe.Pointer(&bufW)), (uintptr)(max_path))
	lastErr = syscall.GetLastError()
	if lastErr != nil {
		fmt.Println("window error: ", lastErr)
	} else {
		s := syscall.UTF16ToString((*[max_path]uint16)(unsafe.Pointer(&bufW[0]))[:r1])
		fmt.Printf("%s\n", s)
	}
}