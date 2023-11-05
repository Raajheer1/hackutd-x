<template>
  <div class="flex w-[100vw]">
    <!-- Company Options -->
    <div class="w-[25%] px-[15px]">
      <div
        v-for="(company, index) in companies"
        :key="index"
        class="rounded-md shadow-md p-[20px]"
        :class="
          selectedCompany === index
            ? 'bg-[#F5F5F5] border-2 border-black scale-105'
            : ''
        "
        @click="selectedCompany = index"
      >
        <p class="text-[25px] font-bold">{{ company.name }}</p>
        <p>{{ company.risk }}</p>
      </div>
    </div>
    <!-- Sliders -->
    <div
      class="flex flex-col w-[75%] p-[20px] justify-center align-items"
    >
      <p class="font-bold -mb-[15px]">Value</p>
      <v-range-slider></v-range-slider>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
type Company = {
  name: string;
  risk: number;
};
const selectedCompany = ref(0);
const companies: Company[] = [
  {
    name: "Company 1",
    risk: 1,
  },
  {
    name: "Company 2",
    risk: 2,
  },
  {
    name: "Company 3",
    risk: 3,
  },
  {
    name: "Company 4",
    risk: 4,
  },
  {
    name: "Company 5",
    risk: 5,
  },
];

async function getOptions() {
  //   const res = await fetch("http://localhost:3000/v1/configs");
  //   const data = await res.json();
  //   console.log(data);
  const res = await fetch("http://localhost:3000/v1/configs");
  const data = await res.json();
  console.log(data);
}
async function getOptions();
</script>
