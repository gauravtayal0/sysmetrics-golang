package services

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

func CPUTimes(perCPU, totalCPU bool) ([]cpu.TimesStat, error) {
	var cpuTimes []cpu.TimesStat
	if perCPU {
		if perCPUTimes, err := cpu.Times(true); err == nil {
			cpuTimes = append(cpuTimes, perCPUTimes...)
		} else {
			return nil, err
		}
	}
	if totalCPU {
		if totalCPUTimes, err := cpu.Times(false); err == nil {
			cpuTimes = append(cpuTimes, totalCPUTimes...)
		} else {
			return nil, err
		}
	}
	return cpuTimes, nil
}

func NetProto() ([]net.ProtoCountersStat, error) {
	return net.ProtoCounters(nil)
}

func NetIO() ([]net.IOCountersStat, error) {
	return net.IOCounters(true)
}

func NetConnections() ([]net.ConnectionStat, error) {
	return net.Connections("all")
}

func VMStat() (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemory()
}

func SwapStat() (*mem.SwapMemoryStat, error) {
	return mem.SwapMemory()
}
