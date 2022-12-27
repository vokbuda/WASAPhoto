import {createApp, reactive} from 'vue'
import App from './App.vue'



import router from './router'
import axios from './services/axios.js';
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'
import { ObserveVisibility } from 'vue-observe-visibility';
import './assets/dashboard.css'
import './assets/main.css'


const app = createApp(App)

app.directive('observe-visibility', {
  beforeMount: (el, binding, vnode) => {
    vnode.context = binding.instance;
    ObserveVisibility.bind(el, binding, vnode);
  },
  update: ObserveVisibility.update,
  unmounted: ObserveVisibility.unbind,
});



app.config.globalProperties.$axios = axios;
app.component("ErrorMsg", ErrorMsg);
app.component("LoadingSpinner", LoadingSpinner);
app.use(router)

app.mount('#app')

