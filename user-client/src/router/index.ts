import { createRouter, createWebHistory } from "vue-router";
import Page1 from "../views/Page1.vue";
import Page2 from "../views/Page2.vue";
import ScoreView from "../views/ScoreView.vue";
import ChatView from "../views/ChatView.vue";
const routes = [
  {
    path: "/",
    component: ScoreView,
  },
  {
    path: "/1",
    component: Page1,
  },
  {
    path: "/2",
    component: Page2,
  },
  {
    path: "/chat",
    component: ChatView,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
