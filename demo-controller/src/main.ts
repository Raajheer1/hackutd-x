import { createApp } from "vue";
import { createPinia } from "pinia";
import "./style.css";
import App from "./App.vue";
import router from "./router";
// Vuetify
import "vuetify/styles";
import { createVuetify } from "vuetify";
import * as components from "vuetify/components";
import * as directives from "vuetify/directives";
import Vue3TouchEvents from "vue3-touch-events";

const vuetify = createVuetify({
  components,
  directives,
});
const pinia = createPinia();
const app = createApp(App);
app.use(router);
app.use(pinia);
app.use(vuetify);
app.use(Vue3TouchEvents, {
  touchClass: "active",
});
app.mount("#app");
