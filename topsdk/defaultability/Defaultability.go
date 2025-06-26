package defaultability

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/defaultability/request"
    "topsdk/defaultability/response"
    "topsdk/util"
)

type Defaultability struct {
    Client *topsdk.TopClient
}

func NewDefaultability(client *topsdk.TopClient) *Defaultability{
    return &Defaultability{client}
}

/*
    查询文件上传多媒体平台的进度结果
*/
func (ability *Defaultability) TaobaoContentMediaUploadGet(req *request.TaobaoContentMediaUploadGetRequest,session string) (*response.TaobaoContentMediaUploadGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.content.media.upload.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoContentMediaUploadGetResponse{}
    if(err != nil){
        log.Println("taobaoContentMediaUploadGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取oss上传的许可secret
*/
func (ability *Defaultability) TaobaoContentMediaUploadSecret(req *request.TaobaoContentMediaUploadSecretRequest) (*response.TaobaoContentMediaUploadSecretResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.content.media.upload.secret",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoContentMediaUploadSecretResponse{}
    if(err != nil){
        log.Println("taobaoContentMediaUploadSecret error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    同步文件到淘宝多媒体平台
*/
func (ability *Defaultability) TaobaoContentMediaUploadPub(req *request.TaobaoContentMediaUploadPubRequest,session string) (*response.TaobaoContentMediaUploadPubResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.content.media.upload.pub",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoContentMediaUploadPubResponse{}
    if(err != nil){
        log.Println("taobaoContentMediaUploadPub error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取当前会话用户出售中的商品列表
*/
func (ability *Defaultability) TaobaoItemsOnsaleGet(req *request.TaobaoItemsOnsaleGetRequest,session string) (*response.TaobaoItemsOnsaleGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.items.onsale.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemsOnsaleGetResponse{}
    if(err != nil){
        log.Println("taobaoItemsOnsaleGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    更新商品信息
*/
func (ability *Defaultability) TaobaoItemUpdate(req *request.TaobaoItemUpdateRequest,session string) (*response.TaobaoItemUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemUpdateResponse{}
    if(err != nil){
        log.Println("taobaoItemUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    添加一个商品
*/
func (ability *Defaultability) TaobaoItemAdd(req *request.TaobaoItemAddRequest,session string) (*response.TaobaoItemAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemAddResponse{}
    if(err != nil){
        log.Println("taobaoItemAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    添加商品图片
*/
func (ability *Defaultability) TaobaoItemImgUpload(req *request.TaobaoItemImgUploadRequest,session string) (*response.TaobaoItemImgUploadResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.img.upload",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemImgUploadResponse{}
    if(err != nil){
        log.Println("taobaoItemImgUpload error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    删除商品图片
*/
func (ability *Defaultability) TaobaoItemImgDelete(req *request.TaobaoItemImgDeleteRequest,session string) (*response.TaobaoItemImgDeleteResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.img.delete",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemImgDeleteResponse{}
    if(err != nil){
        log.Println("taobaoItemImgDelete error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    删除属性图片
*/
func (ability *Defaultability) TaobaoItemPropimgDelete(req *request.TaobaoItemPropimgDeleteRequest,session string) (*response.TaobaoItemPropimgDeleteResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.propimg.delete",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemPropimgDeleteResponse{}
    if(err != nil){
        log.Println("taobaoItemPropimgDelete error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取商家所在分组及其已授权(广播)消息topics
*/
func (ability *Defaultability) TaobaoTmcUserGet(req *request.TaobaoTmcUserGetRequest) (*response.TaobaoTmcUserGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tmc.user.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTmcUserGetResponse{}
    if(err != nil){
        log.Println("taobaoTmcUserGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    添加或修改属性图片
*/
func (ability *Defaultability) TaobaoItemPropimgUpload(req *request.TaobaoItemPropimgUploadRequest,session string) (*response.TaobaoItemPropimgUploadResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.propimg.upload",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemPropimgUploadResponse{}
    if(err != nil){
        log.Println("taobaoItemPropimgUpload error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取SKU
*/
func (ability *Defaultability) TaobaoItemSkuGet(req *request.TaobaoItemSkuGetRequest,session string) (*response.TaobaoItemSkuGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.sku.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemSkuGetResponse{}
    if(err != nil){
        log.Println("taobaoItemSkuGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    根据商品ID列表获取SKU信息
*/
func (ability *Defaultability) TaobaoItemSkusGet(req *request.TaobaoItemSkusGetRequest,session string) (*response.TaobaoItemSkusGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.skus.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemSkusGetResponse{}
    if(err != nil){
        log.Println("taobaoItemSkusGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取单个商品详细信息
*/
func (ability *Defaultability) TaobaoItemSellerGet(req *request.TaobaoItemSellerGetRequest,session string) (*response.TaobaoItemSellerGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.seller.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemSellerGetResponse{}
    if(err != nil){
        log.Println("taobaoItemSellerGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    批量获取商品详细信息
*/
func (ability *Defaultability) TaobaoItemsSellerListGet(req *request.TaobaoItemsSellerListGetRequest,session string) (*response.TaobaoItemsSellerListGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.items.seller.list.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemsSellerListGetResponse{}
    if(err != nil){
        log.Println("taobaoItemsSellerListGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询买家申请的退款列表
*/
func (ability *Defaultability) TaobaoRefundsApplyGet(req *request.TaobaoRefundsApplyGetRequest,session string) (*response.TaobaoRefundsApplyGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.refunds.apply.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoRefundsApplyGetResponse{}
    if(err != nil){
        log.Println("taobaoRefundsApplyGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取单笔退款详情
*/
func (ability *Defaultability) TaobaoRefundGet(req *request.TaobaoRefundGetRequest,session string) (*response.TaobaoRefundGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.refund.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoRefundGetResponse{}
    if(err != nil){
        log.Println("taobaoRefundGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取前台展示的店铺类目
*/
func (ability *Defaultability) TaobaoShopcatsListGet(req *request.TaobaoShopcatsListGetRequest) (*response.TaobaoShopcatsListGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.shopcats.list.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoShopcatsListGetResponse{}
    if(err != nil){
        log.Println("taobaoShopcatsListGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取前台展示的店铺内卖家自定义商品类目
*/
func (ability *Defaultability) TaobaoSellercatsListGet(req *request.TaobaoSellercatsListGetRequest,session string) (*response.TaobaoSellercatsListGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.sellercats.list.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSellercatsListGetResponse{}
    if(err != nil){
        log.Println("taobaoSellercatsListGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    创建码平台常用二维码
*/
func (ability *Defaultability) TaobaoMaQrcodeCommonCreate(req *request.TaobaoMaQrcodeCommonCreateRequest,session string) (*response.TaobaoMaQrcodeCommonCreateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.ma.qrcode.common.create",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoMaQrcodeCommonCreateResponse{}
    if(err != nil){
        log.Println("taobaoMaQrcodeCommonCreate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取后台供卖家发布商品的标准商品类目
*/
func (ability *Defaultability) TaobaoItemcatsGet(req *request.TaobaoItemcatsGetRequest,session string) (*response.TaobaoItemcatsGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.itemcats.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemcatsGetResponse{}
    if(err != nil){
        log.Println("taobaoItemcatsGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询退款留言/凭证列表
*/
func (ability *Defaultability) TaobaoRefundMessagesGet(req *request.TaobaoRefundMessagesGetRequest,session string) (*response.TaobaoRefundMessagesGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.refund.messages.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoRefundMessagesGetResponse{}
    if(err != nil){
        log.Println("taobaoRefundMessagesGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    创建退款留言/凭证
*/
func (ability *Defaultability) TaobaoRefundMessageAdd(req *request.TaobaoRefundMessageAddRequest,session string) (*response.TaobaoRefundMessageAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.refund.message.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoRefundMessageAddResponse{}
    if(err != nil){
        log.Println("taobaoRefundMessageAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取所有的菜鸟标准电子面单模板
*/
func (ability *Defaultability) CainiaoCloudprintStdtemplatesGet(req *request.CainiaoCloudprintStdtemplatesGetRequest) (*response.CainiaoCloudprintStdtemplatesGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("cainiao.cloudprint.stdtemplates.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.CainiaoCloudprintStdtemplatesGetResponse{}
    if(err != nil){
        log.Println("cainiaoCloudprintStdtemplatesGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取用户使用的菜鸟电子面单模板信息
*/
func (ability *Defaultability) CainiaoCloudprintMystdtemplatesGet(req *request.CainiaoCloudprintMystdtemplatesGetRequest,session string) (*response.CainiaoCloudprintMystdtemplatesGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("cainiao.cloudprint.mystdtemplates.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.CainiaoCloudprintMystdtemplatesGetResponse{}
    if(err != nil){
        log.Println("cainiaoCloudprintMystdtemplatesGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取图片分类信息
*/
func (ability *Defaultability) TaobaoPictureCategoryGet(req *request.TaobaoPictureCategoryGetRequest,session string) (*response.TaobaoPictureCategoryGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.picture.category.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoPictureCategoryGetResponse{}
    if(err != nil){
        log.Println("taobaoPictureCategoryGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    服务平台营销链接生成接口
*/
func (ability *Defaultability) TaobaoFuwuSaleLinkGen(req *request.TaobaoFuwuSaleLinkGenRequest) (*response.TaobaoFuwuSaleLinkGenResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.fuwu.sale.link.gen",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoFuwuSaleLinkGenResponse{}
    if(err != nil){
        log.Println("taobaoFuwuSaleLinkGen error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取图片信息
*/
func (ability *Defaultability) TaobaoPictureGet(req *request.TaobaoPictureGetRequest,session string) (*response.TaobaoPictureGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.picture.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoPictureGetResponse{}
    if(err != nil){
        log.Println("taobaoPictureGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    上传单张图片
*/
func (ability *Defaultability) TaobaoPictureUpload(req *request.TaobaoPictureUploadRequest,session string) (*response.TaobaoPictureUploadResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.picture.upload",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoPictureUploadResponse{}
    if(err != nil){
        log.Println("taobaoPictureUpload error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    商家查询物流商产品类型接口
*/
func (ability *Defaultability) CainiaoWaybillIiProduct(req *request.CainiaoWaybillIiProductRequest,session string) (*response.CainiaoWaybillIiProductResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("cainiao.waybill.ii.product",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.CainiaoWaybillIiProductResponse{}
    if(err != nil){
        log.Println("cainiaoWaybillIiProduct error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    商家取消获取的电子面单号
*/
func (ability *Defaultability) CainiaoWaybillIiCancel(req *request.CainiaoWaybillIiCancelRequest,session string) (*response.CainiaoWaybillIiCancelResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("cainiao.waybill.ii.cancel",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.CainiaoWaybillIiCancelResponse{}
    if(err != nil){
        log.Println("cainiaoWaybillIiCancel error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    关键词过滤匹配
*/
func (ability *Defaultability) TaobaoKfcKeywordSearch(req *request.TaobaoKfcKeywordSearchRequest,session string) (*response.TaobaoKfcKeywordSearchResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.kfc.keyword.search",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoKfcKeywordSearchResponse{}
    if(err != nil){
        log.Println("taobaoKfcKeywordSearch error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    商品关联子图
*/
func (ability *Defaultability) TaobaoItemJointImg(req *request.TaobaoItemJointImgRequest,session string) (*response.TaobaoItemJointImgResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.joint.img",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemJointImgResponse{}
    if(err != nil){
        log.Println("taobaoItemJointImg error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    商品关联属性图
*/
func (ability *Defaultability) TaobaoItemJointPropimg(req *request.TaobaoItemJointPropimgRequest,session string) (*response.TaobaoItemJointPropimgResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.joint.propimg",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemJointPropimgResponse{}
    if(err != nil){
        log.Println("taobaoItemJointPropimg error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    异常订单信息获取
*/
func (ability *Defaultability) TaobaoOcTradetraceAlertsGet(req *request.TaobaoOcTradetraceAlertsGetRequest,session string) (*response.TaobaoOcTradetraceAlertsGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.oc.tradetrace.alerts.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoOcTradetraceAlertsGetResponse{}
    if(err != nil){
        log.Println("taobaoOcTradetraceAlertsGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询商家被授权品牌列表和类目列表
*/
func (ability *Defaultability) TaobaoItemcatsAuthorizeGet(req *request.TaobaoItemcatsAuthorizeGetRequest,session string) (*response.TaobaoItemcatsAuthorizeGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.itemcats.authorize.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemcatsAuthorizeGetResponse{}
    if(err != nil){
        log.Println("taobaoItemcatsAuthorizeGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    得到当前会话用户库存中的商品列表
*/
func (ability *Defaultability) TaobaoItemsInventoryGet(req *request.TaobaoItemsInventoryGetRequest,session string) (*response.TaobaoItemsInventoryGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.items.inventory.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemsInventoryGetResponse{}
    if(err != nil){
        log.Println("taobaoItemsInventoryGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    根据外部ID取商品SKU
*/
func (ability *Defaultability) TaobaoSkusCustomGet(req *request.TaobaoSkusCustomGetRequest,session string) (*response.TaobaoSkusCustomGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.skus.custom.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSkusCustomGetResponse{}
    if(err != nil){
        log.Println("taobaoSkusCustomGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取商家的自定义区模板信息
*/
func (ability *Defaultability) CainiaoCloudprintCustomaresGet(req *request.CainiaoCloudprintCustomaresGetRequest,session string) (*response.CainiaoCloudprintCustomaresGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("cainiao.cloudprint.customares.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.CainiaoCloudprintCustomaresGetResponse{}
    if(err != nil){
        log.Println("cainiaoCloudprintCustomaresGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取指定账户的子账号简易信息列表
*/
func (ability *Defaultability) TaobaoSubusersGet(req *request.TaobaoSubusersGetRequest,session string) (*response.TaobaoSubusersGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.subusers.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSubusersGetResponse{}
    if(err != nil){
        log.Println("taobaoSubusersGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取指定账户子账号的详细信息
*/
func (ability *Defaultability) TaobaoSubuserFullinfoGet(req *request.TaobaoSubuserFullinfoGetRequest,session string) (*response.TaobaoSubuserFullinfoGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.subuser.fullinfo.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSubuserFullinfoGetResponse{}
    if(err != nil){
        log.Println("taobaoSubuserFullinfoGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取指定账户的所有部门列表
*/
func (ability *Defaultability) TaobaoSubuserDepartmentsGet(req *request.TaobaoSubuserDepartmentsGetRequest,session string) (*response.TaobaoSubuserDepartmentsGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.subuser.departments.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSubuserDepartmentsGetResponse{}
    if(err != nil){
        log.Println("taobaoSubuserDepartmentsGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取指定账户的所有职务信息列表
*/
func (ability *Defaultability) TaobaoSubuserDutysGet(req *request.TaobaoSubuserDutysGetRequest,session string) (*response.TaobaoSubuserDutysGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.subuser.dutys.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSubuserDutysGetResponse{}
    if(err != nil){
        log.Println("taobaoSubuserDutysGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    修改指定账户子账号的基本信息
*/
func (ability *Defaultability) TaobaoSubuserInfoUpdate(req *request.TaobaoSubuserInfoUpdateRequest,session string) (*response.TaobaoSubuserInfoUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.subuser.info.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSubuserInfoUpdateResponse{}
    if(err != nil){
        log.Println("taobaoSubuserInfoUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    新增图片分类信息
*/
func (ability *Defaultability) TaobaoPictureCategoryAdd(req *request.TaobaoPictureCategoryAddRequest,session string) (*response.TaobaoPictureCategoryAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.picture.category.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoPictureCategoryAddResponse{}
    if(err != nil){
        log.Println("taobaoPictureCategoryAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    电子面单云打印接口
*/
func (ability *Defaultability) CainiaoWaybillIiGet(req *request.CainiaoWaybillIiGetRequest,session string) (*response.CainiaoWaybillIiGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("cainiao.waybill.ii.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.CainiaoWaybillIiGetResponse{}
    if(err != nil){
        log.Println("cainiaoWaybillIiGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    电子面单云打印更新接口
*/
func (ability *Defaultability) CainiaoWaybillIiUpdate(req *request.CainiaoWaybillIiUpdateRequest,session string) (*response.CainiaoWaybillIiUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("cainiao.waybill.ii.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.CainiaoWaybillIiUpdateResponse{}
    if(err != nil){
        log.Println("cainiaoWaybillIiUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    商品编辑增量更新
*/
func (ability *Defaultability) AlibabaItemEditFastupdate(req *request.AlibabaItemEditFastupdateRequest,session string) (*response.AlibabaItemEditFastupdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.item.edit.fastupdate",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaItemEditFastupdateResponse{}
    if(err != nil){
        log.Println("alibabaItemEditFastupdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询appstore应用订购关系
*/
func (ability *Defaultability) TaobaoAppstoreSubscribeGet(req *request.TaobaoAppstoreSubscribeGetRequest) (*response.TaobaoAppstoreSubscribeGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.appstore.subscribe.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoAppstoreSubscribeGetResponse{}
    if(err != nil){
        log.Println("taobaoAppstoreSubscribeGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取单条订单跟踪详情
*/
func (ability *Defaultability) TaobaoJdsTradeTracesGet(req *request.TaobaoJdsTradeTracesGetRequest,session string) (*response.TaobaoJdsTradeTracesGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.jds.trade.traces.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoJdsTradeTracesGetResponse{}
    if(err != nil){
        log.Println("taobaoJdsTradeTracesGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    删除SKU
*/
func (ability *Defaultability) TaobaoItemSkuDelete(req *request.TaobaoItemSkuDeleteRequest,session string) (*response.TaobaoItemSkuDeleteResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.sku.delete",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemSkuDeleteResponse{}
    if(err != nil){
        log.Println("taobaoItemSkuDelete error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取物流服务商电子面单号v1.0
*/
func (ability *Defaultability) TaobaoWlbWaybillIGet(req *request.TaobaoWlbWaybillIGetRequest,session string) (*response.TaobaoWlbWaybillIGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.waybill.i.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbWaybillIGetResponse{}
    if(err != nil){
        log.Println("taobaoWlbWaybillIGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    面单信息更新接口v1.0
*/
func (ability *Defaultability) TaobaoWlbWaybillIFullupdate(req *request.TaobaoWlbWaybillIFullupdateRequest,session string) (*response.TaobaoWlbWaybillIFullupdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.waybill.i.fullupdate",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbWaybillIFullupdateResponse{}
    if(err != nil){
        log.Println("taobaoWlbWaybillIFullupdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    打印确认接口v1.0
*/
func (ability *Defaultability) TaobaoWlbWaybillIPrint(req *request.TaobaoWlbWaybillIPrintRequest,session string) (*response.TaobaoWlbWaybillIPrintResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.waybill.i.print",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbWaybillIPrintResponse{}
    if(err != nil){
        log.Println("taobaoWlbWaybillIPrint error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查面单号状态v1.0
*/
func (ability *Defaultability) TaobaoWlbWaybillIQuerydetail(req *request.TaobaoWlbWaybillIQuerydetailRequest,session string) (*response.TaobaoWlbWaybillIQuerydetailResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.waybill.i.querydetail",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbWaybillIQuerydetailResponse{}
    if(err != nil){
        log.Println("taobaoWlbWaybillIQuerydetail error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    商家取消获取的电子面单号v1.0
*/
func (ability *Defaultability) TaobaoWlbWaybillICancel(req *request.TaobaoWlbWaybillICancelRequest,session string) (*response.TaobaoWlbWaybillICancelResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.waybill.i.cancel",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbWaybillICancelResponse{}
    if(err != nil){
        log.Println("taobaoWlbWaybillICancel error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    添加产品
*/
func (ability *Defaultability) TaobaoFenxiaoProductAdd(req *request.TaobaoFenxiaoProductAddRequest,session string) (*response.TaobaoFenxiaoProductAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.product.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoProductAddResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoProductAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    更新产品
*/
func (ability *Defaultability) TaobaoFenxiaoProductUpdate(req *request.TaobaoFenxiaoProductUpdateRequest,session string) (*response.TaobaoFenxiaoProductUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.product.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoProductUpdateResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoProductUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询产品线列表
*/
func (ability *Defaultability) TaobaoFenxiaoProductcatsGet(req *request.TaobaoFenxiaoProductcatsGetRequest,session string) (*response.TaobaoFenxiaoProductcatsGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.productcats.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoProductcatsGetResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoProductcatsGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询产品列表
*/
func (ability *Defaultability) TaobaoFenxiaoProductsGet(req *request.TaobaoFenxiaoProductsGetRequest,session string) (*response.TaobaoFenxiaoProductsGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.products.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoProductsGetResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoProductsGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    更新合作关系等级
*/
func (ability *Defaultability) TaobaoFenxiaoCooperationUpdate(req *request.TaobaoFenxiaoCooperationUpdateRequest,session string) (*response.TaobaoFenxiaoCooperationUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.cooperation.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoCooperationUpdateResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoCooperationUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    分销商等级查询
*/
func (ability *Defaultability) TaobaoFenxiaoGradesGet(req *request.TaobaoFenxiaoGradesGetRequest,session string) (*response.TaobaoFenxiaoGradesGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.grades.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoGradesGetResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoGradesGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取折扣信息
*/
func (ability *Defaultability) TaobaoFenxiaoDiscountsGet(req *request.TaobaoFenxiaoDiscountsGetRequest,session string) (*response.TaobaoFenxiaoDiscountsGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.discounts.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoDiscountsGetResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoDiscountsGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    商家查询物流商产品类型接口
*/
func (ability *Defaultability) TaobaoWlbWaybillIProduct(req *request.TaobaoWlbWaybillIProductRequest,session string) (*response.TaobaoWlbWaybillIProductResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.waybill.i.product",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbWaybillIProductResponse{}
    if(err != nil){
        log.Println("taobaoWlbWaybillIProduct error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    订购关系查询
*/
func (ability *Defaultability) TaobaoVasSubscribeGet(req *request.TaobaoVasSubscribeGetRequest) (*response.TaobaoVasSubscribeGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.vas.subscribe.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoVasSubscribeGetResponse{}
    if(err != nil){
        log.Println("taobaoVasSubscribeGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    订单记录导出
*/
func (ability *Defaultability) TaobaoVasOrderSearch(req *request.TaobaoVasOrderSearchRequest) (*response.TaobaoVasOrderSearchResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.vas.order.search",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoVasOrderSearchResponse{}
    if(err != nil){
        log.Println("taobaoVasOrderSearch error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    订购记录导出
*/
func (ability *Defaultability) TaobaoVasSubscSearch(req *request.TaobaoVasSubscSearchRequest) (*response.TaobaoVasSubscSearchResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.vas.subsc.search",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoVasSubscSearchResponse{}
    if(err != nil){
        log.Println("taobaoVasSubscSearch error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    宝贝/SKU库存修改
*/
func (ability *Defaultability) TaobaoItemQuantityUpdate(req *request.TaobaoItemQuantityUpdateRequest,session string) (*response.TaobaoItemQuantityUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.quantity.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemQuantityUpdateResponse{}
    if(err != nil){
        log.Println("taobaoItemQuantityUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    更新商品条形码信息
*/
func (ability *Defaultability) TaobaoItemBarcodeUpdate(req *request.TaobaoItemBarcodeUpdateRequest,session string) (*response.TaobaoItemBarcodeUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.barcode.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemBarcodeUpdateResponse{}
    if(err != nil){
        log.Println("taobaoItemBarcodeUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    确认收款
*/
func (ability *Defaultability) TaobaoFenxiaoOrderConfirmPaid(req *request.TaobaoFenxiaoOrderConfirmPaidRequest,session string) (*response.TaobaoFenxiaoOrderConfirmPaidResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.order.confirm.paid",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoOrderConfirmPaidResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoOrderConfirmPaid error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询商品下载记录
*/
func (ability *Defaultability) TaobaoFenxiaoDistributorItemsGet(req *request.TaobaoFenxiaoDistributorItemsGetRequest,session string) (*response.TaobaoFenxiaoDistributorItemsGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.distributor.items.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoDistributorItemsGetResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoDistributorItemsGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    供应商或分销商获取合作关系信息
*/
func (ability *Defaultability) TaobaoFenxiaoCooperationGet(req *request.TaobaoFenxiaoCooperationGetRequest,session string) (*response.TaobaoFenxiaoCooperationGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.cooperation.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoCooperationGetResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoCooperationGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取用户宝贝详情页模板名称
*/
func (ability *Defaultability) TaobaoItemTemplatesGet(req *request.TaobaoItemTemplatesGetRequest,session string) (*response.TaobaoItemTemplatesGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.templates.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemTemplatesGetResponse{}
    if(err != nil){
        log.Println("taobaoItemTemplatesGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    发布单条消息
*/
func (ability *Defaultability) TaobaoTmcMessageProduce(req *request.TaobaoTmcMessageProduceRequest) (*response.TaobaoTmcMessageProduceResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tmc.message.produce",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTmcMessageProduceResponse{}
    if(err != nil){
        log.Println("taobaoTmcMessageProduce error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    取消用户的消息服务
*/
func (ability *Defaultability) TaobaoTmcUserCancel(req *request.TaobaoTmcUserCancelRequest) (*response.TaobaoTmcUserCancelResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tmc.user.cancel",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTmcUserCancelResponse{}
    if(err != nil){
        log.Println("taobaoTmcUserCancel error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    为已授权的用户开通消息服务
*/
func (ability *Defaultability) TaobaoTmcUserPermit(req *request.TaobaoTmcUserPermitRequest,session string) (*response.TaobaoTmcUserPermitResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.tmc.user.permit",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoTmcUserPermitResponse{}
    if(err != nil){
        log.Println("taobaoTmcUserPermit error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    通过订单号查询电子面单通接口
*/
func (ability *Defaultability) CainiaoWaybillIiQueryByTradecode(req *request.CainiaoWaybillIiQueryByTradecodeRequest,session string) (*response.CainiaoWaybillIiQueryByTradecodeResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("cainiao.waybill.ii.query.by.tradecode",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.CainiaoWaybillIiQueryByTradecodeResponse{}
    if(err != nil){
        log.Println("cainiaoWaybillIiQueryByTradecode error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    通过面单号查询面单打印报文
*/
func (ability *Defaultability) CainiaoWaybillIiQueryByWaybillcode(req *request.CainiaoWaybillIiQueryByWaybillcodeRequest,session string) (*response.CainiaoWaybillIiQueryByWaybillcodeResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("cainiao.waybill.ii.query.by.waybillcode",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.CainiaoWaybillIiQueryByWaybillcodeResponse{}
    if(err != nil){
        log.Println("cainiaoWaybillIiQueryByWaybillcode error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取商家使用的标准模板
*/
func (ability *Defaultability) CainiaoCloudprintIsvtemplatesGet(req *request.CainiaoCloudprintIsvtemplatesGetRequest,session string) (*response.CainiaoCloudprintIsvtemplatesGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("cainiao.cloudprint.isvtemplates.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.CainiaoCloudprintIsvtemplatesGetResponse{}
    if(err != nil){
        log.Println("cainiaoCloudprintIsvtemplatesGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    天猫商品/SKU商家编码更新接口
*/
func (ability *Defaultability) TmallItemOuteridUpdate(req *request.TmallItemOuteridUpdateRequest,session string) (*response.TmallItemOuteridUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.item.outerid.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallItemOuteridUpdateResponse{}
    if(err != nil){
        log.Println("tmallItemOuteridUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询面单服务订购及面单使用情况
*/
func (ability *Defaultability) CainiaoWaybillIiSearch(req *request.CainiaoWaybillIiSearchRequest,session string) (*response.CainiaoWaybillIiSearchResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("cainiao.waybill.ii.search",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.CainiaoWaybillIiSearchResponse{}
    if(err != nil){
        log.Println("cainiaoWaybillIiSearch error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询指定账户的子账号列表
*/
func (ability *Defaultability) TaobaoSellercenterSubusersGet(req *request.TaobaoSellercenterSubusersGetRequest,session string) (*response.TaobaoSellercenterSubusersGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.sellercenter.subusers.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSellercenterSubusersGetResponse{}
    if(err != nil){
        log.Println("taobaoSellercenterSubusersGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取指定用户的权限集合
*/
func (ability *Defaultability) TaobaoSellercenterUserPermissionsGet(req *request.TaobaoSellercenterUserPermissionsGetRequest,session string) (*response.TaobaoSellercenterUserPermissionsGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.sellercenter.user.permissions.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSellercenterUserPermissionsGetResponse{}
    if(err != nil){
        log.Println("taobaoSellercenterUserPermissionsGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取商品已生效营销活动更新规则
*/
func (ability *Defaultability) TaobaoItemPromotionRuleGet(req *request.TaobaoItemPromotionRuleGetRequest,session string) (*response.TaobaoItemPromotionRuleGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.promotion.rule.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemPromotionRuleGetResponse{}
    if(err != nil){
        log.Println("taobaoItemPromotionRuleGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取分销用户登录信息
*/
func (ability *Defaultability) TaobaoFenxiaoLoginUserGet(req *request.TaobaoFenxiaoLoginUserGetRequest,session string) (*response.TaobaoFenxiaoLoginUserGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.login.user.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoLoginUserGetResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoLoginUserGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    分销铺货商品发布
*/
func (ability *Defaultability) AlibabaItemPublishDistributeSubmit(req *request.AlibabaItemPublishDistributeSubmitRequest,session string) (*response.AlibabaItemPublishDistributeSubmitResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.item.publish.distribute.submit",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaItemPublishDistributeSubmitResponse{}
    if(err != nil){
        log.Println("alibabaItemPublishDistributeSubmit error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询指定的子账号的权限和角色信息
*/
func (ability *Defaultability) TaobaoSellercenterSubuserPermissionsRolesGet(req *request.TaobaoSellercenterSubuserPermissionsRolesGetRequest,session string) (*response.TaobaoSellercenterSubuserPermissionsRolesGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.sellercenter.subuser.permissions.roles.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSellercenterSubuserPermissionsRolesGetResponse{}
    if(err != nil){
        log.Println("taobaoSellercenterSubuserPermissionsRolesGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取商品发布规则信息 for 分销铺货
*/
func (ability *Defaultability) AlibabaItemPublishSchemaDistributeGet(req *request.AlibabaItemPublishSchemaDistributeGetRequest,session string) (*response.AlibabaItemPublishSchemaDistributeGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.item.publish.schema.distribute.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaItemPublishSchemaDistributeGetResponse{}
    if(err != nil){
        log.Println("alibabaItemPublishSchemaDistributeGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    商品级联属性信息获取-分销铺货
*/
func (ability *Defaultability) AlibabaItemPublishPropsDistributeGet(req *request.AlibabaItemPublishPropsDistributeGetRequest,session string) (*response.AlibabaItemPublishPropsDistributeGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.item.publish.props.distribute.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaItemPublishPropsDistributeGetResponse{}
    if(err != nil){
        log.Println("alibabaItemPublishPropsDistributeGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取指定卖家的角色列表
*/
func (ability *Defaultability) TaobaoSellercenterRolesGet(req *request.TaobaoSellercenterRolesGetRequest,session string) (*response.TaobaoSellercenterRolesGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.sellercenter.roles.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSellercenterRolesGetResponse{}
    if(err != nil){
        log.Println("taobaoSellercenterRolesGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    子账号角色的新增（指定卖家）
*/
func (ability *Defaultability) TaobaoSellercenterRoleAdd(req *request.TaobaoSellercenterRoleAddRequest,session string) (*response.TaobaoSellercenterRoleAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.sellercenter.role.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSellercenterRoleAddResponse{}
    if(err != nil){
        log.Println("taobaoSellercenterRoleAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取oss上传的许可secret.hz
*/
func (ability *Defaultability) TaobaoContentMediaUploadSecretHz(req *request.TaobaoContentMediaUploadSecretHzRequest) (*response.TaobaoContentMediaUploadSecretHzResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.content.media.upload.secret.hz",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoContentMediaUploadSecretHzResponse{}
    if(err != nil){
        log.Println("taobaoContentMediaUploadSecretHz error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    服务平台评价查询接口
*/
func (ability *Defaultability) TaobaoFuwuScoresGet(req *request.TaobaoFuwuScoresGetRequest) (*response.TaobaoFuwuScoresGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.fuwu.scores.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoFuwuScoresGetResponse{}
    if(err != nil){
        log.Println("taobaoFuwuScoresGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    更新图片分类
*/
func (ability *Defaultability) TaobaoPictureCategoryUpdate(req *request.TaobaoPictureCategoryUpdateRequest,session string) (*response.TaobaoPictureCategoryUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.picture.category.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoPictureCategoryUpdateResponse{}
    if(err != nil){
        log.Println("taobaoPictureCategoryUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    替换图片
*/
func (ability *Defaultability) TaobaoPictureReplace(req *request.TaobaoPictureReplaceRequest,session string) (*response.TaobaoPictureReplaceResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.picture.replace",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoPictureReplaceResponse{}
    if(err != nil){
        log.Println("taobaoPictureReplace error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    修改图片名字
*/
func (ability *Defaultability) TaobaoPictureUpdate(req *request.TaobaoPictureUpdateRequest,session string) (*response.TaobaoPictureUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.picture.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoPictureUpdateResponse{}
    if(err != nil){
        log.Println("taobaoPictureUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询图片空间用户的信息
*/
func (ability *Defaultability) TaobaoPictureUserinfoGet(req *request.TaobaoPictureUserinfoGetRequest,session string) (*response.TaobaoPictureUserinfoGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.picture.userinfo.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoPictureUserinfoGetResponse{}
    if(err != nil){
        log.Println("taobaoPictureUserinfoGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    图片是否被引用
*/
func (ability *Defaultability) TaobaoPictureIsreferencedGet(req *request.TaobaoPictureIsreferencedGetRequest,session string) (*response.TaobaoPictureIsreferencedGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.picture.isreferenced.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoPictureIsreferencedGetResponse{}
    if(err != nil){
        log.Println("taobaoPictureIsreferencedGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    SKU库存修改
*/
func (ability *Defaultability) TaobaoSkusQuantityUpdate(req *request.TaobaoSkusQuantityUpdateRequest,session string) (*response.TaobaoSkusQuantityUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.skus.quantity.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSkusQuantityUpdateResponse{}
    if(err != nil){
        log.Println("taobaoSkusQuantityUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    产品SKU删除接口
*/
func (ability *Defaultability) TaobaoFenxiaoProductSkuDelete(req *request.TaobaoFenxiaoProductSkuDeleteRequest,session string) (*response.TaobaoFenxiaoProductSkuDeleteResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.product.sku.delete",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoProductSkuDeleteResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoProductSkuDelete error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    产品sku添加接口
*/
func (ability *Defaultability) TaobaoFenxiaoProductSkuAdd(req *request.TaobaoFenxiaoProductSkuAddRequest,session string) (*response.TaobaoFenxiaoProductSkuAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.product.sku.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoProductSkuAddResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoProductSkuAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    产品sku编辑接口
*/
func (ability *Defaultability) TaobaoFenxiaoProductSkuUpdate(req *request.TaobaoFenxiaoProductSkuUpdateRequest,session string) (*response.TaobaoFenxiaoProductSkuUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.product.sku.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoProductSkuUpdateResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoProductSkuUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    SKU查询接口
*/
func (ability *Defaultability) TaobaoFenxiaoProductSkusGet(req *request.TaobaoFenxiaoProductSkusGetRequest,session string) (*response.TaobaoFenxiaoProductSkusGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.product.skus.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoProductSkusGetResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoProductSkusGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    产品图片上传
*/
func (ability *Defaultability) TaobaoFenxiaoProductImageUpload(req *request.TaobaoFenxiaoProductImageUploadRequest,session string) (*response.TaobaoFenxiaoProductImageUploadResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.product.image.upload",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoProductImageUploadResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoProductImageUpload error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    产品图片删除
*/
func (ability *Defaultability) TaobaoFenxiaoProductImageDelete(req *request.TaobaoFenxiaoProductImageDeleteRequest,session string) (*response.TaobaoFenxiaoProductImageDeleteResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.product.image.delete",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoProductImageDeleteResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoProductImageDelete error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    商品编辑获取schema信息
*/
func (ability *Defaultability) AlibabaItemEditSchemaGet(req *request.AlibabaItemEditSchemaGetRequest,session string) (*response.AlibabaItemEditSchemaGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.item.edit.schema.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaItemEditSchemaGetResponse{}
    if(err != nil){
        log.Println("alibabaItemEditSchemaGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    商品编辑提交schema信息
*/
func (ability *Defaultability) AlibabaItemEditSubmit(req *request.AlibabaItemEditSubmitRequest,session string) (*response.AlibabaItemEditSubmitResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.item.edit.submit",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaItemEditSubmitResponse{}
    if(err != nil){
        log.Println("alibabaItemEditSubmit error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    商品删除
*/
func (ability *Defaultability) AlibabaItemOperateDelete(req *request.AlibabaItemOperateDeleteRequest,session string) (*response.AlibabaItemOperateDeleteResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.item.operate.delete",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaItemOperateDeleteResponse{}
    if(err != nil){
        log.Println("alibabaItemOperateDelete error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    商品下架
*/
func (ability *Defaultability) AlibabaItemOperateDownshelf(req *request.AlibabaItemOperateDownshelfRequest,session string) (*response.AlibabaItemOperateDownshelfResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.item.operate.downshelf",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaItemOperateDownshelfResponse{}
    if(err != nil){
        log.Println("alibabaItemOperateDownshelf error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    商品上架
*/
func (ability *Defaultability) AlibabaItemOperateUpshelf(req *request.AlibabaItemOperateUpshelfRequest,session string) (*response.AlibabaItemOperateUpshelfResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.item.operate.upshelf",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaItemOperateUpshelfResponse{}
    if(err != nil){
        log.Println("alibabaItemOperateUpshelf error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    产品发布规则获取接口
*/
func (ability *Defaultability) TmallProductAddSchemaGet(req *request.TmallProductAddSchemaGetRequest,session string) (*response.TmallProductAddSchemaGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.product.add.schema.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallProductAddSchemaGetResponse{}
    if(err != nil){
        log.Println("tmallProductAddSchemaGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取匹配产品规则
*/
func (ability *Defaultability) TmallProductMatchSchemaGet(req *request.TmallProductMatchSchemaGetRequest,session string) (*response.TmallProductMatchSchemaGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.product.match.schema.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallProductMatchSchemaGetResponse{}
    if(err != nil){
        log.Println("tmallProductMatchSchemaGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    product匹配接口
*/
func (ability *Defaultability) TmallProductSchemaMatch(req *request.TmallProductSchemaMatchRequest,session string) (*response.TmallProductSchemaMatchResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.product.schema.match",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallProductSchemaMatchResponse{}
    if(err != nil){
        log.Println("tmallProductSchemaMatch error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    使用Schema文件发布一个产品
*/
func (ability *Defaultability) TmallProductSchemaAdd(req *request.TmallProductSchemaAddRequest,session string) (*response.TmallProductSchemaAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.product.schema.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallProductSchemaAddResponse{}
    if(err != nil){
        log.Println("tmallProductSchemaAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    通过主账号登陆态分页查询子账号列表
*/
func (ability *Defaultability) TaobaoSellercenterSubusersPage(req *request.TaobaoSellercenterSubusersPageRequest,session string) (*response.TaobaoSellercenterSubusersPageResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.sellercenter.subusers.page",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSellercenterSubusersPageResponse{}
    if(err != nil){
        log.Println("taobaoSellercenterSubusersPage error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    分页获取指定账户的子账号简易信息列表（新isv建议使用taobao.sellercenter.subusers.page接口）
*/
func (ability *Defaultability) TaobaoSubusersPage(req *request.TaobaoSubusersPageRequest,session string) (*response.TaobaoSubusersPageResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.subusers.page",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSubusersPageResponse{}
    if(err != nil){
        log.Println("taobaoSubusersPage error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    isv资源查询
*/
func (ability *Defaultability) CainiaoCloudprintIsvResourcesGet(req *request.CainiaoCloudprintIsvResourcesGetRequest) (*response.CainiaoCloudprintIsvResourcesGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("cainiao.cloudprint.isv.resources.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.CainiaoCloudprintIsvResourcesGetResponse{}
    if(err != nil){
        log.Println("cainiaoCloudprintIsvResourcesGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询卖家用户信息
*/
func (ability *Defaultability) TaobaoUserSellerGet(req *request.TaobaoUserSellerGetRequest,session string) (*response.TaobaoUserSellerGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.user.seller.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoUserSellerGetResponse{}
    if(err != nil){
        log.Println("taobaoUserSellerGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    云打印模板迁移接口
*/
func (ability *Defaultability) CainiaoCloudprintTemplatesMigrate(req *request.CainiaoCloudprintTemplatesMigrateRequest,session string) (*response.CainiaoCloudprintTemplatesMigrateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("cainiao.cloudprint.templates.migrate",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.CainiaoCloudprintTemplatesMigrateResponse{}
    if(err != nil){
        log.Println("cainiaoCloudprintTemplatesMigrate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    根据当前子账号登陆态，获取该子账号基本信息
*/
func (ability *Defaultability) TaobaoSubusersInfoQuery(req *request.TaobaoSubusersInfoQueryRequest,session string) (*response.TaobaoSubusersInfoQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.subusers.info.query",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSubusersInfoQueryResponse{}
    if(err != nil){
        log.Println("taobaoSubusersInfoQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取可用宝贝描述规范化模块
*/
func (ability *Defaultability) TaobaoItemAnchorGet(req *request.TaobaoItemAnchorGetRequest,session string) (*response.TaobaoItemAnchorGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.anchor.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemAnchorGetResponse{}
    if(err != nil){
        log.Println("taobaoItemAnchorGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    特殊退款类型的纠纷单列表查询
*/
func (ability *Defaultability) TaobaoSpecialRefundsReceiveGet(req *request.TaobaoSpecialRefundsReceiveGetRequest,session string) (*response.TaobaoSpecialRefundsReceiveGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.special.refunds.receive.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSpecialRefundsReceiveGetResponse{}
    if(err != nil){
        log.Println("taobaoSpecialRefundsReceiveGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    同意退款
*/
func (ability *Defaultability) TaobaoRpRefundsAgree(req *request.TaobaoRpRefundsAgreeRequest,session string) (*response.TaobaoRpRefundsAgreeResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.rp.refunds.agree",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoRpRefundsAgreeResponse{}
    if(err != nil){
        log.Println("taobaoRpRefundsAgree error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    卖家同意退货
*/
func (ability *Defaultability) TaobaoRpReturngoodsAgree(req *request.TaobaoRpReturngoodsAgreeRequest,session string) (*response.TaobaoRpReturngoodsAgreeResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.rp.returngoods.agree",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoRpReturngoodsAgreeResponse{}
    if(err != nil){
        log.Println("taobaoRpReturngoodsAgree error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    图片获取
*/
func (ability *Defaultability) TaobaoPicturePicturesGet(req *request.TaobaoPicturePicturesGetRequest,session string) (*response.TaobaoPicturePicturesGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.picture.pictures.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoPicturePicturesGetResponse{}
    if(err != nil){
        log.Println("taobaoPicturePicturesGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    图片总数查询
*/
func (ability *Defaultability) TaobaoPicturePicturesCount(req *request.TaobaoPicturePicturesCountRequest,session string) (*response.TaobaoPicturePicturesCountResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.picture.pictures.count",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoPicturePicturesCountResponse{}
    if(err != nil){
        log.Println("taobaoPicturePicturesCount error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
