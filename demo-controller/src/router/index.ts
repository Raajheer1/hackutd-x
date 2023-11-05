import { createRouter, createWebHistory } from "vue-router";
import Page1 from "../views/Page1.vue";
import Page2 from "../views/Page2.vue";
import SelectView from "@/views/SelectView.vue";
const routes = [
  {
    path: "/",
    component: SelectView,
  },
  {
    path: "/1",
    component: Page1,
  },
  {
    path: "/2",
    component: Page2,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
