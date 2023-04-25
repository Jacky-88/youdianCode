import { createApp } from 'vue'
import App from './App.vue'
//引入ant design和主题样式
import Antd from 'ant-design-vue'
import 'ant-design-vue/dist/antd.css'
//全局引入antd图标
import * as Icons from '@ant-design/icons-vue'
//引入router
import router from './router'

const vm = createApp(App)
//引入图标, i是组件名，Icons[i]是具体的组件
for (const i in Icons) {
    vm.component(i, Icons[i])
}
//引入antd组件库
vm.use(Antd)
vm.use(router)
vm.mount('#app')
