package domain


type TaobaoContentMediaUploadGetGetMediaUploadResultResponse struct {
    /*
        文件id     */
    FileId  *string `json:"file_id,omitempty" `

    /*
        类型     */
    MimeType  *string `json:"mime_type,omitempty" `

    /*
        是否完成上传     */
    Done  *bool `json:"done,omitempty" `

}

func (s *TaobaoContentMediaUploadGetGetMediaUploadResultResponse) SetFileId(v string) *TaobaoContentMediaUploadGetGetMediaUploadResultResponse {
    s.FileId = &v
    return s
}
func (s *TaobaoContentMediaUploadGetGetMediaUploadResultResponse) SetMimeType(v string) *TaobaoContentMediaUploadGetGetMediaUploadResultResponse {
    s.MimeType = &v
    return s
}
func (s *TaobaoContentMediaUploadGetGetMediaUploadResultResponse) SetDone(v bool) *TaobaoContentMediaUploadGetGetMediaUploadResultResponse {
    s.Done = &v
    return s
}
