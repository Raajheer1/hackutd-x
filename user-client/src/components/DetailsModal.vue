<template>
  <div
    class="z-5 fixed top-[64px] inset-x-0 h-[100%] border-t rounded-t-2xl bg-gradient-to-b to-white"
    :class="{
      'from-clr-good-shade-10':
        props.actionItem.type == 'default' ||
        props.actionItem.type == 'finance' ||
        props.actionItem.type == 'saftey',
      'from-clr-mid': props.actionItem.type == 'tax',
      'from-clr-bad-tint-30': props.actionItem.type == 'legal',
    }"
    @click.stop="() => $emit('closeModal')"
  >
    <!-- Header (Profile pic, title, X button) -->
    <div class="flex p-[10px]">
      <!-- Profile Picture -->
      <div
        class="w-[80px] h-[80px] rounded-full flex items-center justify-center shadow-xl"
        :class="{
          'bg-clr-good':
            props.actionItem.type == 'default' ||
            props.actionItem.type == 'finance' ||
            props.actionItem.type == 'saftey',
          'bg-clr-mid-shade-7': props.actionItem.type == 'tax',
          'bg-clr-bad': props.actionItem.type == 'legal',
        }"
        style="flex-shrink: 0"
      >
        <ArchiveIcon
          :size="45"
          color="white"
          v-if="
            props.actionItem.type == 'default' ||
            props.actionItem.type == 'tax'
          "
        />
        <GavelIcon
          :size="50"
          color="white"
          v-if="props.actionItem.type == 'legal'"
        />
        <BankIcon
          :size="35"
          color="white"
          v-if="props.actionItem.type == 'finance'"
        />
        <ConeIcon
          :size="45"
          color="white"
          v-if="props.actionItem.type == 'saftey'"
        />
      </div>
      <!-- Title -->
      <div>
        <p class="text-[30px] font-bold ml-[15px] mt-[15px]">
          {{ props.actionItem.title }}
        </p>
      </div>
      <!-- Close button
      <CloseIcon
        :size="50"
        color="black"
        class="ml-auto mt-[11px] -mr-[6px]"
      /> -->
    </div>
    <!-- Details -->
    <p class="px-[25px] text-[20px]">
      {{ props.actionItem.details }}
    </p>
    <!-- Buttons -->
    <div class="flex items-center mt-[60px]">
      <div
        class="flex h-[75px] w-[75px] rounded-xl flex items-center justify-center shadow-xl mr-auto ml-[30px]"
        :class="buttonColor"
        v-touch
        @click="() => $emit('complete')"
      >
        <!-- <p class="font-bold text-[25px]">Complete</p> -->
        <CheckIcon :size="40" color="black" />
      </div>
      <div
        class="w-[90px] h-[90px] rounded-full flex items-center justify-center shadow-xl ml-auto mr-[30px]"
        :class="buttonColor"
        v-touch
        @click.stop="() => $emit('openChat')"
      >
        <ChatIcon :size="45" color="black" class="mt-[5px]" />
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { defineProps, defineEmits, computed } from "vue";
import GavelIcon from "@/components/icons/GavelIcon.vue";
import BankIcon from "@/components/icons/BankIcon.vue";
import ConeIcon from "@/components/icons/ConeIcon.vue";
import ArchiveIcon from "@/components/icons/ArchiveIcon.vue";
import CheckIcon from "@/components/icons/CheckIcon.vue";
import ChatIcon from "@/components/icons/ChatIcon.vue";
import { ActionItemType } from "@/types";
const props = defineProps({
  actionItem: {
    type: Object as () => ActionItemType,
    default: {
      type: "default",
      title: "Default",
      details: "Default",
    },
  },
});
const emits = defineEmits([
  "openChat",
  "closeModal",
  "complete",
]);

const buttonColor = computed(() => {
  if (
    props.actionItem.type == "default" ||
    props.actionItem.type == "finance" ||
    props.actionItem.type == "saftey"
  ) {
    return "bg-clr-good";
  } else if (props.actionItem.type == "tax") {
    return "bg-clr-mid";
  } else if (props.actionItem.type == "legal") {
    return "bg-clr-bad";
  }
});
</script>
