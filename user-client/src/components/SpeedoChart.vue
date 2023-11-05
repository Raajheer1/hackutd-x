<template>
  <canvas ref="scoreChart" id="score-chart"></canvas>
</template>
<script setup lang="ts">
import Chart from "chart.js/auto";
import { onMounted, watch, ref } from "vue";

const props = defineProps({
  v1: {
    type: Number,
    default: 0,
  },
  v2: {
    type: Number,
    default: 0,
  },
  v3: {
    type: Number,
    default: 0,
  },
});
//Make computed values so the chart can be reactive
const v1 = ref(props.v1);
const v2 = ref(props.v2);
const v3 = ref(props.v3);
const data = {
  labels: ["Red", "Yellow", "Blue"],
  datasets: [
    {
      label: "My First Dataset",
      data: [v1.value, v2.value, v3.value],
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
let chart: Chart<"doughnut", number[], string>;
onMounted(() => {
  ctx = document.getElementById(
    "score-chart"
  ) as HTMLCanvasElement;
  chart = new Chart(ctx, {
    type: "doughnut",
    data: data,
    options: {
      plugins: {
        legend: {
          display: false,
          position: "bottom",
          align: "center",
        },
      },
    },
  });
});
watch(
  () => props,
  () => {
    console.log("Updating chart");
    chart.data.datasets[0].data = [
      v1.value,
      v2.value,
      v3.value,
    ];
    chart.update();
  }
);
</script>
