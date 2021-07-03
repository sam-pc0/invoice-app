import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import pluginsIndex from "./plugins/pluginsIndex.js";

Vue.config.productionTip = false;

new Vue({
  ...pluginsIndex,
  router,
  render: (h) => h(App),
}).$mount("#app");
