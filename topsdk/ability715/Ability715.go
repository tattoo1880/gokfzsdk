package ability715

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/ability715/request"
    "topsdk/ability715/response"
    "topsdk/util"
)

type Ability715 struct {
    Client *topsdk.TopClient
}

func NewAbility715(client *topsdk.TopClient) *Ability715{
    return &Ability715{client}
}

/*
    是否关注
*/
func (ability *Ability715) TaobaoWeitaoFeedIsrelation(req *request.TaobaoWeitaoFeedIsrelationRequest) (*response.TaobaoWeitaoFeedIsrelationResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability715 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.weitao.feed.isrelation",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoWeitaoFeedIsrelationResponse{}
    if(err != nil){
        log.Println("taobaoWeitaoFeedIsrelation error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    微淘是否关注
*/
func (ability *Ability715) TaobaoWeitaoFollowIsrelation(req *request.TaobaoWeitaoFollowIsrelationRequest,session string) (*response.TaobaoWeitaoFollowIsrelationResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability715 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.weitao.follow.isrelation",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWeitaoFollowIsrelationResponse{}
    if(err != nil){
        log.Println("taobaoWeitaoFollowIsrelation error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
