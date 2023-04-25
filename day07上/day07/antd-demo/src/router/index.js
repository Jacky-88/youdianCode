import {createRouter, createWebHistory} from 'vue-router'

//路由规则
const routes = [
    {
        path: '/btn',
        component: () => import('@/views/test/btn.vue')
    },
    {
        path: '/card',
        component: () => import('@/views/layout/card.vue')
    },
    {
        path: '/affix',
        component: () => import('@/views/layout/affix.vue')
    },
    {
        path: '/grid',
        component: () => import('@/views/layout/grid.vue')
    }
]

//创建路由实例
const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router