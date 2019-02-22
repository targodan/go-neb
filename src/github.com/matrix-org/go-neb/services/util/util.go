package util

import (
	"strings"

	"github.com/matrix-org/go-neb/types"
	"github.com/matrix-org/gomatrix"
)

const ServiceType = "util"

const rollHelp = "Rolls dice for you. You can for example say \"!roll 3d20 2d6\" to roll three 20-sided dice and two 6-sided dice."

type Service struct {
	types.DefaultService
}

func (s *Service) roll(roomID, userID string, args []string) (interface{}, error) {
	dice := NewDice()
	for _, arg := range args {
		dice = dice.Join(ParseDice(arg))
	}

	results := ItoaSlice(dice.Roll())

	return &gomatrix.TextMessage{
		MsgType: "m.notice",
		Body:    "You rolled: " + strings.Join(results, ", "),
	}, nil
}

func (s *Service) Commands(cli *gomatrix.Client) []types.Command {
	return []types.Command{
		types.Command{
			Path:    []string{"roll"},
			Command: s.roll,
			Help:    rollHelp,
		},
	}
}

func init() {
	types.RegisterService(func(serviceID, serviceUserID, webhookEndpointURL string) types.Service {
		return &Service{
			DefaultService: types.NewDefaultService(serviceID, serviceUserID, ServiceType),
		}
	})
}
