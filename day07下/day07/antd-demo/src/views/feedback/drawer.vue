<template>
    <div>
        <a-button type="primary" @click="showDrawer">Open</a-button>
        <a-drawer
            v-model:visible="visible"
            style="color:red;"
            title="这是一个抽屉的标题"
            placement="right"
            :bodyStyle="{padding: 50 + 'px'}">
            <p>这是抽屉的内容。。。</p>
            <p>这是抽屉的内容。。。</p>
            <p>这是抽屉的内容。。。</p>
        </a-drawer>
    </div>
</template>

<script>
import { ref } from 'vue'
export default({
    setup() {
        const visible = ref(false)
        function showDrawer() {
            visible.value = true
        }
        return {
            visible,
            showDrawer
        }
    },
})
</script>
