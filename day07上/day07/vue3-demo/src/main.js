// 导入：{}导入的是组件的方法
import { createApp } from 'vue'
// 导入：导入组件的别名
import App from './App.vue'
// @代表在src目录下
import Test from '@/components/Test.vue'
//引入vue-router
import router from './router'

const vm = createApp(App)
//全局引入vue-router实例
vm.use(router)
//基于vue实例，全局注册
vm.component('Test1', Test)
vm.mount('#app')