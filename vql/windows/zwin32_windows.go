// MACHINE GENERATED BY 'go generate' COMMAND; DO NOT EDIT

package windows

import (
	"syscall"
	"unsafe"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modntdll    = NewLazySystemDLL("ntdll.dll")
	modkernel32 = NewLazySystemDLL("kernel32.dll")
	modpsapi    = NewLazySystemDLL("psapi.dll")
	modAdvapi32 = NewLazySystemDLL("Advapi32.dll")
	modnetapi32 = NewLazySystemDLL("netapi32.dll")
	modwintrust = NewLazySystemDLL("wintrust.dll")

	procNtOpenThreadToken                    = modntdll.NewProc("NtOpenThreadToken")
	procReadProcessMemory                    = modkernel32.NewProc("ReadProcessMemory")
	procGetProcessMemoryInfo                 = modpsapi.NewProc("GetProcessMemoryInfo")
	procGetProcessIoCounters                 = modkernel32.NewProc("GetProcessIoCounters")
	procQueryFullProcessImageNameW           = modkernel32.NewProc("QueryFullProcessImageNameW")
	procNtOpenDirectoryObject                = modntdll.NewProc("NtOpenDirectoryObject")
	procAdjustTokenPrivileges                = modAdvapi32.NewProc("AdjustTokenPrivileges")
	procLookupPrivilegeValueW                = modAdvapi32.NewProc("LookupPrivilegeValueW")
	procNtDuplicateObject                    = modntdll.NewProc("NtDuplicateObject")
	procNtQueryInformationProcess            = modntdll.NewProc("NtQueryInformationProcess")
	procNtQueryInformationThread             = modntdll.NewProc("NtQueryInformationThread")
	procNtQueryObject                        = modntdll.NewProc("NtQueryObject")
	procNtQuerySystemInformation             = modntdll.NewProc("NtQuerySystemInformation")
	procCloseHandle                          = modkernel32.NewProc("CloseHandle")
	procOpenProcess                          = modkernel32.NewProc("OpenProcess")
	procGetSystemInfo                        = modkernel32.NewProc("GetSystemInfo")
	procModule32NextW                        = modkernel32.NewProc("Module32NextW")
	procModule32FirstW                       = modkernel32.NewProc("Module32FirstW")
	procCreateToolhelp32Snapshot             = modkernel32.NewProc("CreateToolhelp32Snapshot")
	procGetMappedFileNameW                   = modpsapi.NewProc("GetMappedFileNameW")
	procVirtualQueryEx                       = modkernel32.NewProc("VirtualQueryEx")
	procNetApiBufferFree                     = modnetapi32.NewProc("NetApiBufferFree")
	procNetUserEnum                          = modnetapi32.NewProc("NetUserEnum")
	procNetUserGetGroups                     = modnetapi32.NewProc("NetUserGetGroups")
	procCryptCATAdminAcquireContext2         = modwintrust.NewProc("CryptCATAdminAcquireContext2")
	procCryptCATAdminReleaseContext          = modwintrust.NewProc("CryptCATAdminReleaseContext")
	procCryptCATAdminCalcHashFromFileHandle2 = modwintrust.NewProc("CryptCATAdminCalcHashFromFileHandle2")
	procCryptCATAdminEnumCatalogFromHash     = modwintrust.NewProc("CryptCATAdminEnumCatalogFromHash")
	procCryptCATCatalogInfoFromContext       = modwintrust.NewProc("CryptCATCatalogInfoFromContext")
	procCryptCATAdminReleaseCatalogContext   = modwintrust.NewProc("CryptCATAdminReleaseCatalogContext")
	procWinVerifyTrust                       = modwintrust.NewProc("WinVerifyTrust")
	procWTHelperProvDataFromStateData        = modwintrust.NewProc("WTHelperProvDataFromStateData")
	procWTHelperGetProvSignerFromChain       = modwintrust.NewProc("WTHelperGetProvSignerFromChain")
)

func NtOpenThreadToken(thread_handle syscall.Handle, DesiredAccess uint32, open_as_self bool, token_handle *syscall.Handle) (status uint32) {
	var _p0 uint32
	if open_as_self {
		_p0 = 1
	} else {
		_p0 = 0
	}
	r0, _, _ := syscall.Syscall6(procNtOpenThreadToken.Addr(), 4, uintptr(thread_handle), uintptr(DesiredAccess), uintptr(_p0), uintptr(unsafe.Pointer(token_handle)), 0, 0)
	status = uint32(r0)
	return
}

