<template>
    <div>
        <button @click="btn()">Test</button>
    </div>
</template>

<script>
export default({
    setup() {
        function btn() {
            alert("你好，这里是测试组件")
        }
        return {
            btn
        }
    },
})
</script>
