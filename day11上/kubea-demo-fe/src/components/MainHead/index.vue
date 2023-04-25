<template>
    <div>
        <a-card :bodyStyle="{padding: '10px'}">
            <a-row>
                <a-col :span="20">
                    <div style="text-align:left;">
                        <!-- 选择框 -->
                        <!-- 当收到父组件传过来的namespace属性为true时,才展示这个选择框 -->
                        <span v-if="namespace" style="font-size:14px;">命名空间：</span>
                        <a-select
                            v-if="namespace"
                            show-search
                            style="width:140px;margin-right:20px;"
                            size="small"
                            v-model:value="namespaceValue"
                            placeholder="请选择"
                            @change="nsChange">
                            <a-select-option
                                v-for="item, index in namespaceList"
                                :key="index"
                                :value="item.metadata.name">
                                {{ item.metadata.name }}
                            </a-select-option>
                        </a-select>
                        <!-- 搜索框 -->
                        <a-input
                            allow-clear
                            @change="searchChange"
                            style="width:200px;margin-right:10px;"
                            v-model:value="searchValue"
                            size="small"
                            :placeholder="searchDescribe">
                        </a-input>
                        <a-button size="small" type="primary" ghost @click="$emit('dataList')">
                            <template #icon>
                                <SearchOutlined />
                            </template>
                            搜索
                        </a-button>
                    </div>
                </a-col>
                <a-col :span="4">
                    <div style="text-align:right;">
                        <a-button v-if="add" style="margin-right:10px" size="small" type="primary" ghost @click="$emit('addFunc')">
                            <template #icon>
                                <PlusOutlined />
                            </template>
                            新增
                        </a-button>
                        <a-button size="small" @click="$emit('dataList')">
                            <template #icon>
                                <UndoOutlined />
                            </template>
                            刷新
                        </a-button>
                    </div>
                </a-col>
            </a-row>
        </a-card>
    </div>
</template>

<script>
import { ref, reactive, onMounted } from 'vue';
import common from "@/config";
import httpClient from '@/request';
import { message } from 'ant-design-vue';
export default ({
    //定义父传子的属性
    name: 'MainHead',
    props: {
        namespace: {type: Boolean, default: false},
        searchDescribe: {type: String, default: ''},
        add: {type: Boolean, default: false},
    },
    //定义子传父的事件
    //namespaceChange 当namespace放生变化时，要把最新的namespace的值给到父组件
    //namespaceList 当deployment、service、ingress这些有创建功能的页面，创建时需要选择namespace
    //dataList 父组件传递获取列表的方法，这里去执行
    emits: ['searchChange', 'namespaceChange', 'namespaceList', 'dataList', 'addFunc'],
    setup(props, ctx) {
        const searchValue = ref('')
        const namespaceValue = ref('default')
        //命名空间列表
        const namespaceList = ref()
        const namespaceListData = reactive({
            url: common.k8sNamespaceList,
            params: {
                cluster: ''
            }
        })
        //【方法】
        function searchChange(val) {
            ctx.emit('searchChange', searchValue.value)
        }
        function nsChange(val) {
            namespaceValue.value = val
            localStorage.setItem('namespace', val)
            //子传父
            ctx.emit('namespaceChange', val)
            //重新获取一次资源列表
            ctx.emit('dataList')
        }
        function getNamespaceList() {
            namespaceListData.params.cluster =  localStorage.getItem('k8s_cluster')
            httpClient.get(namespaceListData.url, {params: namespaceListData.params})
            .then(res => {
                namespaceList.value = res.data.items
                //父组件要定义namespaceList事件对应的方法，获取到namespaceList.value
                ctx.emit('namespaceList', namespaceList.value)
            })
            .catch(res => {
                message.error(res.msg)
            })
        }
        onMounted(() => {
            //第一次进来是default的命名空间，当选了某个命名空间后，切换页面时，命名空间会继承
            if (localStorage.getItem('namespace')) {
                nsChange(localStorage.getItem('namespace'))
            } else {
                nsChange(namespaceValue.value)
            }
            getNamespaceList()
        })
        return {
            namespaceValue,
            namespaceList,
            nsChange,
            searchValue,
            searchChange
        }
    },
})
</script>
