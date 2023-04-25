<template>
    <div style="margin:20px;">
        <a-space>
            <a-button type="primary" @click="success">Success</a-button>
            <a-button type="primary" @click="error">Error</a-button>
            <a-button type="primary" @click="warning">Warning</a-button>
        </a-space>
    </div>
</template>

<script>
// 这里的message提示，实际上是antd提供的js功能
import { message } from 'ant-design-vue'
export default ({
    setup() {
        function success() {
            message.success({
                content: () => '这里是成功的提示',
                class: 'custom-class',
                style: {
                    marginTop: '20vh',
                }
            })
        }
        function error() {
            message.error('这里是错误的提示')
        }
        function warning() {
            message.warning('这里是警告的提示')
        }
        return {
            success,
            error,
            warning
        }
    },
})
</script>

<style scoped>
    .custom-class {
        color: red;
    }
</style>
