<template>
    <div style="margin:20px;width:500px;">
        <a-form :model="formState" :label-col="labelCol">
            <!-- 输入框 -->
            <a-form-item label="Activity name">
                <a-input v-model:value="formState.name"></a-input>
            </a-form-item>
            <!-- 开关 -->
            <a-form-item label="Instance delivery">
                <a-switch v-model:checked="formState.delivery" />
            </a-form-item>
            <!-- 多选 -->
            <a-form-item label="Activity Type">
                <a-checkbox-group v-model:value="formState.type">
                    <a-checkbox value="1">Online</a-checkbox>
                    <a-checkbox value="2">Promotion</a-checkbox>
                    <a-checkbox value="3">Offline</a-checkbox>
                </a-checkbox-group>
            </a-form-item>
            <!-- 单选 -->
            <a-form-item label="Resources">
                <a-radio-group v-model:value="formState.resource">
                    <a-radio value="1">Sponsor</a-radio>
                    <a-radio value="2">Venue</a-radio>
                </a-radio-group>
            </a-form-item>
            <a-form-item label="Activity Desc">
                <a-textarea v-model:value="formState.desc"></a-textarea>
            </a-form-item>
        </a-form>
    </div>
</template>

<script>
import { reactive } from "@vue/reactivity"

export default ({
    setup() {
        const formState = reactive({
            name: '',
            delivery: false,
            type: [],
            resource: '',
            desc: ''
        })
        const labelCol = reactive({
            style: {
                width: '150px'
            }
        })
        return {
            formState,
            labelCol
        }
    },
})
</script>
