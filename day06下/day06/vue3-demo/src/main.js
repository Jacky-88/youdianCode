// 导入：{}导入的是组件的方法
import { createApp } from 'vue'
// 导入：导入组件的别名
import App from './App.vue'
// @代表在src目录下
import Test from '@/components/Test.vue'

const vm = createApp(App)
//基于vue实例，全局注册
vm.component('Test1', Test)
vm.mount('#app')