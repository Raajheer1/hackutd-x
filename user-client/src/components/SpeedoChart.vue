<template>
  <canvas ref="scoreChart" id="score-chart"></canvas>
</template>
<script setup lang="ts">
import Chart from "chart.js/auto";
import { onMounted } from "vue";

const props = defineProps({
  v1: {
    type: Number,
    default: 50,
  },
  v2: {
    type: Number,
    default: 50,
  },
  v3: {
    type: Number,
    default: 100,
  },
});
const data = {
  labels: ["Red", "Yellow", "Blue"],
  datasets: [
    {
      label: "My First Dataset",
      data: [props.v1, props.v2, props.v3],
      backgroundColor: ["#f57fce", "#f5ed7f", "#7fcef5"],
      borderColor: "#faf9f6",
      borderRadius: 10,
      borderWidth: 3,
      radius: 150,
      cutout: 130,
      hoverOffset: 4,
      rotation: 125,
      circumference: 110,
    },
  ],
};
let ctx: HTMLCanvasElement;
onMounted(() => {
  ctx = document.getElementById(
    "score-chart"
  ) as HTMLCanvasElement;
  new Chart(ctx, {
    type: "doughnut",
    data: data,
  });
});
</script>
