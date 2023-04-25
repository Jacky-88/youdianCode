import { createApp } from 'vue'
import App from './App.vue'
//引入router
import router from './router'
//引入ant
import Antd from 'ant-design-vue'
//引入ant暗黑风格主题以及图标
import 'ant-design-vue/dist/antd.dark.css'
import * as Icons from '@ant-design/icons-vue'
import CButton from '@/components/CButton'
//codemirror编辑器
import { GlobalCmComponent } from "codemirror-editor-vue3";
// 引入主题 可以从 codemirror/theme/ 下引入多个
import 'codemirror/theme/dracula.css'
// 引入语言模式 可以从 codemirror/mode/ 下引入多个
import 'codemirror/mode/yaml/yaml.js'

const app = createApp(App)
//图标注册全局组件
for (const i in Icons) {
    app.component(i, Icons[i])
}
app.component('c-button', CButton)
app.use(GlobalCmComponent, { componentName: "codemirror" });
//引入antd
app.use(Antd)
//引入router
app.use(router)
app.mount('#app')
