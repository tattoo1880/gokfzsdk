package ability1092

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/ability1092/request"
    "topsdk/ability1092/response"
    "topsdk/util"
)

type Ability1092 struct {
    Client *topsdk.TopClient
}

func NewAbility1092(client *topsdk.TopClient) *Ability1092{
    return &Ability1092{client}
}

/*
    编辑门店覆盖范围
*/
func (ability *Ability1092) TaobaoShopCoverageManage(req *request.TaobaoShopCoverageManageRequest,session string) (*response.TaobaoShopCoverageManageResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability1092 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.shop.coverage.manage",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoShopCoverageManageResponse{}
    if(err != nil){
        log.Println("taobaoShopCoverageManage error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询门店覆盖范围
*/
func (ability *Ability1092) TaobaoShopCoverageQuery(req *request.TaobaoShopCoverageQueryRequest,session string) (*response.TaobaoShopCoverageQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability1092 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.shop.coverage.query",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoShopCoverageQueryResponse{}
    if(err != nil){
        log.Println("taobaoShopCoverageQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    编辑仓库覆盖范围
*/
func (ability *Ability1092) TaobaoRegionWarehouseManage(req *request.TaobaoRegionWarehouseManageRequest,session string) (*response.TaobaoRegionWarehouseManageResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability1092 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.region.warehouse.manage",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoRegionWarehouseManageResponse{}
    if(err != nil){
        log.Println("taobaoRegionWarehouseManage error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询仓库覆盖范围
*/
func (ability *Ability1092) TaobaoRegionWarehouseQuery(req *request.TaobaoRegionWarehouseQueryRequest,session string) (*response.TaobaoRegionWarehouseQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability1092 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.region.warehouse.query",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoRegionWarehouseQueryResponse{}
    if(err != nil){
        log.Println("taobaoRegionWarehouseQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    编辑区域价格
*/
func (ability *Ability1092) TaobaoRegionPriceManage(req *request.TaobaoRegionPriceManageRequest,session string) (*response.TaobaoRegionPriceManageResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability1092 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.region.price.manage",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoRegionPriceManageResponse{}
    if(err != nil){
        log.Println("taobaoRegionPriceManage error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    区域价格查询
*/
func (ability *Ability1092) TaobaoRegionPriceQuery(req *request.TaobaoRegionPriceQueryRequest,session string) (*response.TaobaoRegionPriceQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability1092 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.region.price.query",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoRegionPriceQueryResponse{}
    if(err != nil){
        log.Println("taobaoRegionPriceQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
