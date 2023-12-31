<template>
  <div class="flex flex-col">
    <div
      class="fixed top-[50px] left-0 h-[390px] w-[100%] bg-app-bg-clr flex flex-col justify-center items-center"
    >
      <!-- SCORE -->
      <div
        class="flex flex-col justify-center items-center -mb-[180px]"
      >
        <div
          v-if="scoreLoading || activeConfig.riskValue == -1"
          class="m-auto h-[140px]"
        >
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
              activeConfig.riskValue
                ? (100 - activeConfig.riskValue).toFixed(0)
                : ""
            }}
          </div>
          <div class="text-[20px] drop-shadow-md text-center">
            Score
          </div>
        </div>
      </div>
      <!-- SCORE CHART -->
      <canvas ref="scoreChart" id="score-chart"></canvas>

      <!-- Legend -->
      <div class="flex fixed top-[305px]">
        <div class="flex mx-[5px]">
          <div
            class="w-[20px] h-[20px] mr-[3px] bg-clr-good rounded-[4px]"
          ></div>
          <div>Bankruptcy</div>
        </div>
        <div class="flex mx-[5px]">
          <div
            class="w-[20px] h-[20px] mr-[3px] bg-clr-mid rounded-[4px]"
          ></div>
          <div>Baseline</div>
        </div>
        <div class="flex mx-[5px]">
          <div
            class="w-[20px] h-[20px] mr-[3px] bg-clr-bad rounded-[4px]"
          ></div>
          <div>Recs</div>
        </div>
      </div>
    </div>
    <!-- ACTIONS -->
    <div
      class="px-[10px] h-fit fixed top-[340px] left-0 h-[589px] bg-app-bg-clr"
    >
      <div
        class="text-[45px] drop-shadow-md ml-[20px] -mb-[8px]"
      >
        Actions
      </div>
      <div
        class="text-[18px] drop-shadow-md ml-[21px] mb-[5px]"
      >
        Do these to increase your score
      </div>
      <div
        class="flex flex-col h-[550px] overflow-y-auto pb-[50px]"
      >
        <div
          v-if="
            reccomendationsLoading ||
            activeReccomendations.length === 0 ||
            activeConfig.riskValue == -1
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
import { ref, watch, onMounted } from "vue";
import Chart from "chart.js/auto";
import ActionItem from "@/components/ActionItem.vue";
import ChatView from "@/views/ChatView.vue";
import DetailsModal from "@/components/DetailsModal.vue";
import LoadingIcon from "@/components/icons/LoadingIcon.vue";
import {
  ActionItemType,
  CompanyConfig,
  Reccomendation,
} from "@/types";

const validRecTypes = [
  "default",
  "tax",
  "legal",
  "finance",
  "saftey",
];
const activeConfig = ref<CompanyConfig>({} as CompanyConfig);
const activeReccomendations = ref<ActionItemType[]>(
  [] as ActionItemType[]
);
const weightedBankruptcy = ref(0);
const weightedBaseline = ref(0);
const weightedRecs = ref(0);
const scoreLoading = ref(false);
const reccomendationsLoading = ref(false);
async function pollActiveConfig() {
  console.log("Polling active config...");
  const response = await fetch(
    import.meta.env.VITE_API_URL + "/v1/active-config"
  );
  const data = await response.json();
  //If polling returns a new company name, the config has changed
  if (data.riskValue !== activeConfig.value.riskValue) {
    console.log("New active config", data);
    scoreLoading.value = true;
    activeConfig.value = data;
    weightedBankruptcy.value =
      activeConfig.value.weightedBankruptcy || 0;
    weightedBaseline.value =
      activeConfig.value.weightedBaseline || 0;
    weightedRecs.value = activeConfig.value.weightedRecs || 0;

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
            if (
              !validRecTypes.includes(
                reccomendation.type.toLocaleLowerCase()
              )
            )
              return;
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
    activeReccomendations.value.filter((reccomendation) => {
      if (reccomendation === undefined) return;
      reccomendation.details !== action.details;
    });
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
const data = {
  labels: ["Red", "Yellow", "Blue"],
  datasets: [
    {
      label: "My First Dataset",
      data: [
        weightedBankruptcy.value,
        weightedBaseline.value,
        weightedRecs.value,
      ],
      backgroundColor: ["#f57fce", "#f5ed7f", "#7fcef5"],
      borderColor: "#faf9f6",
      borderRadius: 10,
      borderWidth: 3,
      radius: 150,
      cutout: 140,
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
watch(weightedBankruptcy, () => {
  console.log("Updating chart");
  chart.data.datasets[0].data = [
    weightedBankruptcy.value,
    weightedBaseline.value,
    weightedRecs.value,
  ];
  chart.update();
});
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