func _ReadProcessMemory(handle syscall.Handle, baseAddress uintptr, buffer uintptr, size uintptr, numRead *uintptr) (err error) {
	r1, _, e1 := syscall.Syscall6(procReadProcessMemory.Addr(), 5, uintptr(handle), uintptr(baseAddress), uintptr(buffer), uintptr(size), uintptr(unsafe.Pointer(numRead)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func GetProcessMemoryInfo(handle syscall.Handle, memCounters *PROCESS_MEMORY_COUNTERS, cb uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetProcessMemoryInfo.Addr(), 3, uintptr(handle), uintptr(unsafe.Pointer(memCounters)), uintptr(cb))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func GetProcessIoCounters(hProcess syscall.Handle, lpIoCounters *IO_COUNTERS) (ok bool) {
	r0, _, _ := syscall.Syscall(procGetProcessIoCounters.Addr(), 2, uintptr(hProcess), uintptr(unsafe.Pointer(lpIoCounters)), 0)
	ok = r0 != 0
	return
}

func QueryFullProcessImageName(handle syscall.Handle, dwFlags uint32, buffer *byte, length *uint32) (ok bool) {
	r0, _, _ := syscall.Syscall6(procQueryFullProcessImageNameW.Addr(), 4, uintptr(handle), uintptr(dwFlags), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(length)), 0, 0)
	ok = r0 != 0
	return
}

func NtOpenDirectoryObject(DirectoryHandle *uint32, DesiredAccess uint32, ObjectAttributes *OBJECT_ATTRIBUTES) (status uint32) {
	r0, _, _ := syscall.Syscall(procNtOpenDirectoryObject.Addr(), 3, uintptr(unsafe.Pointer(DirectoryHandle)), uintptr(DesiredAccess), uintptr(unsafe.Pointer(ObjectAttributes)))
	status = uint32(r0)
	return
}

func AdjustTokenPrivileges(TokenHandle syscall.Token, DisableAllPrivileges bool, NewState uintptr, BufferLength int, PreviousState uintptr, ReturnLength *int) (err error) {
	var _p0 uint32
	if DisableAllPrivileges {
		_p0 = 1
	} else {
		_p0 = 0
	}
	r1, _, e1 := syscall.Syscall6(procAdjustTokenPrivileges.Addr(), 6, uintptr(TokenHandle), uintptr(_p0), uintptr(NewState), uintptr(BufferLength), uintptr(PreviousState), uintptr(unsafe.Pointer(ReturnLength)))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func LookupPrivilegeValue(lpSystemName uintptr, lpName uintptr, out uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procLookupPrivilegeValueW.Addr(), 3, uintptr(lpSystemName), uintptr(lpName), uintptr(out))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func NtDuplicateObject(SourceProcessHandle syscall.Handle, SourceHandle syscall.Handle, TargetProcessHandle syscall.Handle, TargetHandle *syscall.Handle, DesiredAccess uint32, InheritHandle uint32, Options uint32) (status uint32) {
	r0, _, _ := syscall.Syscall9(procNtDuplicateObject.Addr(), 7, uintptr(SourceProcessHandle), uintptr(SourceHandle), uintptr(TargetProcessHandle), uintptr(unsafe.Pointer(TargetHandle)), uintptr(DesiredAccess), uintptr(InheritHandle), uintptr(Options), 0, 0)
	status = uint32(r0)
	return
}

func NtQueryInformationProcess(Handle syscall.Handle, ObjectInformationClass uint32, ProcessInformation *byte, ProcessInformationLength uint32, ReturnLength *uint32) (status uint32) {
	r0, _, _ := syscall.Syscall6(procNtQueryInformationProcess.Addr(), 5, uintptr(Handle), uintptr(ObjectInformationClass), uintptr(unsafe.Pointer(ProcessInformation)), uintptr(ProcessInformationLength), uintptr(unsafe.Pointer(ReturnLength)), 0)
	status = uint32(r0)
	return
}

func NtQueryInformationThread(Handle syscall.Handle, ObjectInformationClass uint32, ThreadInformation *byte, ThreadInformationLength uint32, ReturnLength *uint32) (status uint32, err error) {
	r0, _, e1 := syscall.Syscall6(procNtQueryInformationThread.Addr(), 5, uintptr(Handle), uintptr(ObjectInformationClass), uintptr(unsafe.Pointer(ThreadInformation)), uintptr(ThreadInformationLength), uintptr(unsafe.Pointer(ReturnLength)), 0)
	status = uint32(r0)
	if status == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func NtQueryObject(Handle syscall.Handle, ObjectInformationClass uint32, ObjectInformation *byte, ObjectInformationLength uint32, ReturnLength *uint32) (status uint32, err error) {
	r0, _, e1 := syscall.Syscall6(procNtQueryObject.Addr(), 5, uintptr(Handle), uintptr(ObjectInformationClass), uintptr(unsafe.Pointer(ObjectInformation)), uintptr(ObjectInformationLength), uintptr(unsafe.Pointer(ReturnLength)), 0)
	status = uint32(r0)
	if status == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func NtQuerySystemInformation(SystemInformationClass uint32, SystemInformation *byte, SystemInformationLength uint32, ReturnLength *uint32) (status uint32) {
	r0, _, _ := syscall.Syscall6(procNtQuerySystemInformation.Addr(), 4, uintptr(SystemInformationClass), uintptr(unsafe.Pointer(SystemInformation)), uintptr(SystemInformationLength), uintptr(unsafe.Pointer(ReturnLength)), 0, 0)
	status = uint32(r0)
	return
}

func CloseHandle(h syscall.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procCloseHandle.Addr(), 1, uintptr(h), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func OpenProcess(dwDesiredAccess uint32, bInheritHandle bool, dwProcessId uint32) (handle syscall.Handle, err error) {
	var _p0 uint32
	if bInheritHandle {
		_p0 = 1
	} else {
		_p0 = 0
	}
	r0, _, e1 := syscall.Syscall(procOpenProcess.Addr(), 3, uintptr(dwDesiredAccess), uintptr(_p0), uintptr(dwProcessId))
	handle = syscall.Handle(r0)
	if handle == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func GetSystemInfo(lpSystemInfo *SYSTEM_INFO) (err error) {
	r1, _, e1 := syscall.Syscall(procGetSystemInfo.Addr(), 1, uintptr(unsafe.Pointer(lpSystemInfo)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func Module32Next(hSnapshot syscall.Handle, me *MODULEENTRY32W) (err error) {
	r1, _, e1 := syscall.Syscall(procModule32NextW.Addr(), 2, uintptr(hSnapshot), uintptr(unsafe.Pointer(me)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func Module32First(hSnapshot syscall.Handle, me *MODULEENTRY32W) (err error) {
	r1, _, e1 := syscall.Syscall(procModule32FirstW.Addr(), 2, uintptr(hSnapshot), uintptr(unsafe.Pointer(me)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func CreateToolhelp32Snapshot(dwFlags uint32, th32ProcessID uint32) (handle syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procCreateToolhelp32Snapshot.Addr(), 2, uintptr(dwFlags), uintptr(th32ProcessID), 0)
	handle = syscall.Handle(r0)
	if handle == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func GetMappedFileNameW(hProcess syscall.Handle, address uint64, lpFilename *uint16, nSize uint32) (len uint32, err error) {
	r0, _, e1 := syscall.Syscall6(procGetMappedFileNameW.Addr(), 4, uintptr(hProcess), uintptr(address), uintptr(unsafe.Pointer(lpFilename)), uintptr(nSize), 0, 0)
	len = uint32(r0)
	if len == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func VirtualQueryEx(handle syscall.Handle, address uint64, info *MEMORY_BASIC_INFORMATION, info_size uintptr) (size int32, err error) {
	r0, _, e1 := syscall.Syscall6(procVirtualQueryEx.Addr(), 4, uintptr(handle), uintptr(address), uintptr(unsafe.Pointer(info)), uintptr(info_size), 0, 0)
	size = int32(r0)
	if size == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func NetApiBufferFree(Buffer uintptr) (status NET_API_STATUS) {
	r0, _, _ := syscall.Syscall(procNetApiBufferFree.Addr(), 1, uintptr(Buffer), 0, 0)
	status = NET_API_STATUS(r0)
	return
}

func NetUserEnum(servername *uint16, level uint32, filter uint32, bufptr *uintptr, prefmaxlen uint32, entriesread *uint32, totalentries *uint32, resume_handle *uint32) (status NET_API_STATUS) {
	r0, _, _ := syscall.Syscall9(procNetUserEnum.Addr(), 8, uintptr(unsafe.Pointer(servername)), uintptr(level), uintptr(filter), uintptr(unsafe.Pointer(bufptr)), uintptr(prefmaxlen), uintptr(unsafe.Pointer(entriesread)), uintptr(unsafe.Pointer(totalentries)), uintptr(unsafe.Pointer(resume_handle)), 0)
	status = NET_API_STATUS(r0)
	return
}

func NetUserGetGroups(servername *LPCWSTR, username *LPCWSTR, level DWORD, bufptr *LPBYTE, prefmaxlen DWORD, entriesread *LPDWORD, totalentries *LPDWORD) (status NET_API_STATUS) {
	r0, _, _ := syscall.Syscall9(procNetUserGetGroups.Addr(), 7, uintptr(unsafe.Pointer(servername)), uintptr(unsafe.Pointer(username)), uintptr(level), uintptr(unsafe.Pointer(bufptr)), uintptr(prefmaxlen), uintptr(unsafe.Pointer(entriesread)), uintptr(unsafe.Pointer(totalentries)), 0, 0)
	status = NET_API_STATUS(r0)
	return
}

func CryptCATAdminAcquireContext2(handle *syscall.Handle, pgSubsystem *GUID, pwszHashAlgorithm *byte, pStrongHashPolicy *byte, dwFlags uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procCryptCATAdminAcquireContext2.Addr(), 5, uintptr(unsafe.Pointer(handle)), uintptr(unsafe.Pointer(pgSubsystem)), uintptr(unsafe.Pointer(pwszHashAlgorithm)), uintptr(unsafe.Pointer(pStrongHashPolicy)), uintptr(dwFlags), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func CryptCATAdminReleaseContext(handle syscall.Handle, unused int32) (ok bool) {
	r0, _, _ := syscall.Syscall(procCryptCATAdminReleaseContext.Addr(), 2, uintptr(handle), uintptr(unused), 0)
	ok = r0 != 0
	return
}

func CryptCATAdminCalcHashFromFileHandle2(handle syscall.Handle, fd uintptr, pcbHash *uint32, pbHash *byte, dwFlags uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procCryptCATAdminCalcHashFromFileHandle2.Addr(), 5, uintptr(handle), uintptr(fd), uintptr(unsafe.Pointer(pcbHash)), uintptr(unsafe.Pointer(pbHash)), uintptr(dwFlags), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func CryptCATAdminEnumCatalogFromHash(handle syscall.Handle, pbHash *byte, pcbHash uint32, dwFlags uint32, phPrevCatInfo *syscall.Handle) (HCATINFO syscall.Handle) {
	r0, _, _ := syscall.Syscall6(procCryptCATAdminEnumCatalogFromHash.Addr(), 5, uintptr(handle), uintptr(unsafe.Pointer(pbHash)), uintptr(pcbHash), uintptr(dwFlags), uintptr(unsafe.Pointer(phPrevCatInfo)), 0)
	HCATINFO = syscall.Handle(r0)
	return
}

func CryptCATCatalogInfoFromContext(handle syscall.Handle, psCatInfo *CATALOG_INFO, dwFlags uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procCryptCATCatalogInfoFromContext.Addr(), 3, uintptr(handle), uintptr(unsafe.Pointer(psCatInfo)), uintptr(dwFlags))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func CryptCATAdminReleaseCatalogContext(handle syscall.Handle, handle2 syscall.Handle, dwFlags uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procCryptCATAdminReleaseCatalogContext.Addr(), 3, uintptr(handle), uintptr(handle2), uintptr(dwFlags))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func WinVerifyTrust(handle syscall.Handle, action *GUID, data *WINTRUST_DATA) (ret uint32, err error) {
	r0, _, e1 := syscall.Syscall(procWinVerifyTrust.Addr(), 3, uintptr(handle), uintptr(unsafe.Pointer(action)), uintptr(unsafe.Pointer(data)))
	ret = uint32(r0)
	if ret != 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func WTHelperProvDataFromStateData(handle syscall.Handle) (provider *CRYPT_PROVIDER_DATA, err error) {
	r0, _, e1 := syscall.Syscall(procWTHelperProvDataFromStateData.Addr(), 1, uintptr(handle), 0, 0)
	provider = (*CRYPT_PROVIDER_DATA)(unsafe.Pointer(r0))
	if provider == nil {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func WTHelperGetProvSignerFromChain(pProvData *CRYPT_PROVIDER_DATA, idxSigner uint32, fCounterSigner bool, idxCounterSigner uint32) (signer *CRYPT_PROVIDER_SGNR, err error) {
	var _p0 uint32
	if fCounterSigner {
		_p0 = 1
	} else {
		_p0 = 0
	}
	r0, _, e1 := syscall.Syscall6(procWTHelperGetProvSignerFromChain.Addr(), 4, uintptr(unsafe.Pointer(pProvData)), uintptr(idxSigner), uintptr(_p0), uintptr(idxCounterSigner), 0, 0)
	signer = (*CRYPT_PROVIDER_SGNR)(unsafe.Pointer(r0))
	if signer == nil {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}
