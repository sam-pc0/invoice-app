import Vue from "vue";
import Buefy from "buefy";
import VueToast from "vue-toast-notification";
import VAnimateCss from "v-animate-css";

import "vue-toast-notification/dist/theme-sugar.css";
import "@/assets/sass/index.scss";
import "buefy/dist/buefy.css";

Vue.use(VueToast, {
  position: "top",
});

Vue.use(Buefy);
Vue.use(VAnimateCss);
