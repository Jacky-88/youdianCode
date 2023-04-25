import { createApp } from 'vue'
import App from './App.vue'
//引入router
import router from './router'
//引入ant
import Antd from 'ant-design-vue'
//引入ant暗黑风格主题以及图标
import 'ant-design-vue/dist/antd.dark.css'
import * as Icons from '@ant-design/icons-vue'
import CButton from '@/components/CButton'

const app = createApp(App)
//图标注册全局组件
for (const i in Icons) {
    app.component(i, Icons[i])
}
app.component('c-button', CButton)
//引入antd
app.use(Antd)
//引入router
app.use(router)
app.mount('#app')
