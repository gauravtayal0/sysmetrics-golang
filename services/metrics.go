package services

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"net/http"
)

func Cpu(w http.ResponseWriter, r *http.Request) {
	v, _ := cpu.Percent(0, false)
	WriteResponse(w, NewAPIResponse(v))
}

func Memory(w http.ResponseWriter, r *http.Request) {

	v, _ := VMStat()
	WriteResponse(w, NewAPIResponse(v))
}

func Load(w http.ResponseWriter, r *http.Request) {

	v, _ := load.Avg()
	WriteResponse(w, NewAPIResponse(v))
}

func Host(w http.ResponseWriter, r *http.Request) {

	v, _ := host.Info()
	WriteResponse(w, NewAPIResponse(v))
}

func NetPro(w http.ResponseWriter, r *http.Request) {

	v, _ := NetProto()
	WriteResponse(w, NewAPIResponse(v))
}

func Net(w http.ResponseWriter, r *http.Request) {

	v, _ := NetIO()
	WriteResponse(w, NewAPIResponse(v))
}

func NetConn(w http.ResponseWriter, r *http.Request) {

	v, _ := NetConnections()
	WriteResponse(w, NewAPIResponse(v))
}

func Swap(w http.ResponseWriter, r *http.Request) {

	v, _ := SwapStat()
	WriteResponse(w, NewAPIResponse(v))
}
