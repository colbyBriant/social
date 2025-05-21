import service from '@/utils/request'

// @Tags BbsCategoryApi
// @Summary 
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bbs.BBSCategory true "/bbscategory/page"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbscategory/page [post]
export const loadBbsCategoryData = (data) => {
  return service({
    url: '/bbscategory/page',
    method: 'post',
    data
  })
}

export const findBbsCategory = () => {
  return service({
    url: '/bbscategory/find',
    method: 'get'
  })
}

// 保存帖子分类
export const saveBbsCategory = (data) => {
  return service({
    url: '/bbscategory/save',
    method: 'post',
    data
  })
}

// 根据id删除分类
export const delBbsCategory = (id) => {
  return service({
    url: `/bbscategory/delete/${id}`,
    method: 'delete',
  })
}

// 批量删除
export const delBbsCategorys = (ids) => {
  return service({
    url: `/bbscategory/deletes?ids=${ids}`,
    method: 'delete',
  })
}

// 修改帖子分类
export const updateBbsCategory = (data) => {
  return service({
    url: '/bbscategory/update',
    method: 'post',
    data
  })
}


// 修改帖子分类状态
export const updateStatus = (data) => {
  return service({
    url: '/bbscategory/update/status',
    method: 'post',
    data
  })
}
