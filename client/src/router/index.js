import Vue from "vue";
import VueRouter from "vue-router";
import SessionUtil from "@/plugins/session";
Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    redirect: "login",
  },
  {
    path: "/login",
    name: "login",
    component: () => import("@/views/Login.vue"),
  },
  {
    path: "/invoices",
    name: "invoices",
    component: () => import("@/views/Invoices.vue"),
  },
  {
    path: "/invoices/:invoiceId",
    name: "invoice",
    component: () => import("@/views/Invoice.vue"),
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

router.beforeEach((to, from, next) => {
  if (to.name !== "login" && !SessionUtil.isLogged()) next({ name: "login" });
  else next();
});

export default router;
