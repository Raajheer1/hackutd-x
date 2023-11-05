<template>
  <div class="text-3xl font-bold">Page 1</div>
  <div class="text-xl">Count: {{ counter.count }}</div>
  <button class="m-1" @click="routeToPage2">Page 2</button>
  <button class="m-1" @click="incrementCounter">Count</button>
</template>

<script lang="ts" setup>
import { useCounterStore } from "@/stores/index.ts";
import { useRouter } from "vue-router";
const router = useRouter();
const counter = useCounterStore();

function routeToPage2() {
  router.push("/2");
}
function incrementCounter() {
  counter.increment();
}
</script>
<style scoped></style>
