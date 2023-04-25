<template>
    <div style="margin:20px;">
        <!-- fileList决定前端展示那些图片或者文件 -->
        <!-- action和customRequest用其一 -->
        <a-upload
            v-model:file-list="fileList"
            list-type="picture"
            @change="onChange"
            :customRequest="customRequest">
            <a-button>
                <upload-outlined></upload-outlined>
                上传
            </a-button>
        </a-upload>
    </div>
</template>

<script>
import {ref} from 'vue'
export default({
    setup() {
        const fileList = ref([{
            uid: '-1',
            name: 'xxx.png',
            status: 'done',
            url:'https://zos.alipayobjects.com/rmsportal/jkjgkEfvpUPVyRjUImniVslZfWPnJuuZ.png',
            thumbUrl:'https://zos.alipayobjects.com/rmsportal/jkjgkEfvpUPVyRjUImniVslZfWPnJuuZ.png',
            }, {
            uid: '-2',
            name: 'yyy.png',
            status: 'done',
            // url用于展示完整图片
            url:'https://zos.alipayobjects.com/rmsportal/jkjgkEfvpUPVyRjUImniVslZfWPnJuuZ.png',
            // thumbUrl用于展示图片的缩略图
            thumbUrl:'https://zos.alipayobjects.com/rmsportal/jkjgkEfvpUPVyRjUImniVslZfWPnJuuZ.png11',
        }]);

        function onChange() {
            //console.log(fileList)
        }
        //data是上传文件的所有信息和内容
        function customRequest(data) {
            console.log(data)
        }
        return {
            fileList,
            onChange,
            customRequest
        }
    },
})
</script>
