<template>
  <div class="p-[15px] bg-clr-user-chat-tint-70">
    <!-- Profile picture and text -->
    <div class="flex">
      <!-- Profile Picture -->
      <div
        class="w-[60px] h-[60px] rounded-full flex items-center justify-center shadow-xl bg-clr-user-chat-tint-30"
        style="flex-shrink: 0"
      >
        <UserIcon :size="35" color="white" />
      </div>
      <!-- Chat title and content -->
      <div class="ml-[5px]">
        <p class="font-bold text-[25px]">You</p>
        <p class="ml-[2px]">{{ props.content }}</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { defineProps } from "vue";
import UserIcon from "@/components/icons/UserIcon.vue";

const props = defineProps({
  content: {
    type: String,
    default: "...",
  },
});
</script>
