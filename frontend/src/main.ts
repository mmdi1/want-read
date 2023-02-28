import { createApp } from 'vue'
import App from './App.vue'
import naive from 'naive-ui'
let app = createApp(App)
app.use(naive) // 挂载路由
app.mount('#app')
