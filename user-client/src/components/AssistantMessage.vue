<template>
  <div
    class="p-[15px]"
    :class="{
      'bg-clr-good-tint-50':
        props.type == 'default' ||
        props.type == 'financial' ||
        props.type == 'saftey',
      'bg-clr-mid': props.type == 'tax',
      'bg-clr-bad': props.type == 'legal',
    }"
  >
    <!-- Profile picture and text -->
    <div class="flex">
      <!-- Profile Picture -->
      <div
        class="w-[60px] h-[60px] rounded-full flex items-center justify-center shadow-xl"
        :class="{
          'bg-clr-good':
            props.type == 'default' ||
            props.type == 'finance' ||
            props.type == 'saftey',
          'bg-clr-mid-shade-7': props.type == 'tax',
          'bg-clr-bad': props.type == 'legal',
        }"
        style="flex-shrink: 0"
      >
        <ArchiveIcon
          :size="35"
          color="white"
          v-if="props.type == 'default' || props.type == 'tax'"
        />
        <GavelIcon
          :size="40"
          color="white"
          v-if="props.type == 'legal'"
        />
        <BankIcon
          :size="35"
          color="white"
          v-if="props.type == 'financial'"
        />
        <ConeIcon
          :size="33"
          color="white"
          v-if="props.type == 'saftey'"
        />
      </div>
      <!-- Chat title and content -->
      <div class="ml-[5px]">
        <p class="font-bold text-[25px]">{{ props.title }}</p>
        <p class="ml-[2px]">
          Lets talk about your reccomendation:
          {{ props.details }}
        </p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import GavelIcon from "@/components/icons/GavelIcon.vue";
import BankIcon from "@/components/icons/BankIcon.vue";
import ConeIcon from "@/components/icons/ConeIcon.vue";
import ArchiveIcon from "@/components/icons/ArchiveIcon.vue";
const props = defineProps({
  type: {
    type: String,
    default: "default",
  },
  title: {
    type: String,
    default: "Title",
  },
  details: {
    type: String,
    default: "Details",
  },
});
</script>
