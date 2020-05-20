import App from './App'
import router from './router'
import VueClipboard from 'vue-clipboard2'

Vue.use(VueClipboard)
Vue.prototype.axios = axios

new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
