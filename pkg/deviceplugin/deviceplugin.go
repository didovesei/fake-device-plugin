package deviceplugin

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/didovesei/fake-device-plugin/api"
	"google.golang.org/grpc"
	k8sdp "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
)

type DevicePlugin struct {
	server        *grpc.Server
	stop          chan api.Empty
	sokcetAddress string
}

func (dp *DevicePlugin) Start() error {
	if dp.server != nil {
		return fmt.Errorf("gRPC server already started")
	}

	listener, err := net.Listen("unix", dp.sokcetAddress)
	if err != nil {
		fmt.Errorf("could not create gRPC server socket: %w", err)
	}

	dp.server = grpc.NewServer()
	k8sdp.RegisterDevicePluginServer(dp.server, dp)

	go dp.server.Serve(listener)

	err = waitForGrpcServer(dp.socketPath, connectionTimeout)
	if err != nil {
		// this err is returned at the end of the Start function
		log.Printf("[%s] Error connecting to GRPC server: %v", dp.deviceName, err)
	}

	err = dp.Register()
	if err != nil {
		log.Printf("[%s] Error registering with device plugin manager: %v", dp.deviceName, err)
		return err
	}

	go dp.healthCheck()

	log.Println(dp.deviceName + " Device plugin server ready")

	return err
}

func (dp *DevicePlugin) GetDevicePluginOptions(context.Context, *k8sdp.Empty) (*k8sdp.DevicePluginOptions, error) {
	return nil, nil
}

func (dp *DevicePlugin) ListAndWatch(*k8sdp.Empty, k8sdp.DevicePlugin_ListAndWatchServer) error {
	return nil
}

func (dp *DevicePlugin) GetPreferredAllocation(context.Context, *k8sdp.PreferredAllocationRequest) (*k8sdp.PreferredAllocationResponse, error) {
	return nil, nil
}

func (dp *DevicePlugin) Allocate(context.Context, *k8sdp.AllocateRequest) (*k8sdp.AllocateResponse, error) {
	return nil, nil
}

func (dp *DevicePlugin) PreStartContainer(context.Context, *k8sdp.PreStartContainerRequest) (*k8sdp.PreStartContainerResponse, error) {
	return nil, nil
}
