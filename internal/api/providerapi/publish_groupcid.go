package providerapi

import (
	"net"

	"github.com/ConsenSys/fc-retrieval-gateway/pkg/logging"
	"github.com/ConsenSys/fc-retrieval-gateway/pkg/messages"
)

func handleProviderPublishGroupCIDRequest(conn net.Conn, request *messages.ProviderDHTPublishGroupCIDRequest) error {
	logging.Info("Provider request from: %s", request.ProviderID.ToString())
	return nil
}
