package request


type AlibabaItemOperateDeleteRequest struct {
    /*
        商品ID     */
    ItemId  *int64 `json:"item_id" required:"true" `
}

func (s *AlibabaItemOperateDeleteRequest) SetItemId(v int64) *AlibabaItemOperateDeleteRequest {
    s.ItemId = &v
    return s
}

func (req *AlibabaItemOperateDeleteRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    return paramMap
}

func (req *AlibabaItemOperateDeleteRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}