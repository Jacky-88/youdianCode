<template>
    <div style="margin:20px;width:500px;">
        <a-form ref="formRef" :model="formState" :label-col="labelCol">
            <!-- 输入框 -->
            <!-- rules参数： required是否必填，message提示信息 -->
            <a-form-item
                label="Activity name"
                name="name"
                :rules="[{required: true, message: '请输入Activity name'}]">
                <a-input v-model:value="formState.name"></a-input>
            </a-form-item>
            <!-- 开关 -->
            <a-form-item 
                label="Instance delivery"
                name="delivery"
                :rules="[{required: true, message: '请选择Instance delivery'}]">
                <a-switch v-model:checked="formState.delivery" />
            </a-form-item>
            <!-- 多选 -->
            <a-form-item
                label="Activity Type"
                name="type"
                :rules="[{required: true, message: '请选择Activity Type'}]">
                <a-checkbox-group v-model:value="formState.type">
                    <a-checkbox value="1">Online</a-checkbox>
                    <a-checkbox value="2">Promotion</a-checkbox>
                    <a-checkbox value="3">Offline</a-checkbox>
                </a-checkbox-group>
            </a-form-item>
            <!-- 单选 -->
            <a-form-item
                label="Resources"
                name="resource"
                :rules="[{required: true, message: '请选择Resources'}]">
                <a-radio-group v-model:value="formState.resource">
                    <a-radio value="1">Sponsor</a-radio>
                    <a-radio value="2">Venue</a-radio>
                </a-radio-group>
            </a-form-item>
            <a-form-item
                label="Activity Desc"
                name="desc"
                :rules="[{required: true, message: '请输入Activity Desc'}]">
                <a-textarea v-model:value="formState.desc"></a-textarea>
            </a-form-item>
            <!-- 定义操作按钮 -->
            <a-form-item :wrapper-col="{span:14, offset: 8}">
                <a-button type="primary" @click="formSubmit()">Submit</a-button>
                <a-button style="margin-left:10px;" type="primary" ghost @click="resetForm()">Reset</a-button>
                <a-button style="margin-left:10px;">Cancel</a-button>
            </a-form-item>
        </a-form>
    </div>
</template>

<script>
import { reactive, ref } from "@vue/reactivity"

export default ({
    setup() {
        const formState = reactive({
            name: '',
            delivery: '',
            type: [],
            resource: '',
            desc: ''
        })
        const labelCol = reactive({
            style: {
                width: '150px'
            }
        })
        const formRef = ref()
        //async和await用于方法的同步执行
        //async用于修饰在哪个方法里面使用await
        async function formSubmit() {
            console.log(formState)
            //捕捉方法同步执行的异常
            try {
                //实际等待方法执行完成
                await formRef.value.validateFields()
            } catch (error) {
                console.log('Failed:', error)
            }
        }
        //表单重置
        //重置常用于新增数据成功后，重置表单，这样下次打开，表单就是空的
        function resetForm() {
            formRef.value.resetFields()
        }
        return {
            formState,
            labelCol,
            formRef,
            formSubmit,
            resetForm
        }
    },
})
</script>
