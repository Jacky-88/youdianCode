<template>
    <div style="margin:20px;">
        <a-space>
            <!-- 标签组件的用法 -->
            <a-button type="primary" @click="showModal">Modal</a-button>
            <!-- js的用法 -->
            <a-button @click="confirm">Confirm</a-button>
        </a-space>
        <a-modal
            v-model:visible="visible"
            title="Modal"
            ok-text="确认"
            cancel-text="取消"
            @ok="handleModal">
            <p>这里是Modal的内容</p>
            <p>这里是Modal的内容</p>
            <p>这里是Modal的内容</p>
        </a-modal>
    </div>
</template>

<script>
import {ref, createVNode} from 'vue'
import {Modal} from 'ant-design-vue'
import {ExclamationCircleOutlined} from '@ant-design/icons-vue'
export default({
    setup() {
        const visible =ref(false)
        function showModal() {
            visible.value = true
        }
        function handleModal() {
            visible.value = false
        }
        function confirm() {
            //Modal的确认框
            Modal.confirm({
                title: 'Confirm',
                //使用createVNode去定义modal提示的图标
                //createVNode用于返回html
                // content: createVNode('div', {
                //   style: 'color:red;',
                // }, 'Some descriptions'),
                icon: createVNode(ExclamationCircleOutlined),
                content: '这里是Modal的内容',
                okText: '确认',
                cancelText: '取消',
                onOk() {
                    console.log('ok')
                },
                onCancel() {
                    console.log('cancel')
                }
            })
        }
        return {
            visible,
            showModal,
            handleModal,
            confirm
        }
    },
})
</script>
