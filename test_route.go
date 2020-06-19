package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"syscall"
	"unsafe"
)

func GetLocalIp() (string, error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}

		}
	}
	return "", errors.New("can not GetLocalIp")
}

func Inet_ntoa(ipnr uint32, isBig bool) string {
	ip := net.IPv4(0, 0, 0, 0)
	var bo binary.ByteOrder
	if isBig {
		bo = binary.BigEndian
	} else {
		bo = binary.LittleEndian
	}
	bo.PutUint32([]byte(ip.To4()), ipnr)
	return ip.String()
}

func Inet_aton(ip string, isBig bool) uint32 {
	var bo binary.ByteOrder
	if isBig {
		bo = binary.BigEndian
	} else {
		bo = binary.LittleEndian
	}
	return bo.Uint32(
		[]byte(net.ParseIP(ip).To4()),
	)
}

type RouteRow struct {
	dwForwardDest      uint32
	dwForwardMask      uint32
	dwForwardPolicy    uint32
	dwForwardNextHop   uint32
	dwForwardIfIndex   uint32
	dwForwardType      uint32
	dwForwardProto     uint32
	dwForwardAge       uint32
	dwForwardNextHopAS uint32
	dwForwardMetric1   uint32
	dwForwardMetric2   uint32
	dwForwardMetric3   uint32
	dwForwardMetric4   uint32
	dwForwardMetric5   uint32
}

type SliceHeader struct {
	Addr uintptr
	Len  int
	Cap  int
}

type DynamicMemory struct {
	mem []byte
}

func NewDynamicMemory(bytes uint32) *DynamicMemory {
	return &DynamicMemory{
		mem: make([]byte, bytes, bytes),
	}
}

func (this *DynamicMemory) Len() uint32 {
	return uint32(len(this.mem))
}

func (this *DynamicMemory) Address() uintptr {
	return (*SliceHeader)(unsafe.Pointer(&this.mem)).Addr
}

type RouteTable struct {
	dll                  *syscall.DLL
	getIpForwardTable    *syscall.Proc
	createIpForwardEntry *syscall.Proc
	deleteIpForwardEntry *syscall.Proc
}

func NewRouteTable() (*RouteTable, error) {
	dll, err := syscall.LoadDLL("iphlpapi.dll")
	if err != nil {
		return nil, err
	}

	getIpForwardTable, err := dll.FindProc("GetIpForwardTable")
	if err != nil {
		return nil, err
	}

	createIpForwardEntry, err := dll.FindProc("CreateIpForwardEntry")
	if err != nil {
		return nil, err
	}

	deleteIpForwardEntry, err := dll.FindProc("DeleteIpForwardEntry")
	if err != nil {
		return nil, err
	}

	return &RouteTable{
		dll:                  dll,
		getIpForwardTable:    getIpForwardTable,
		createIpForwardEntry: createIpForwardEntry,
		deleteIpForwardEntry: deleteIpForwardEntry,
	}, nil
}

func (this *RouteTable) Close() error {
	return this.dll.Release()
}

/*
https://msdn.microsoft.com/en-us/library/windows/desktop/aa366852(v=vs.85).aspx
typedef struct _MIB_IPFORWARDTABLE {
  DWORD            dwNumEntries;
  MIB_IPFORWARDROW table[ANY_SIZE];
} MIB_IPFORWARDTABLE, *PMIB_IPFORWARDTABLE;
*/
func (this *RouteTable) Routes() ([]RouteRow, error) {
	mem := NewDynamicMemory(
		uint32(
			4 + unsafe.Sizeof(RouteRow{}),
		),
	)
	table_size := uint32(0)
	_, r2, err := this.getIpForwardTable.Call(
		mem.Address(),
		uintptr(unsafe.Pointer(&table_size)),
		0,
	)
	// msdn https://msdn.microsoft.com/en-us/library/windows/desktop/aa365953(v=vs.85).aspx
	if r2 != 0 {
		return nil, err
	}

	mem = NewDynamicMemory(table_size)
	_, r2, err = this.getIpForwardTable.Call(
		mem.Address(),
		uintptr(unsafe.Pointer(&table_size)),
		0,
	)
	if r2 != 0 {
		return nil, err
	}

	num := *(*uint32)(unsafe.Pointer(mem.Address()))

	rows := []RouteRow{}
	sh_rows := (*SliceHeader)(unsafe.Pointer(&rows))
	sh_rows.Addr = mem.Address() + 4
	sh_rows.Len = int(num)
	sh_rows.Cap = int(num)
	return rows, nil
}

func (this *RouteTable) AddRoute(rr RouteRow) error {
	// https://msdn.microsoft.com/en-us/library/windows/desktop/aa365860(v=vs.85).aspx
	_, r2, err := this.createIpForwardEntry.Call(uintptr(unsafe.Pointer(&rr)))
	if r2 != 0 {
		return err
	}
	fmt.Println(err)
	return nil
}

func (this *RouteTable) DeleteRoute(rr RouteRow) error {
	// https://msdn.microsoft.com/en-us/library/windows/desktop/aa365860(v=vs.85).aspx
	_, r2, err := this.deleteIpForwardEntry.Call(uintptr(unsafe.Pointer(&rr)))
	if r2 != 0 {
		return err
	}
	fmt.Println(err)
	return nil
}

func main() {
	table, err := NewRouteTable()
	if err != nil {
		panic(err.Error())
	}
	defer table.Close()

	rows, err := table.Routes()
	if err != nil {
		panic(err.Error())
	}

	for _, row := range rows {
		fmt.Sprintf(
			"%v--->%v---->%v",
			Inet_ntoa(row.dwForwardDest, false),
			Inet_ntoa(row.dwForwardMask, false),
			Inet_ntoa(row.dwForwardNextHop, false),
		)
	}

	local_ip, err := GetLocalIp()
	if err != nil {
		panic(err.Error())
	}

	local_ip_uint := Inet_aton(local_ip, false)

	for _, row := range rows {
		// fmt.Println(Inet_ntoa(row.ForwardDest, false))
		if row.dwForwardNextHop == local_ip_uint {
			fmt.Println(row)
			// route add 14.215.177.37 mask 255.255.255.255 local_ip_uint
			row.dwForwardDest = Inet_aton("14.215.177.37", false)
			row.dwForwardMask = Inet_aton("255.255.255.255", false)
			row.dwForwardAge = 0
			row.dwForwardNextHopAS = 0
			fmt.Println(row)
			if err := table.AddRoute(row); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(Inet_ntoa(row.dwForwardDest, false))
				/* if err := table.DeleteRoute(row); err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(Inet_ntoa(row.ForwardDest, false))
				} */
			}
			break
		}
	}
}
