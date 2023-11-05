<template>
  <div class="text-3xl font-bold">Page 2</div>
  <div class="text-xl">Count: {{ counter.count }}</div>
  <button class="m-1" @click="routeToPage1">Page 1</button>
  <button class="m-1" @click="incrementCounter">Count</button>
</template>

<script lang="ts" setup>
import { useCounterStore } from "@/stores/index.ts";
import { useRouter } from "vue-router";
const router = useRouter();
const counter = useCounterStore();

function routeToPage1() {
  router.push("/1");
}
function incrementCounter() {
  counter.increment();
}
</script>
<style scoped></style>
