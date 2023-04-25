<template>
    <div style="margin:20px;text-align:center;">
        <a-button type="primary">Primary Button</a-button>
        <br>
        <a-button>Default Button</a-button>
        <br>
        <a-button danger>Danger Button</a-button>
        <br>
        <a-button type="dashed">Dashed Button</a-button>
        <br>
        <a-button type="text">Text Button</a-button>
        <br>
        <a-button type="link">Link Button</a-button>
    </div>
</template>

<style scoped>
    /* 给组件添加样式使用class时，一般都是ant-开头 */
    .ant-btn {
        margin-bottom: 10px;
    }
</style>