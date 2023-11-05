<template>
  <div
    class="overflow-y-auto z-10 w-[100vw] h-[100vh] fixed top-[10px] left-0 bg-app-bg-clr"
  >
    <!-- Header -->
    <div
      class="flex fixed top-0 pt-[55px] pb-[10px] h-max items-center justify-center w-[100vw] bg-app-bg-clr"
    >
      <!-- Back button -->
      <div class="flex-1" v-touch @click="$emit('closeChat')">
        <LeftArrowIcon
          class="ml-[5px]"
          :size="40"
          color="black"
        />
      </div>

      <!-- Title -->
      <div class="">
        <p class="text-[30px] font-bold">Chat</p>
      </div>
      <!-- Chat Icon -->
      <div class="flex flex-1">
        <ChatIcon
          :size="40"
          color="black"
          class="ml-auto mr-[10px] mt-[5px]"
        />
      </div>
    </div>
    <!-- Chat Messages -->
    <div class="mt-[100px] mb-[100px]">
      <!-- Assistant message -->
      <div
        class="p-[15px]"
        :class="{
          'bg-clr-good-tint-50':
            props.actionItem.type == 'default' ||
            props.actionItem.type == 'financial' ||
            props.actionItem.type == 'saftey',
          'bg-clr-mid': props.actionItem.type == 'tax',
          'bg-clr-bad': props.actionItem.type == 'legal',
        }"
      >
        <!-- Profile picture and text over reccomendation display -->
        <div class="flex flex-col">
          <!-- Profile picture and text -->
          <div class="flex">
            <!-- Profile Picture -->
            <div
              class="w-[60px] h-[60px] rounded-full flex items-center justify-center shadow-xl"
              :class="{
                'bg-clr-good':
                  props.actionItem.type == 'default' ||
                  props.actionItem.type == 'finance' ||
                  props.actionItem.type == 'saftey',
                'bg-clr-mid-shade-7':
                  props.actionItem.type == 'tax',
                'bg-clr-bad': props.actionItem.type == 'legal',
              }"
              style="flex-shrink: 0"
            >
              <ArchiveIcon
                :size="35"
                color="white"
                v-if="
                  props.actionItem.type == 'default' ||
                  props.actionItem.type == 'tax'
                "
              />
              <GavelIcon
                :size="40"
                color="white"
                v-if="props.actionItem.type == 'legal'"
              />
              <BankIcon
                :size="35"
                color="white"
                v-if="props.actionItem.type == 'financial'"
              />
              <ConeIcon
                :size="33"
                color="white"
                v-if="props.actionItem.type == 'saftey'"
              />
            </div>
            <!-- Chat title and content -->
            <div class="ml-[5px]">
              <p class="font-bold text-[25px]">
                {{ props.actionItem.title }}
              </p>
              <p>Lets talk about your reccomendation:</p>
            </div>
          </div>
          <!-- Reccomendation displaybox -->
          <div
            class="p-[10px] rounded-xl my-[10px]"
            :class="{
              'bg-clr-good-tint-10':
                props.actionItem.type == 'default' ||
                props.actionItem.type == 'financial' ||
                props.actionItem.type == 'saftey',
              'bg-clr-mid-shade-10':
                props.actionItem.type == 'tax',
              'bg-clr-bad': props.actionItem.type == 'legal',
            }"
          >
            <p>
              {{
                props.actionItem.details.length >
                maxReccomendationLength
                  ? props.actionItem.details.slice(
                      0,
                      maxReccomendationLength
                    ) + "..."
                  : props.actionItem.details
              }}
            </p>
          </div>
        </div>
      </div>
      <!-- Chat Messages -->
      <div
        v-for="(message, index) in chatMessages"
        :key="index"
      >
        <AssistantMessage
          v-if="message.role === 'assistant' && index !== 0"
          :type="props.actionItem.type"
          :title="props.actionItem.title"
          :details="message.content"
        />
        <UserMessage
          v-if="message.role === 'user' && index !== 0"
          :content="message.content"
        />
      </div>
    </div>
    <!-- Chat Input -->
    <div class="flex items-center justify-center w-[100vw]">
      <!-- Input -->
      <form
        class="p-0 m-0 border-0 w-[100%]"
        @submit="sendMessage"
      >
        <input
          v-model="chatInput"
          type="text"
          class="flex-1 rounded-xl p-[15px] bg-gray-200 text-[20px] fixed bottom-0 w-[100%]"
          placeholder="Type a message..."
        />
      </form>
    </div>
  </div>
</template>
<script setup lang="ts">
import LeftArrowIcon from "@/components/icons/LeftArrowIcon.vue";
import ChatIcon from "@/components/icons/ChatIcon.vue";
import GavelIcon from "@/components/icons/GavelIcon.vue";
import BankIcon from "@/components/icons/BankIcon.vue";
import ConeIcon from "@/components/icons/ConeIcon.vue";
import ArchiveIcon from "@/components/icons/ArchiveIcon.vue";
import UserMessage from "@/components/UserMessage.vue";
import AssistantMessage from "@/components/AssistantMessage.vue";
import { defineProps, ref } from "vue";
import { ChatMessage, ActionItemType } from "@/types";

const maxReccomendationLength = 70;
const props = defineProps({
  actionItem: {
    type: Object as () => ActionItemType,
    default: () => ({
      type: "default",
      title: "",
      details: "",
    }),
  },
});
const emits = defineEmits(["closeChat"]);
const chatMessages = ref<ChatMessage[]>([
  {
    role: "assistant",
    content:
      "Lets talk about your reccomendation: " +
      props.actionItem.details,
  },
]);

const chatInput = ref("");
async function sendMessage(event: Event) {
  event.preventDefault();
  console.log("send message");
  chatMessages.value.push({
    role: "user",
    content: chatInput.value,
  });
  chatInput.value = "";

  const response = await fetch(
    import.meta.env.VITE_API_URL + "/v1/chat",
    {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(chatMessages.value),
    }
  );
  const data = await response.json();
  console.log(data);
  chatMessages.value = data;
}
</script>
