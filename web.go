// Package gweb
// Support HTTP Server, restful api handler, with timeout setting
//
// Support RPC Server, but this only called by golang rpc client,
// not good enough
package gweb

import (
	. "github.com/yeqown/gweb/logger"
	. "github.com/yeqown/gweb/utils"

	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

var rpc_server = rpc.NewServer()

// StartHttpServer over HTTP as for api server
func StartHttpServer(servConfig *ServerConfig) {
	InitLogger(servConfig.Logpath)

	server := &http.Server{
		Addr: Fstring(":%d", servConfig.Port),
		Handler: http.TimeoutHandler(ApiHdl,
			5*time.Second,
			TimeoutJsonResp,
		),
	}

	AppL.Infof("Http Server listening on: %d\n", servConfig.Port)
	if err := server.ListenAndServe(); err != nil {
		AppL.Fatal(err.Error())
	}
}

// GetRpcServer get rcp.server instance to register
func GetRpcServer() *rpc.Server {
	return rpc_server
}

// StartRpcSerevr running a server to deal with rpc request
// default set jsonrpc
func StartRpcServer(rpcConfig *RpcServerConfig) {
	// DefaultRPCPath = "/_goRPC_"
	// DefaultDebugPath = "/debug/rpc"
	rpc_server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)

	l, err := net.Listen(rpcConfig.Network,
		Fstring("%s:%d",
			rpcConfig.Host,
			rpcConfig.Port,
		),
	)
	if err != nil {
		AppL.Fatal(err.Error())
	}

	AppL.Infof("Json-Rpc Listening on: %d\n", rpcConfig.Port)
	// loop listening
	for {
		conn, err := l.Accept()
		if err != nil {
			AppL.Fatal(err.Error())
		}

		AppL.Info("A new Rpc request received!")
		go rpc_server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
