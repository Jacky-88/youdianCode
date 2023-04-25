<template>
  <div style="margin: 20px">
    <!-- 没有标题的卡片 -->
    <a-card style="width: 300px" :bodyStyle="{ padding: '50px' }">
      <p>这是一张卡片</p>
      <p>这是一张卡片</p>
      <p>这是一张卡片</p>
    </a-card>
    <!-- 有标题没边框的卡片 -->
    <a-card
      style="width: 300px"
      title="卡片标题"
      :bordered="false"
      :bodyStyle="{ padding: '50px' }"
    >
      <p>这是一张卡片</p>
      <p>这是一张卡片</p>
      <p>这是一张卡片</p>
    </a-card>
    <!--带扩展项的卡片 -->
    <a-card
      style="width: 300px"
      title="卡片标题"
      :bodyStyle="{ padding: '5px' }"
    >
      <!-- 这里是个插槽，extra是插槽的名字 -->
      <!-- 插槽就是灵活的定义某个区域的内容，定义玩之后会自动插入到这个区域 -->
      <template #extra>
        <a-button type="link">more</a-button>
      </template>
      <p>这是一张卡片</p>
      <p>这是一张卡片</p>
      <p>这是一张卡片</p>
    </a-card>
  </div>
</template>