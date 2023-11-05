import { createApp } from "vue";
import { createPinia } from "pinia";
import "./style.css";
import App from "./App.vue";
import router from "./router";
import Vue3TouchEvents from "vue3-touch-events";

const pinia = createPinia();
const app = createApp(App);
app.use(router);
app.use(pinia);
app.use(Vue3TouchEvents, {
  touchClass: "active",
});
app.mount("#app");
