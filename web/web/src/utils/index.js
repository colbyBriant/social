import { ElMessageBox } from 'element-plus'
import { ElMessage } from 'element-plus'
// type : success ,error，warning，message
// export const KsdMessage = (message, type = 'success') => {
//   return ElMessage({
//     message,
//     type
//   })
// }

export const KsdBox = {
    confirm(title = '提示', content = '请输入内容', options={cbtn:'取消',sbtn:'确认',type:'warning'}){
        return ElMessageBox.confirm(content, title, { confirmButtonText: options.sbtn, cancelButtonText: options.cbtn, type: options.type});
    }
}
export const KsdMessage = {
  main(message, type) {
    return ElMessage({message,type})
  },
  success(message) {
    this.main(message, 'success')
  },
  warn(message) {
    this.main(message, 'warning')
  },
  message(message) {
    this.main(message, 'message')
  },
  error(message) {
    this.main(message, 'error')
  }
}

export default { ...KsdBox, ...KsdMessage }