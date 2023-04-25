import axios from 'axios'

const instance = axios.create({
    baseURL: 'http://www.httpbin.org',
    timeout: 3000,
})
//拦截器，请求拦截
instance.interceptors.request.use(config => {
    //在请求发送之前做些什么
    //例如添加请求头
    config.headers['Test-Header'] = '123456'
    return config
    },
    error => {
        //处理请求错误
        return Promise.reject(error)
    }
)

//拦截器，响应拦截
instance.interceptors.response.use(
    response => {
        return response
    },
    error => {
        console.log(error, 'error')
        return Promise.reject(error)   // 返回接口返回的错误信息
    }
)

export default instance