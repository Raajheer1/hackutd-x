<template>
  <div
    class="rounded-lg shadow-md p-5 mb-[5px] flex bg-[#fcfbf5]"
    v-touch
    @click="$emit('selectAction')"
  >
    <!-- Profile Picture -->
    <div
      class="w-[60px] h-[60px] rounded-full flex items-center justify-center"
      :class="{
        'bg-clr-good':
          props.actionItem.type == 'default' ||
          props.actionItem.type == 'saftey' ||
          props.actionItem.type == 'finance',
        'bg-clr-mid': props.actionItem.type == 'tax',
        'bg-clr-bad': props.actionItem.type == 'legal',
      }"
      style="flex-shrink: 0"
    >
      <ArchiveIcon
        :size="30"
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
        :size="33"
        color="white"
        v-if="props.actionItem.type == 'saftey'"
      />
    </div>
    <!-- Content -->
    <div class="ml-2">
      <p class="text-[20px] font-bold">
        {{ props.actionItem.title }}
      </p>
      <p class="text-gray-500">
        {{ truncatedDetails }}
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { defineProps, computed } from "vue";
import GavelIcon from "@/components/icons/GavelIcon.vue";
import BankIcon from "@/components/icons/BankIcon.vue";
import ConeIcon from "@/components/icons/ConeIcon.vue";
import ArchiveIcon from "@/components/icons/ArchiveIcon.vue";
import { ActionItemType } from "@/types";
const props = defineProps({
  actionItem: {
    type: Object as () => ActionItemType,
    default: {
      type: "default",
      title: "",
      details: "",
    },
  },
});
const emits = defineEmits(["openChat", "selectAction"]);

const maxLength = 50;
const truncatedDetails = computed(() => {
  return props.actionItem.details.length > maxLength
    ? props.actionItem.details.substring(0, maxLength) + "..."
    : props.actionItem.details;
});
</script>
<style scoped>
.v-enter-active {
  animation: slide-up 0.3s;
}

.v-leave-active {
  animation: slide-up 0.3s reverse;
}

@keyframes slide-up {
  0% {
    transform: translateY(60%);
  }
  100% {
    transform: translateY(0px);
  }
}
</style>
