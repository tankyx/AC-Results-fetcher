import Vue from 'vue'
import App from './App.vue'
import router from './router'
import 'beautify-scrollbar/dist/index.css';
import 'v2-table/dist/index.css';
import V2Table from 'v2-table';

Vue.config.productionTip = false;
Vue.use(V2Table)

var track = new Vue({
  data: {
    name: ''
  },
  template: '<div>{{ message }}</div>'
})

track.name = "Rainbow Road"

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
