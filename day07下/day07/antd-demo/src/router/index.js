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
    },
    {
        path: '/layout',
        component: () => import('@/views/layout/layout.vue')
    },
    {
        path: '/menu',
        component: () => import('@/views/layout/menu.vue')
    },
    {
        path: '/form',
        component: () => import('@/views/form/form.vue')
    },
    {
        path: '/validate',
        component: () => import('@/views/form/validate.vue')
    },
    {
        path: '/custom_validate',
        component: () => import('@/views/form/custom_validate.vue')
    },
    {
        path: '/upload',
        component: () => import('@/views/form/upload.vue')
    },
    {
        path: '/table',
        component: () => import('@/views/table/table.vue')
    },
    {
        path: '/select',
        component: () => import('@/views/table/select.vue')
    },
    {
        path: '/sort',
        component: () => import('@/views/table/sort.vue')
    },
    {
        path: '/pagination',
        component: () => import('@/views/table/pagination.vue')
    },
    {
        path: '/drawer',
        component: () => import('@/views/feedback/drawer.vue')
    },
    {
        path: '/message',
        component: () => import('@/views/feedback/message.vue')
    },
    {
        path: '/modal',
        component: () => import('@/views/feedback/modal.vue')
    },
    {
        path: '/popover',
        component: () => import('@/views/feedback/popover.vue')
    }
]

//创建路由实例
const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router