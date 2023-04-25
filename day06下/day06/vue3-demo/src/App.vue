<template>
  <img alt="Vue logo" src="./assets/logo.png">
  <!-- 局部注册组件的使用 -->
  <!-- 子传父的事件方法，不能使用（），因为这样就默认不接收参数了 -->
  <!-- <HelloWorld :msg="text" @childMsg="getContent"/> -->
  <!-- 全局注册的组件，直接使用即可，不需要再import了 -->
  <Test1></Test1>
  <button @click="getData()">获取后端数据</button>
  {{ data }}
  <p v-if="error" style="color: red;">连接服务器失败，请稍后再试！</p>
</template>

<script>
//导入组件
//import HelloWorld from './components/HelloWorld.vue'
import {onMounted, ref} from 'vue'
//这里导入的是axios实例的别名
//import axios from 'axios'
import aaa from '@/request'
export default {
  name: 'App',
  //局部注册
  // components: {
  //   HelloWorld
  // },
  setup() {
    const text = ref('Hello, Vue!')
    const data = ref()
    const error = ref(false)
    //emit事件函数，接收emit子传父的数据
    function getContent(val1,val2) {
      console.log(val1,val2)
    }
    //axios发起请求
    function getData() {
      aaa.get('/get')
      .then(res => {
        data.value = res.data
      })
      .catch(res => {
        console.log(res)
        error.value = true
      })
      //.finally()
    }
    //生命周期钩子
    onMounted(() => {
      getData()
    })
    return {
      text,
      getContent,
      data,
      getData,
      error
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
