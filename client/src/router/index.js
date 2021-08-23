import Vue from "vue";
import VueRouter from "vue-router";
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

export default router;
