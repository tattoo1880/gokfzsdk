package ability222

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/ability222/request"
    "topsdk/ability222/response"
    "topsdk/util"
)

type Ability222 struct {
    Client *topsdk.TopClient
}

func NewAbility222(client *topsdk.TopClient) *Ability222{
    return &Ability222{client}
}

/*
    ECRM创建淘短链服务
*/
func (ability *Ability222) TaobaoCrmServiceChannelShortlinkCreate(req *request.TaobaoCrmServiceChannelShortlinkCreateRequest,session string) (*response.TaobaoCrmServiceChannelShortlinkCreateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability222 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.crm.service.channel.shortlink.create",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoCrmServiceChannelShortlinkCreateResponse{}
    if(err != nil){
        log.Println("taobaoCrmServiceChannelShortlinkCreate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
