<template>
    <div style="margin:20px">
        <a-table :columns="columns" :data-source="data" :pagination="pagination" @change="handleTableChange"/>
    </div>
</template>
<script>
import { reactive, ref } from 'vue';
export default ({
    setup() {
        //分页器
        const pagination = reactive({
            showSizeChanger: true,  //展示单页条数
            showQuickJumper: true, //快速跳转
            total: 100,  //总共多少条数据
            currentPage: 1, //第几页
            pageSize: 10, //单页大小
            pageSizeOptions: ["10", "20", "50", "100"], //单页大小
            showTotal: total => `共${total}条`
        })
        //分页的回调函数
        function handleTableChange(val) {
            pagination.currentPage = val.current,
            pagination.pageSize = val.pageSize
            //getList()
        }
        const columns = ref([{
            title: 'Name',
            dataIndex: 'name',
        }, {
            title: 'Chinese Score',
            dataIndex: 'chinese',
            sorter: {
                compare: (a, b) => a.chinese - b.chinese,
                multiple: 3,
            },
        }, {
            title: 'Math Score',
            dataIndex: 'math',
            sorter: {
                compare: (a, b) => a.math - b.math,
                multiple: 2,
            },
        }, {
            title: 'English Score',
            dataIndex: 'english',
            sorter: {
                compare: (a, b) => a.english - b.english,
                multiple: 1,
            },
        }])
        const data = ref([{
            key: '1',
            name: 'John Brown',
            chinese: 98,
            math: 60,
            english: 70,
        }, {
            key: '2',
            name: 'Jim Green',
            chinese: 98,
            math: 66,
            english: 89,
        }, {
            key: '3',
            name: 'Joe Black',
            chinese: 98,
            math: 90,
            english: 70,
        }, {
            key: '4',
            name: 'Jim Red',
            chinese: 88,
            math: 99,
            english: 89,
        }])
        return {
            data,
            columns,
            pagination,
            handleTableChange
        }
    },
});
</script>