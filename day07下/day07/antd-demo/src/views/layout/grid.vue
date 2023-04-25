<template>
    <div>
        <!-- row中的col水平和垂直的间隔都是10px -->
        <!-- justify定義col的水平方向的位置 -->
        <a-row :gutter="[10, 10]" justify="center">
            <a-col :span="6">
                <div class="box" style="height:200px;">col-6</div>
            </a-col>
            <a-col :span="6">
                <div class="box">col-6</div>
            </a-col>
            
            <a-col :span="6">
                <div class="box">col-6</div>
            </a-col>
            <a-col :span="6">
                <div class="box">col-6</div>
            </a-col>
            <a-col :span="6">
                <div class="box">col-6</div>
            </a-col>
            <a-col :span="6">
                <div class="box">col-6</div>
            </a-col>
        </a-row>
    </div>
</template>

<style scoped>
    .box {
        background: blue;
        color:aliceblue;
    }
</style>