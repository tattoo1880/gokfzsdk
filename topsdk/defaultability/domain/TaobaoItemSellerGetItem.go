package domain

import (
        "topsdk/util"
    )

type TaobaoItemSellerGetItem struct {
    /*
        售后服务ID,该字段仅在taobao.item.get接口中返回     */
    AfterSaleId  *int64 `json:"after_sale_id,omitempty" `

    /*
        商品上传后的状态。onsale出售中，instock库中     */
    ApproveStatus  *string `json:"approve_status,omitempty" `

    /*
        商品的积分返点比例。如:5,表示:返点比例0.5%     */
    AuctionPoint  *int64 `json:"auction_point,omitempty" `

    /*
        代充商品类型。在代充商品的类目下，不传表示不标记商品类型（交易搜索中就不能通过标记搜到相关的交易了）。可选类型： no_mark(不做类型标记) time_card(点卡软件代充) fee_card(话费软件代充)     */
    AutoFill  *string `json:"auto_fill,omitempty" `

    /*
        自动重发,true/false     */
    AutoRepost  *bool `json:"auto_repost,omitempty" `

    /*
        商品级别的条形码     */
    Barcode  *string `json:"barcode,omitempty" `

    /*
        商品所属的叶子类目 id     */
    Cid  *int64 `json:"cid,omitempty" `

    /*
        货到付款运费模板ID     */
    CodPostageId  *int64 `json:"cod_postage_id,omitempty" `

    /*
        Item的发布时间，目前仅供taobao.item.add和taobao.item.get可用     */
    Created  *util.LocalTime `json:"created,omitempty" `

    /*
        定制工具Id     */
    CustomMadeTypeId  *string `json:"custom_made_type_id,omitempty" `

    /*
        下架时间（格式：yyyy-MM-dd HH:mm:ss）     */
    DelistTime  *util.LocalTime `json:"delist_time,omitempty" `

    /*
        商品描述, 字数要大于5个字符，小于25000个字符     */
    Desc  *string `json:"desc,omitempty" `

    /*
        宝贝描述规范化模块锚点信息     */
    DescModuleInfo  *TaobaoItemSellerGetDescModuleInfo `json:"desc_module_info,omitempty" `

    /*
        商品url     */
    DetailUrl  *string `json:"detail_url,omitempty" `

    /*
        ems费用,格式：5.00；单位：元；精确到：分     */
    EmsFee  *string `json:"ems_fee,omitempty" `

    /*
        快递费用,格式：5.00；单位：元；精确到：分     */
    ExpressFee  *string `json:"express_fee,omitempty" `

    /*
        宝贝特征值，只有在Top支持的特征值才能保存到宝贝上     */
    Features  *string `json:"features,omitempty" `

    /*
        食品安全信息，包括：生产许可证编号、产品标准号、厂名、厂址等     */
    FoodSecurity  *TaobaoItemSellerGetFoodSecurity `json:"food_security,omitempty" `

    /*
        运费承担方式,seller（卖家承担），buyer(买家承担）     */
    FreightPayer  *string `json:"freight_payer,omitempty" `

    /*
        全球购商品采购地信息（地区/国家），代表全球购商品的产地信息。     */
    GlobalStockCountry  *string `json:"global_stock_country,omitempty" `

    /*
        全球购商品采购地信息（库存类型），有两种库存类型：现货和代购;参数值为1时代表现货，值为2时代表代购     */
    GlobalStockType  *string `json:"global_stock_type,omitempty" `

    /*
        支持会员打折,true/false     */
    HasDiscount  *bool `json:"has_discount,omitempty" `

    /*
        是否有发票,true/false     */
    HasInvoice  *bool `json:"has_invoice,omitempty" `

    /*
        橱窗推荐,true/false     */
    HasShowcase  *bool `json:"has_showcase,omitempty" `

    /*
        是否有保修,true/false     */
    HasWarranty  *bool `json:"has_warranty,omitempty" `

    /*
        加价幅度。如果为0，代表系统代理幅度。在竞拍中，为了超越上一个出价，会员需要在当前出价上增加金额，这个金额就是加价幅度。卖家在发布宝贝的时候可以自定义加价幅度，也可以让系统自动代理加价。系统自动代理加价的加价幅度随着当前出价金额的增加而增加，我们建议会员使用系统自动代理加价，并请买家在出价前看清楚加价幅度的具体金额。另外需要注意是，此功能只适用于拍卖的商品。以下是系统自动代理加价幅度表：当前价（加价幅度 ）1-40（ 1 ）、41-100（ 2 ）、101-200（5 ）、201-500 （10）、501-1001（15）、001-2000（25）、2001-5000（50）、5001-10000（100）10001以上         200     */
    Increment  *string `json:"increment,omitempty" `

    /*
        用户内店宝贝装修模板id     */
    InnerShopAuctionTemplateId  *int64 `json:"inner_shop_auction_template_id,omitempty" `

    /*
        用户自行输入的类目属性ID串。结构："pid1,pid2,pid3"，如："20000"（表示品牌） 注：通常一个类目下用户可输入的关键属性不超过1个。     */
    InputPids  *string `json:"input_pids,omitempty" `

    /*
        用户自行输入的子属性名和属性值，结构:"父属性值;一级子属性名;一级子属性值;二级子属性名;自定义输入值,....",如：“耐克;耐克系列;科比系列;科比系列;2K5”，input_str需要与input_pids一一对应，注：通常一个类目下用户可输入的关键属性不超过1个。所有属性别名加起来不能超过 3999字节。     */
    InputStr  *string `json:"input_str,omitempty" `

    /*
        是否是3D淘宝的商品     */
    Is3D  *bool `json:"is_3D,omitempty" `

    /*
        是否在外部网店显示     */
    IsEx  *bool `json:"is_ex,omitempty" `

    /*
        普通商品:0;代销商品:1;经销商品:2;供销产品:3     */
    IsFenxiao  *int64 `json:"is_fenxiao,omitempty" `

    /*
        是否24小时闪电发货     */
    IsLightningConsignment  *bool `json:"is_lightning_consignment,omitempty" `

    /*
        商品是否为先行赔付taobao.items.search和taobao.items.vip.search专用     */
    IsPrepay  *bool `json:"is_prepay,omitempty" `

    /*
        是否在淘宝显示     */
    IsTaobao  *bool `json:"is_taobao,omitempty" `

    /*
        是否定时上架商品     */
    IsTiming  *bool `json:"is_timing,omitempty" `

    /*
        虚拟商品的状态字段     */
    IsVirtual  *bool `json:"is_virtual,omitempty" `

    /*
        标示商品是否为新品。值含义：true-是，false-否。     */
    IsXinpin  *bool `json:"is_xinpin,omitempty" `

    /*
        商品图片列表(包括主图)。fields中只设置item_img可以返回ItemImg结构体中所有字段，如果设置为item_img.id、item_img.url、item_img.position等形式就只会返回相应的字段     */
    ItemImgs  *[]TaobaoItemSellerGetItemImg `json:"item_imgs,omitempty" `

    /*
        表示商品的体积，用于按体积计费的运费模板。该值的单位为立方米（m3）。该值支持两种格式的设置：格式1：bulk:3,单位为立方米(m3),表示直接设置为商品的体积。格式2：weight:10;breadth:10;height:10，单位为米（m）     */
    ItemSize  *string `json:"item_size,omitempty" `

    /*
        商品的重量，用于按重量计费的运费模板。注意：单位为kg     */
    ItemWeight  *string `json:"item_weight,omitempty" `

    /*
        上架时间（格式：yyyy-MM-dd HH:mm:ss）     */
    ListTime  *util.LocalTime `json:"list_time,omitempty" `

    /*
        本地生活电子交易凭证业务，目前此字段只涉及到的信息为有效期:如果有效期为起止日期类型，此值为2012-08-06,2012-08-16如果有效期为【购买成功日 至】类型则格式为2012-08-16如果有效期为天数类型则格式为3     */
    LocalityLife  *TaobaoItemSellerGetLocalityLife `json:"locality_life,omitempty" `

    /*
        商品所在地     */
    Location  *TaobaoItemSellerGetLocation `json:"location,omitempty" `

    /*
        商品修改时间（格式：yyyy-MM-dd HH:mm:ss）     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

    /*
        宝贝主图视频的数据信息，包括：视频ID，视频缩略图URL，视频时长，视频状态等信息。     */
    MpicVideo  *TaobaoItemSellerGetMpicVideo `json:"mpic_video,omitempty" `

    /*
        是否为新消保法中的7天无理由退货     */
    Newprepay  *string `json:"newprepay,omitempty" `

    /*
        卖家昵称     */
    Nick  *string `json:"nick,omitempty" `

    /*
        商品数量     */
    Num  *int64 `json:"num,omitempty" `

    /*
        商品数字id     */
    NumIid  *int64 `json:"num_iid,omitempty" `

    /*
        是否淘1站商品     */
    OneStation  *bool `json:"one_station,omitempty" `

    /*
        商家外部编码(可与商家外部系统对接)。需要授权才能获取。     */
    OuterId  *string `json:"outer_id,omitempty" `

    /*
        用户外店装修模板id     */
    OuterShopAuctionTemplateId  *int64 `json:"outer_shop_auction_template_id,omitempty" `

    /*
        用于保存拍卖有关的信息     */
    PaimaiInfo  *TaobaoItemSellerGetPaimaiInfo `json:"paimai_info,omitempty" `

    /*
        周期销售库存     */
    PeriodSoldQuantity  *int64 `json:"period_sold_quantity,omitempty" `

    /*
        商品主图片地址     */
    PicUrl  *string `json:"pic_url,omitempty" `

    /*
        平邮费用,格式：5.00；单位：元；精确到：分     */
    PostFee  *string `json:"post_fee,omitempty" `

    /*
        宝贝所属的运费模板ID，如果没有返回则说明没有使用运费模板     */
    PostageId  *int64 `json:"postage_id,omitempty" `

    /*
        商品价格，格式：5.00；单位：元；精确到：分     */
    Price  *string `json:"price,omitempty" `

    /*
        宝贝所属产品的id(可能为空). 该字段可以通过taobao.products.search 得到     */
    ProductId  *int64 `json:"product_id,omitempty" `

    /*
        消保类型，多个类型以,分割。可取以下值：2：假一赔三；4：7天无理由退换货；taobao.items.search和taobao.items.vip.search专用     */
    PromotedService  *string `json:"promoted_service,omitempty" `

    /*
        商品属性图片列表。fields中只设置prop_img可以返回PropImg结构体中所有字段，如果设置为prop_img.id、prop_img.url、prop_img.properties、prop_img.position等形式就只会返回相应的字段     */
    PropImgs  *[]TaobaoItemSellerGetPropImg `json:"prop_imgs,omitempty" `

    /*
        属性值别名,比如颜色的自定义名称     */
    PropertyAlias  *string `json:"property_alias,omitempty" `

    /*
        商品属性 格式：pid:vid;pid:vid     */
    Props  *string `json:"props,omitempty" `

    /*
        商品属性名称。标识着props内容里面的pid和vid所对应的名称。格式为：pid1:vid1:pid_name1:vid_name1;pid2:vid2:pid_name2:vid_name2……(<strong>注：</strong><font color="red">属性名称中的冒号":"被转换为："#cln#";  分号";"被转换为："#scln#"</font>)     */
    PropsName  *string `json:"props_name,omitempty" `

    /*
        商品所属卖家的信用等级数，1表示1心，2表示2心……，只有调用商品搜索:taobao.items.get和taobao.items.search的时候才能返回     */
    Score  *int64 `json:"score,omitempty" `

    /*
        秒杀商品类型。打上秒杀标记的商品，用户只能下架并不能再上架，其他任何编辑或删除操作都不能进行。如果用户想取消秒杀标记，需要联系小二进行操作。如果秒杀结束需要自由编辑请联系活动负责人（小二）去掉秒杀标记。可选类型web_only(只能通过web网络秒杀)wap_only(只能通过wap网络秒杀)web_and_wap(既能通过web秒杀也能通过wap秒杀)     */
    SecondKill  *string `json:"second_kill,omitempty" `

    /*
        商品卖点信息，天猫商家使用字段，最长150个字符。     */
    SellPoint  *string `json:"sell_point,omitempty" `

    /*
        是否承诺退换货服务!     */
    SellPromise  *bool `json:"sell_promise,omitempty" `

    /*
        商品所属的店铺内卖家自定义类目列表     */
    SellerCids  *string `json:"seller_cids,omitempty" `

    /*
        <a href="http://open.taobao.com/dev/index.php/Sku">Sku</a>列表。fields中只设置sku可以返回Sku结构体中所有字段，如果设置为sku.sku_id、sku.properties、sku.quantity等形式就只会返回相应的字段     */
    Skus  *[]TaobaoItemSellerGetSku `json:"skus,omitempty" `

    /*
        商品销量     */
    SoldQuantity  *int64 `json:"sold_quantity,omitempty" `

    /*
        商品新旧程度(全新:new，闲置:unused，二手：second)     */
    StuffStatus  *string `json:"stuff_status,omitempty" `

    /*
        商品是否支持拍下减库存:1支持;2取消支持(付款减库存);0(默认)不更改 集市卖家默认拍下减库存; 商城卖家默认付款减库存     */
    SubStock  *int64 `json:"sub_stock,omitempty" `

    /*
        页面模板id     */
    TemplateId  *string `json:"template_id,omitempty" `

    /*
        商品标题,不能超过60字节     */
    Title  *string `json:"title,omitempty" `

    /*
        商品类型(fixed:一口价;auction:拍卖)注：取消团购     */
    Type  *string `json:"type,omitempty" `

    /*
        有效期,7或者14（默认是7天）     */
    ValidThru  *int64 `json:"valid_thru,omitempty" `

    /*
        该字段废弃，请勿使用。     */
    VideoId  *int64 `json:"video_id,omitempty" `

    /*
        商品视频列表(目前只支持单个视频关联)。fields中只设置video可以返回Video结构体中所有字段，如果设置为video.id、video.video_id、video.url等形式就只会返回相应的字段     */
    Videos  *[]TaobaoItemSellerGetVideo `json:"videos,omitempty" `

    /*
        商品是否违规，违规：true , 不违规：false     */
    Violation  *bool `json:"violation,omitempty" `

    /*
        对应搜索商品列表页的最近成交量,只有调用商品搜索:taobao.items.get和taobao.items.search的时候才能返回     */
    Volume  *int64 `json:"volume,omitempty" `

    /*
        不带html标签的desc文本信息，该字段只在taobao.item.get接口中返回     */
    WapDesc  *string `json:"wap_desc,omitempty" `

    /*
        适合wap应用的商品详情url ，该字段只在taobao.item.get接口中返回     */
    WapDetailUrl  *string `json:"wap_detail_url,omitempty" `

    /*
        无线的宝贝描述     */
    WirelessDesc  *string `json:"wireless_desc,omitempty" `

    /*
        预扣库存，即付款减库存的商品现在有多少处于未付款状态的订单     */
    WithHoldQuantity  *int64 `json:"with_hold_quantity,omitempty" `

    /*
        商品所属的商家的旺旺在线状况，taobao.items.search和taobao.items.vip.search专用     */
    WwStatus  *bool `json:"ww_status,omitempty" `

    /*
        商品描述模块化，模块列表，由List转化成jsonArray存入，后端逻辑验证通过，拼装成模块内容+锚点导航后存入desc中。数据结构具体参见Item_Desc_Module     */
    DescModules  *string `json:"desc_modules,omitempty" `

    /*
        村淘特有商品级数据结构     */
    CuntaoItemSpecific  *TaobaoItemSellerGetCuntaoItemSpecific `json:"cuntao_item_specific,omitempty" `

    /*
        门店大屏图     */
    LargeScreenImageUrl  *string `json:"large_screen_image_url,omitempty" `

    /*
        true:商品是区域限售商品；false:商品不是区域限售商品。     */
    IsAreaSale  *bool `json:"is_area_sale,omitempty" `

    /*
        属性值的备注.格式:pid:vid:备注信息1;pid2:vid2:备注信息2;     */
    CpvMemo  *string `json:"cpv_memo,omitempty" `

    /*
        3:4主图     */
    ItemRectangleImgs  *[]TaobaoItemSellerGetItemImg `json:"item_rectangle_imgs,omitempty" `

    /*
        商品首次上架时间     */
    FirstStartsTime  *util.LocalTime `json:"first_starts_time,omitempty" `

}

