<template>
    <div style="margin:200px;">
        <!-- 默认的弹出框 -->
        <a-popover>
            <template #content>
                <p>这里是popover的内容</p>
                <p>这里是popover的内容</p>
                <p>这里是popover的内容</p>
            </template>
            <a-button type="primary">Hover me</a-button>
        </a-popover>
        <!-- 从上面弹出来 -->
        <a-popover title="这里是popover的标题" placement="bottom">
            <template #content>
                <p>这里是popover的内容</p>
                <p>这里是popover的内容</p>
                <p>这里是popover的内容</p>
            </template>
            <a-button type="primary">Hover me2</a-button>
        </a-popover>
        <!-- 触发条件为点击 -->
        <a-popover title="这里是popover的标题" trigger="click">
            <template #content>
                <p>这里是popover的内容</p>
                <p>这里是popover的内容</p>
                <p>这里是popover的内容</p>
            </template>
            <a-button type="primary">Hover me3</a-button>
        </a-popover>
    </div>
</template>