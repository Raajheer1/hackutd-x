/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  theme: {
    extend: {
      colors: {
        "app-bg-clr": "#faf9f6",
        "clr-good": "#7fcef5",
        "clr-mid": "#f5ed7f",
        "clr-bad": "#f57fce",
        //Good shades
        "clr-good-shade-10": "#72b9dd",
        "clr-good-shade-20": "#66a5c4",
        "clr-good-shade-30": "#5990ac",
        "clr-good-shade-40": "#4c7c93",
        "clr-good-shade-50": "#40677b",
        //Good tints
        "clr-good-tint-10": "#8cd3f6",
        "clr-good-tint-20": "#99d8f7",
        "clr-good-tint-30": "#a5ddf8",
        "clr-good-tint-40": "#b2e2f9",
        "clr-good-tint-50": "#bfe7fa",
        //Mid shades
        "clr-mid-shade-5": "#f3e965",
        "clr-mid-shade-7": "#f1e64e",
        "clr-mid-shade-10": "#ddd572",
        "clr-mid-shade-20": "#c4be66",
        "clr-mid-shade-30": "#aca659",
        "clr-mid-shade-40": "#938e4c",
        "clr-mid-shade-50": "#7b7740",
        //Mid tints
        "clr-mid-tint-10": "#f6ef8c",
        "clr-mid-tint-20": "#f7f199",
        "clr-mid-tint-30": "#f8f2a5",
        "clr-mid-tint-40": "#f9f4b2",
        "clr-mid-tint-50": "#faf6bf",
        //Bad shades
        "clr-bad-shade-10": "#dd72b9",
        "clr-bad-shade-20": "#c466a5",
        "clr-bad-shade-30": "#ac5990",
        "clr-bad-shade-40": "#934c7c",
        "clr-bad-shade-50": "#7b4067",
        //Bad tints
        "clr-bad-tint-10": "#f68cd3",
        "clr-bad-tint-20": "#f799d8",
        "clr-bad-tint-30": "#f8a5dd",
        "clr-bad-tint-40": "#f9b2e2",
        "clr-bad-tint-50": "#fabfe7",
        //User chat colors
        "clr-user-chat": "#5c5c5c",
        //User chat shades
        "clr-user-chat-shade-10": "#535353",
        "clr-user-chat-shade-20": "#4a4a4a",
        "clr-user-chat-shade-30": "#404040",
        "clr-user-chat-shade-40": "#373737",
        "clr-user-chat-shade-50": "#2e2e2e",
        //User chat tints
        "clr-user-chat-tint-10": "#6c6c6c",
        "clr-user-chat-tint-20": "#7d7d7d",
        "clr-user-chat-tint-30": "#8d8d8d",
        "clr-user-chat-tint-40": "#9d9d9d",
        "clr-user-chat-tint-50": "#aeaeae",
        "clr-user-chat-tint-60": "#bebebe",
        "clr-user-chat-tint-70": "#cecece",
      },
    },
  },
  plugins: [],
};