func (s *TaobaoItemSellerGetItem) SetAfterSaleId(v int64) *TaobaoItemSellerGetItem {
    s.AfterSaleId = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetApproveStatus(v string) *TaobaoItemSellerGetItem {
    s.ApproveStatus = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetAuctionPoint(v int64) *TaobaoItemSellerGetItem {
    s.AuctionPoint = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetAutoFill(v string) *TaobaoItemSellerGetItem {
    s.AutoFill = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetAutoRepost(v bool) *TaobaoItemSellerGetItem {
    s.AutoRepost = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetBarcode(v string) *TaobaoItemSellerGetItem {
    s.Barcode = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetCid(v int64) *TaobaoItemSellerGetItem {
    s.Cid = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetCodPostageId(v int64) *TaobaoItemSellerGetItem {
    s.CodPostageId = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetCreated(v util.LocalTime) *TaobaoItemSellerGetItem {
    s.Created = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetCustomMadeTypeId(v string) *TaobaoItemSellerGetItem {
    s.CustomMadeTypeId = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetDelistTime(v util.LocalTime) *TaobaoItemSellerGetItem {
    s.DelistTime = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetDesc(v string) *TaobaoItemSellerGetItem {
    s.Desc = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetDescModuleInfo(v TaobaoItemSellerGetDescModuleInfo) *TaobaoItemSellerGetItem {
    s.DescModuleInfo = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetDetailUrl(v string) *TaobaoItemSellerGetItem {
    s.DetailUrl = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetEmsFee(v string) *TaobaoItemSellerGetItem {
    s.EmsFee = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetExpressFee(v string) *TaobaoItemSellerGetItem {
    s.ExpressFee = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetFeatures(v string) *TaobaoItemSellerGetItem {
    s.Features = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetFoodSecurity(v TaobaoItemSellerGetFoodSecurity) *TaobaoItemSellerGetItem {
    s.FoodSecurity = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetFreightPayer(v string) *TaobaoItemSellerGetItem {
    s.FreightPayer = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetGlobalStockCountry(v string) *TaobaoItemSellerGetItem {
    s.GlobalStockCountry = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetGlobalStockType(v string) *TaobaoItemSellerGetItem {
    s.GlobalStockType = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetHasDiscount(v bool) *TaobaoItemSellerGetItem {
    s.HasDiscount = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetHasInvoice(v bool) *TaobaoItemSellerGetItem {
    s.HasInvoice = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetHasShowcase(v bool) *TaobaoItemSellerGetItem {
    s.HasShowcase = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetHasWarranty(v bool) *TaobaoItemSellerGetItem {
    s.HasWarranty = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetIncrement(v string) *TaobaoItemSellerGetItem {
    s.Increment = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetInnerShopAuctionTemplateId(v int64) *TaobaoItemSellerGetItem {
    s.InnerShopAuctionTemplateId = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetInputPids(v string) *TaobaoItemSellerGetItem {
    s.InputPids = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetInputStr(v string) *TaobaoItemSellerGetItem {
    s.InputStr = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetIs3D(v bool) *TaobaoItemSellerGetItem {
    s.Is3D = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetIsEx(v bool) *TaobaoItemSellerGetItem {
    s.IsEx = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetIsFenxiao(v int64) *TaobaoItemSellerGetItem {
    s.IsFenxiao = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetIsLightningConsignment(v bool) *TaobaoItemSellerGetItem {
    s.IsLightningConsignment = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetIsPrepay(v bool) *TaobaoItemSellerGetItem {
    s.IsPrepay = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetIsTaobao(v bool) *TaobaoItemSellerGetItem {
    s.IsTaobao = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetIsTiming(v bool) *TaobaoItemSellerGetItem {
    s.IsTiming = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetIsVirtual(v bool) *TaobaoItemSellerGetItem {
    s.IsVirtual = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetIsXinpin(v bool) *TaobaoItemSellerGetItem {
    s.IsXinpin = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetItemImgs(v []TaobaoItemSellerGetItemImg) *TaobaoItemSellerGetItem {
    s.ItemImgs = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetItemSize(v string) *TaobaoItemSellerGetItem {
    s.ItemSize = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetItemWeight(v string) *TaobaoItemSellerGetItem {
    s.ItemWeight = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetListTime(v util.LocalTime) *TaobaoItemSellerGetItem {
    s.ListTime = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetLocalityLife(v TaobaoItemSellerGetLocalityLife) *TaobaoItemSellerGetItem {
    s.LocalityLife = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetLocation(v TaobaoItemSellerGetLocation) *TaobaoItemSellerGetItem {
    s.Location = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetModified(v util.LocalTime) *TaobaoItemSellerGetItem {
    s.Modified = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetMpicVideo(v TaobaoItemSellerGetMpicVideo) *TaobaoItemSellerGetItem {
    s.MpicVideo = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetNewprepay(v string) *TaobaoItemSellerGetItem {
    s.Newprepay = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetNick(v string) *TaobaoItemSellerGetItem {
    s.Nick = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetNum(v int64) *TaobaoItemSellerGetItem {
    s.Num = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetNumIid(v int64) *TaobaoItemSellerGetItem {
    s.NumIid = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetOneStation(v bool) *TaobaoItemSellerGetItem {
    s.OneStation = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetOuterId(v string) *TaobaoItemSellerGetItem {
    s.OuterId = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetOuterShopAuctionTemplateId(v int64) *TaobaoItemSellerGetItem {
    s.OuterShopAuctionTemplateId = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetPaimaiInfo(v TaobaoItemSellerGetPaimaiInfo) *TaobaoItemSellerGetItem {
    s.PaimaiInfo = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetPeriodSoldQuantity(v int64) *TaobaoItemSellerGetItem {
    s.PeriodSoldQuantity = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetPicUrl(v string) *TaobaoItemSellerGetItem {
    s.PicUrl = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetPostFee(v string) *TaobaoItemSellerGetItem {
    s.PostFee = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetPostageId(v int64) *TaobaoItemSellerGetItem {
    s.PostageId = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetPrice(v string) *TaobaoItemSellerGetItem {
    s.Price = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetProductId(v int64) *TaobaoItemSellerGetItem {
    s.ProductId = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetPromotedService(v string) *TaobaoItemSellerGetItem {
    s.PromotedService = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetPropImgs(v []TaobaoItemSellerGetPropImg) *TaobaoItemSellerGetItem {
    s.PropImgs = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetPropertyAlias(v string) *TaobaoItemSellerGetItem {
    s.PropertyAlias = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetProps(v string) *TaobaoItemSellerGetItem {
    s.Props = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetPropsName(v string) *TaobaoItemSellerGetItem {
    s.PropsName = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetScore(v int64) *TaobaoItemSellerGetItem {
    s.Score = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetSecondKill(v string) *TaobaoItemSellerGetItem {
    s.SecondKill = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetSellPoint(v string) *TaobaoItemSellerGetItem {
    s.SellPoint = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetSellPromise(v bool) *TaobaoItemSellerGetItem {
    s.SellPromise = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetSellerCids(v string) *TaobaoItemSellerGetItem {
    s.SellerCids = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetSkus(v []TaobaoItemSellerGetSku) *TaobaoItemSellerGetItem {
    s.Skus = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetSoldQuantity(v int64) *TaobaoItemSellerGetItem {
    s.SoldQuantity = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetStuffStatus(v string) *TaobaoItemSellerGetItem {
    s.StuffStatus = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetSubStock(v int64) *TaobaoItemSellerGetItem {
    s.SubStock = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetTemplateId(v string) *TaobaoItemSellerGetItem {
    s.TemplateId = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetTitle(v string) *TaobaoItemSellerGetItem {
    s.Title = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetType(v string) *TaobaoItemSellerGetItem {
    s.Type = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetValidThru(v int64) *TaobaoItemSellerGetItem {
    s.ValidThru = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetVideoId(v int64) *TaobaoItemSellerGetItem {
    s.VideoId = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetVideos(v []TaobaoItemSellerGetVideo) *TaobaoItemSellerGetItem {
    s.Videos = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetViolation(v bool) *TaobaoItemSellerGetItem {
    s.Violation = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetVolume(v int64) *TaobaoItemSellerGetItem {
    s.Volume = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetWapDesc(v string) *TaobaoItemSellerGetItem {
    s.WapDesc = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetWapDetailUrl(v string) *TaobaoItemSellerGetItem {
    s.WapDetailUrl = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetWirelessDesc(v string) *TaobaoItemSellerGetItem {
    s.WirelessDesc = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetWithHoldQuantity(v int64) *TaobaoItemSellerGetItem {
    s.WithHoldQuantity = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetWwStatus(v bool) *TaobaoItemSellerGetItem {
    s.WwStatus = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetDescModules(v string) *TaobaoItemSellerGetItem {
    s.DescModules = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetCuntaoItemSpecific(v TaobaoItemSellerGetCuntaoItemSpecific) *TaobaoItemSellerGetItem {
    s.CuntaoItemSpecific = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetLargeScreenImageUrl(v string) *TaobaoItemSellerGetItem {
    s.LargeScreenImageUrl = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetIsAreaSale(v bool) *TaobaoItemSellerGetItem {
    s.IsAreaSale = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetCpvMemo(v string) *TaobaoItemSellerGetItem {
    s.CpvMemo = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetItemRectangleImgs(v []TaobaoItemSellerGetItemImg) *TaobaoItemSellerGetItem {
    s.ItemRectangleImgs = &v
    return s
}
func (s *TaobaoItemSellerGetItem) SetFirstStartsTime(v util.LocalTime) *TaobaoItemSellerGetItem {
    s.FirstStartsTime = &v
    return s
}
