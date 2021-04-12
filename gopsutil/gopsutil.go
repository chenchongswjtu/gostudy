package main

//获取硬件信息

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	//"github.com/shirou/gopsutil/load"
)

//获取cpu信息
func check_cpu() {
	c, _ := cpu.Counts(false)
	fmt.Println("cpu核数：", c)

	c1, _ := cpu.Info()
	fmt.Println("cpu型号：", c1[0].ModelName)
}

//获取内存信息
func check_mem() {
	v, _ := mem.VirtualMemory()
	fmt.Println("总大小：", v.Total/1024/1024, "M")
	fmt.Println("可用：", v.Available/1024/1024, "M")
	fmt.Println("被使用：", v.Used/1024/1024, "M")
	fmt.Println()
}

//获取硬盘信息
func check_disk() {
	p, _ := disk.Partitions(false)
	var coun uint64
	for i := 0; i < len(p); i++ {
		co, _ := disk.Usage(p[i].Mountpoint)
		if co.Total < 1024000000 {
			fmt.Println("分区", p[i].Mountpoint, "总容量", co.Total/1024/1024, "M")
		} else {
			fmt.Println("分区", p[i].Mountpoint, "总容量", co.Total/1024/1024/1024, "G")
		}
		coun += co.Total
	}
	fmt.Println("硬盘总容量：", coun/1024/1024/1024, "G")
}

//获取网卡信息
func check_net() {
	inter, _ := net.Interfaces()

	for i := 0; i < len(inter); i++ {
		fmt.Println("网卡名：", inter[i].Name)
		if inter[i].HardwareAddr == "" {
			fmt.Println("MAC：", "无")
		} else {
			fmt.Println("MAC：", inter[i].HardwareAddr)
		}
		//fmt.Println("IP：", inter[i].Addrs)
		var ip_str string
		for _, ip := range inter[i].Addrs {
			ip_str += ip.Addr + "   "
		}
		fmt.Println("IP:", ip_str)
	}
}

func main() {

	fmt.Println("--------------cpu------------------")
	check_cpu()
	fmt.Println()

	fmt.Println("--------------memory---------------")
	check_mem()
	fmt.Println()

	fmt.Println("--------------disk-----------------")
	check_disk()
	fmt.Println()

	fmt.Println("--------------net------------------")
	check_net()
	fmt.Println()
}
