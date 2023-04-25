<template>
    <div style="text-align:center">
        <a-layout style="height:100vh">
            <a-layout-header>Header</a-layout-header>
            <!-- 在a-layout中，定义时垂直排列的布局 -->
            <a-layout>
                <!-- 第一种方式 -->
                <!-- <a-layout-sider>Sider</a-layout-sider>
                <a-layout>
                    <a-layout-content>Content</a-layout-content>
                    <a-layout-footer>Footer</a-layout-footer>
                </a-layout> -->
                <!-- 第二种方式 -->
                <a-layout>
                    <a-layout-sider>Sider</a-layout-sider>
                    <a-layout-content>Content</a-layout-content>
                </a-layout>
                <a-layout-footer>Footer</a-layout-footer>
            </a-layout>
        </a-layout>
    </div>
</template>

<style scoped>
    .ant-layout-header {
        background: rgb(131, 202, 235);
    }
    .ant-layout-sider {
        background: rgb(90, 168, 204);
    }
    .ant-layout-content {
        background: rgb(61, 136, 172);
    }
    .ant-layout-footer {
        background: rgb(23, 86, 116);
    }
</style>