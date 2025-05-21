package req

type BbsCategorySaveReq struct {
	ID          uint   `json:"id" form:"id"`
	Title       string `form:"title" json:"title"`
	Description string `form:"description" json:"description"`
	Parent_id   uint   `form:"parentId" json:"parentId"`
	Sorted      int    `form:"sorted" json:"sorted"`
	Status      int8   `form:"status" json:"status"`
	IsDelete    int8   `form:"isDelete" json:"isDelete"`
}
