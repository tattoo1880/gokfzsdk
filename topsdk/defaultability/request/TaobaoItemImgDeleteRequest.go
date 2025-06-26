package request


type TaobaoItemImgDeleteRequest struct {
    /*
        商品数字ID     */
    NumIid  *int64 `json:"num_iid" required:"true" `
    /*
        商品图片ID；如果是竖图，请将id的值设置为1     */
    Id  *int64 `json:"id" required:"true" `
    /*
        标记是否要删除第6张图，因为第6张图与普通商品图片不是存储在同一个位置的无图片ID，所以要通过一个标记来判断是否为第6张图，目前第6张图业务主要用在女装业务下 defalutValue��false    */
    IsSixthPic  *bool `json:"is_sixth_pic,omitempty" required:"false" `
}

func (s *TaobaoItemImgDeleteRequest) SetNumIid(v int64) *TaobaoItemImgDeleteRequest {
    s.NumIid = &v
    return s
}
func (s *TaobaoItemImgDeleteRequest) SetId(v int64) *TaobaoItemImgDeleteRequest {
    s.Id = &v
    return s
}
func (s *TaobaoItemImgDeleteRequest) SetIsSixthPic(v bool) *TaobaoItemImgDeleteRequest {
    s.IsSixthPic = &v
    return s
}

func (req *TaobaoItemImgDeleteRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.NumIid != nil) {
        paramMap["num_iid"] = *req.NumIid
    }
    if(req.Id != nil) {
        paramMap["id"] = *req.Id
    }
    if(req.IsSixthPic != nil) {
        paramMap["is_sixth_pic"] = *req.IsSixthPic
    }
    return paramMap
}

func (req *TaobaoItemImgDeleteRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}