<template>
  <div class="flex flex-col">
    <div
      class="fixed top-[50px] left-0 h-[390px] w-[100%] bg-app-bg-clr flex flex-col justify-center items-center"
    >
      <!-- SCORE -->
      <div
        class="flex flex-col justify-center items-center -mb-[180px]"
      >
        <div v-if="scoreLoading" class="m-auto h-[140px]">
          <LoadingIcon />
        </div>
        <div v-else>
          <p
            class="text-[25px] -mb-[30px] drop-shadow-md text-center"
          >
            {{ activeConfig.active?.companyName || "" }}
          </p>
          <div
            class="text-[110px] -mb-[35px] drop-shadow-md text-center"
          >
            {{
              100 -
                parseFloat(
                  activeConfig.zValue?.toPrecision(2)
                ) || ""
            }}
          </div>
          <div class="text-[20px] drop-shadow-md text-center">
            Score
          </div>
        </div>
      </div>
      <!-- SCORE CHART -->
      <SpeedoChart />
    </div>

    <!-- ACTIONS -->
    <div
      class="px-[10px] h-fit fixed top-[310px] left-0 h-[589px] bg-app-bg-clr"
    >
      <div
        class="text-[45px] drop-shadow-md ml-[20px] mb-[5px]"
      >
        Actions
      </div>
      <div class="flex flex-col h-[550px] overflow-y-auto">
        <div
          v-if="
            reccomendationsLoading ||
            activeReccomendations.length === 0
          "
          class="h-[124px] w-[370px] flex"
        >
          <LoadingIcon class="m-auto" />
        </div>
        <div
          v-else
          v-for="(action, index) in activeReccomendations"
        >
          <ActionItem
            v-if="action?.completed === false"
            :key="index"
            :action-item="action"
            @open-chat="toggleChat"
            @select-action="() => selectAction(action)"
          />
        </div>
      </div>
    </div>
  </div>

  <!-- CHAT -->
  <Transition name="chat-slide-left">
    <ChatView
      v-if="showChat"
      :action-item="selectedAction"
      @close-chat="toggleChat"
    />
  </Transition>
  <!-- Modal -->
  <Transition name="modal-slide-up">
    <DetailsModal
      v-if="showDetailsModal"
      :action-item="selectedAction"
      @openChat="toggleChat"
      @close-modal="closeModal"
      @complete="() => completeAction(selectedAction)"
    />
  </Transition>
</template>
<script setup lang="ts">
import { ref } from "vue";
import SpeedoChart from "@/components/SpeedoChart.vue";
import ActionItem from "@/components/ActionItem.vue";
import ChatView from "@/views/ChatView.vue";
import DetailsModal from "@/components/DetailsModal.vue";
import LoadingIcon from "@/components/icons/LoadingIcon.vue";
import {
  ActionItemType,
  CompanyConfig,
  Reccomendation,
} from "@/types";

const activeConfig = ref<CompanyConfig>({} as CompanyConfig);
const activeReccomendations = ref<ActionItemType[]>(
  [] as ActionItemType[]
);
const scoreLoading = ref(false);
const reccomendationsLoading = ref(false);
async function pollActiveConfig() {
  console.log("Polling active config...");
  const response = await fetch(
    import.meta.env.VITE_API_URL + "/v1/active-config"
  );
  const data = await response.json();
  //If polling returns a new company name, the config has changed
  if (
    data.active?.companyName !==
    activeConfig.value.active?.companyName
  ) {
    console.log("New active config", data);
    scoreLoading.value = true;
    reccomendationsLoading.value = true;
    const reccomendationsResponse = await fetch(
      import.meta.env.VITE_API_URL + "/v1/recommendations"
    );
    const reccomendationsData =
      await reccomendationsResponse.json();
    scoreLoading.value = false;
    setTimeout(() => {
      reccomendationsLoading.value = false;
      //Map reccomendations to action types
      const reccomendedActions: ActionItemType[] =
        reccomendationsData.map(
          (reccomendation: Reccomendation) => {
            if (reccomendation.completed) return;
            const action: ActionItemType = {
              type: reccomendation.type.toLocaleLowerCase(),
              title: reccomendation.title,
              details: reccomendation.details,
              completed: reccomendation.completed,
              id: (
                Math.floor(Math.random() * 1000000) + 1
              ).toString(),
            };
            return action;
          }
        );
      console.log("reccomended actions", reccomendedActions);
      activeReccomendations.value = reccomendedActions;
    }, 1000);
  }
  activeConfig.value = data;
  setTimeout(pollActiveConfig, 1000);
  return data;
}
pollActiveConfig();
const selectedAction = ref<ActionItemType>({
  type: "default",
  title: "",
  details: "",
  id: "",
  completed: false,
});
function selectAction(action: ActionItemType) {
  console.log("selected action", action);
  selectedAction.value = action;
  activateModal();
}
async function completeAction(action: ActionItemType) {
  console.log("complete action", action);
  //Remove reccomendation from the UI immediately
  activeReccomendations.value =
    activeReccomendations.value.filter(
      (reccomendation) => reccomendation.id !== action.id
    );
  //Send request to complete reccomendation
  const res = await fetch(
    import.meta.env.VITE_API_URL + "/v1/toggle-recommendation",
    {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        title: action.title,
      }),
    }
  );
  const data = await res.json();
  console.log(data);
}
const showChat = ref(false);
function toggleChat() {
  console.log("toggle chat");
  showChat.value = !showChat.value;
}
const showDetailsModal = ref(false);
function activateModal() {
  console.log("activate modal");
  showDetailsModal.value = true;
}
function closeModal() {
  console.log("close modal");
  showDetailsModal.value = false;
}
</script>

<style scoped>
.chat-slide-left-enter-active {
  animation: slide-left 0.3s;
}

.chat-slide-left-leave-active {
  animation: slide-left 0.3s reverse;
}

@keyframes slide-left {
  0% {
    transform: translateX(60%);
  }
  100% {
    transform: translateX(0px);
  }
}

.modal-slide-up-enter-active {
  animation: slide-up 0.3s;
}

.modal-slide-up-leave-active {
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