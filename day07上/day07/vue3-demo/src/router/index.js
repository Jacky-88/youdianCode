//导入vue原生方法
//1、创建路由方法
//2、路由类型
import {createRouter, createWebHistory} from 'vue-router'
//导入组件方式，先导入，下面再引入
//注意：一开始就会引入，即使不进入页面
import Home from '@/views/Home.vue'

//定义路由策略
const routes = [
    {
        path: '/home', //访问路径
        name: 'Home', //路由名称
        component: Home  //引入组件
    },
    {
        path: '/about',
        name: 'About',
        component: () => import('@/views/About.vue') //引入组件，懒加载，只有在打开页面的时候才加载这个组件
    }
]

//创建路由实例
const router = createRouter({
    //路由模式 createWebHistory createWebHashHistory
    history: createWebHistory(),
    routes
})

//前置路由守卫,在进入页面之前执行
router.beforeEach((to, from, next) => {
    console.log(to,"去哪里")
    console.log(from,"从哪里来")
    //假装这里的home是login页面
    if (to.path == '/home') {
        return next()
    }
    const token = '' //模拟token，实际场景中是从cookie或者localstorage中获取
    if (token) {
        return next()
    } else {
        return next('/home') // 如果没有token，就跳转到登录页
    }
})

//后置路由守卫,在进入页面之后执行
//router.afterEach()


export default router