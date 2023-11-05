<template>
  <div class="flex w-[100vw]">
    <!-- Company Options -->
    <div class="w-[25%] px-[15px]">
      <div
        v-for="(company, index) in companies"
        :key="index"
        class="rounded-md shadow-md p-[20px]"
        :class="
          selectedCompanyIndex === index
            ? 'bg-[#F5F5F5] border-2 border-black scale-105'
            : ''
        "
        @click="
          () => {
            selectedCompany = company;
            selectedCompanyIndex = index;
          }
        "
      >
        <p class="text-[25px] font-bold">
          {{ company.companyName }}
        </p>
        <p>{{ company.presetLevel }}</p>
      </div>
    </div>
    <!-- Sliders -->
    <div
      class="flex flex-col flex-grow p-[20px] pt-0 justify-center align-items"
    >
      <div v-for="(value, key) in selectedCompany" :key="key">
        <div v-if="typeof value == 'number'">
          <p class="font-bold -mb-[15px]">
            {{ key }} - {{ value }}
          </p>
          <v-slider
            v-model="selectedCompany[key]"
            :color="sliderColors[0]"
            :max="1000 * 1000"
            :min="0"
            :step="1000"
          ></v-slider>
        </div>
      </div>
    </div>
    <!-- Submit -->
    <div class="flex flex-col justify-end p-[15px] pb-[20px]">
      <button
        class="bg-[#f57fce] rounded-md shadow-md p-[20px] font-bold"
        v-touch
        @click="submitOption"
      >
        Submit
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
type Company = {
  companyName: string;
  industry: string;
  city: string;
  revenue: number;
  cogs: number;
  depreciation: number;
  longTermAssets: number;
  shortTermAssets: number;
  longTermLiability: number;
  shortTermLiability: number;
  operatingExpense: number;
  retainedEarnings: number;
  yearsInBusiness: number;
  presetLevel: string;
};
const selectedCompany = ref<Company>({} as Company);
const selectedCompanyIndex = ref<number>(0);
const companies = ref<Company[]>([] as Company[]);

async function getOptions() {
  //   const res = await fetch("http://localhost:3000/v1/configs");
  //   const data = await res.json();
  //   console.log(data);
  const res = await fetch(
    import.meta.env.VITE_API_URL + "/v1/configs"
  );
  const data = await res.json();
  companies.value = data;
  selectedCompany.value = data[0];
  console.log(typeof data[0]["revenue"]);
}
getOptions();

async function submitOption() {
  console.log(selectedCompany.value);
  const res = await fetch(
    import.meta.env.VITE_API_URL + "/v1/set-config",
    {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(selectedCompany.value),
    }
  );
  const data = await res.json();
  console.log(data);
}
const sliderColors = ["#7fcef5", "#f5ed7f", "#f57fce"];
</script>
