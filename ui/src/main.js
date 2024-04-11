import {createApp} from "vue";
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import router from "@/components/router/router.ts";
import App from "./App.vue";
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import zhCn from 'element-plus/dist/locale/zh-cn.mjs'
import {runConfig} from "@/components/mod/config.ts";

const app = createApp(App)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)

}
app.use(ElementPlus, {
    locale: zhCn,
})
app.use(router)
app.mount('#app')

export  let baseUrl = "http://10.5.10.87:55001"
if (import.meta.env.PROD){
    runConfig.server =  window.location.origin
}else {
    runConfig.server = baseUrl
}