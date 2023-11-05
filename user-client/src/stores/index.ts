// stores/counter.js
import { defineStore } from "pinia";

export const useCounterStore = defineStore("counter", {
  state: () => {
    return { count: 0 };
    //For persistant storage
    // return {
    //     count: useLocalStorage('count', 0), //useLocalStorage takes in a key of 'count' and default value of 0
    //    };
  },

  actions: {
    increment() {
      this.count++;
    },
  },
});
