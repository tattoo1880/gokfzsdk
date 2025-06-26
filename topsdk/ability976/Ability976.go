package ability976

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/ability976/request"
    "topsdk/ability976/response"
    "topsdk/util"
)

type Ability976 struct {
    Client *topsdk.TopClient
}

func NewAbility976(client *topsdk.TopClient) *Ability976{
    return &Ability976{client}
}

/*
    校验用户是否已经登录
*/
func (ability *Ability976) AlibabaInteractUserIslogin(req *request.AlibabaInteractUserIsloginRequest) (*response.AlibabaInteractUserIsloginResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability976 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.interact.user.islogin",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaInteractUserIsloginResponse{}
    if(err != nil){
        log.Println("alibabaInteractUserIslogin error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
