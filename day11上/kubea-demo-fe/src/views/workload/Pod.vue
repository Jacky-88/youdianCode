<template>
    <div>
        <MainHead
            namespace
            searchDescribe="请输入"
            @searChange="getSearchValue"
            @namespaceChange="getNamespaceValue"
            @dataList="getPodList">
        </MainHead>
    </div>
</template>

<script>
//子组件的局部引入
import MainHead from '@/components/MainHead';
export default({
    components: {
        MainHead
    },
    setup() {
        function getSearchValue(val) {}
        function getNamespaceValue(val) {}
        function getPodList() {}
        return {
            getSearchValue,
            getNamespaceValue,
            getPodList
        }
    },
})
</script>
