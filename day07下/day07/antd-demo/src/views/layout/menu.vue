<template>
    <div>
        <a-menu
            v-model:selectedKeys="selectedKeys"
            style="width: 256px; height:100vh"
            mode="inline"
            :open-keys="openKeys"
            @openChange="onOpenChange">
            <a-sub-menu key="sub1">
                <!-- 两个插槽，定义icon与标题 -->
                <template #icon>
                <MailOutlined />
                </template>
                <template #title>Navigation One</template>
                <a-menu-item key="1">Option 1</a-menu-item>
                <a-menu-item key="2">Option 2</a-menu-item>
                <a-menu-item key="3">Option 3</a-menu-item>
                <a-menu-item key="4">Option 4</a-menu-item>
            </a-sub-menu>
            <a-sub-menu key="sub2">
                <template #icon></template>
                <template #title>
                <AppstoreOutlined />
                Navigation Two
                </template>
                <a-menu-item key="5">Option 5</a-menu-item>
                <a-menu-item key="6">Option 6</a-menu-item>
                <a-sub-menu key="sub3" title="Submenu">
                <a-menu-item key="7">Option 7</a-menu-item>
                <a-menu-item key="8">Option 8</a-menu-item>
                </a-sub-menu>
            </a-sub-menu>
            <a-sub-menu key="sub4">
                <template #icon>
                <SettingOutlined />
                </template>
                <template #title>Navigation Three</template>
                <a-menu-item key="9">Option 9</a-menu-item>
                <a-menu-item key="10">Option 10</a-menu-item>
                <a-menu-item key="11">Option 11</a-menu-item>
                <a-menu-item key="12">Option 12</a-menu-item>
            </a-sub-menu>
            <a-menu-item key="13">Option 13</a-menu-item>
        </a-menu>
    </div>
</template>

<script>
import { reactive, toRefs } from 'vue'
export default({
    setup() {
        const state = reactive({
            selectedKeys: ['1'],
            openKeys: ['sub1']
        })
        function onOpenChange(openKeys) {
            console.log(openKeys)
        }
        return {
            //将对象的属性直接返回出去，并带有ref属性
            ...toRefs(state),
            onOpenChange
        }
    },
})
</script>
