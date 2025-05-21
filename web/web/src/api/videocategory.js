import service from '@/utils/request'


// 查询分类信息
export const loadCategory = () => {
    return service.get('/videocategory/find')
}



// 修改视频分类状态
export const updateStatus = (data) => {
    return service({
      url: '/videocategory/update/status',
      method: 'post',
      data
    })
  }
  