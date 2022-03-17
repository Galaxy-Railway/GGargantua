package single_request

import "github.com/Galaxy-Railway/GGargantua/internal/gargantua/app/single_request/module"

func SendSingleRequest(req *module.SingleRequest)(*module.SingleResponse, error){
	sender := module.BuildSenderBySingleRequest(req)
	err := sender.Send()
	if err != nil{
		return nil, err
	}
	return module.GetSingleResponseFromSender(sender), nil
}
