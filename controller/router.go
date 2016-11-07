package controller

import (
	"github.com/moonfrog/sysutil/services"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Get Cpu",
		"GET",
		"/metrics/cpu",
		services.Cpu,
	},
	Route{
		"Get Memory",
		"GET",
		"/metrics/mem",
		services.Memory,
	},
	Route{
		"Get Load",
		"GET",
		"/metrics/load",
		services.Load,
	},
	Route{
		"Get Host Info",
		"GET",
		"/metrics/host",
		services.Host,
	},
	Route{
		"Get Net connections",
		"GET",
		"/metrics/nc",
		services.NetConn,
	},
	Route{
		"Get Net Protocol",
		"GET",
		"/metrics/np",
		services.NetPro,
	},
	Route{
		"Get network IO",
		"GET",
		"/metrics/ni",
		services.Net,
	},
	Route{
		"Get swap stats",
		"GET",
		"/metrics/swap",
		services.Swap,
	},
}
