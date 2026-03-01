<script lang="ts" setup>
import { onMounted, ref, shallowRef } from "vue";
import * as zebar from "zebar";
import Clock from "./components/Clock.vue";
import Island from "./components/Island.vue";
import MediaPlayer from "./components/MediaPlayer.vue";
import SystemStats from "./components/SystemStats.vue";
import Workspaces from "./components/Workspaces.vue";

const providers = zebar.createProviderGroup({
  glazewm: { type: "glazewm" },
  date: { type: "date", formatting: "HH:mm" },
  battery: { type: "battery" },
  media: { type: "media" },
  cpu: { type: "cpu" },
  memory: { type: "memory" },
});

const output = shallowRef(providers.outputMap);

const loadClockState = () => {
  try {
    const saved = localStorage.getItem('clockState');
    return saved ? parseInt(saved, 10) : 0;
  } catch { return 0; }
};

const clockState = ref(loadClockState());
const isMediaHidden = ref(false);

onMounted(() => {
  providers.onOutput(() => {
    output.value = { ...providers.outputMap };
  });
});

const toggleClock = () => {
  clockState.value = (clockState.value + 1) % 3;
  try {
    localStorage.setItem('clockState', clockState.value.toString());
  } catch { /* ignore */ }
};
const toggleMedia = () => { isMediaHidden.value = !isMediaHidden.value; };

</script>

<template>
  <div class="bg-transparent flex items-center justify-between py-0.5 select-none text-[10px] text-ctp-subtext0 w-full">

    <div class="flex gap-2 items-center">
      <Island position="left" class="px-3">
        <i class="mt-px nf nf-fa-circle_notch text-[12px] text-ctp-blue"></i>
      </Island>

      <Workspaces :glazewm="output.glazewm" />
    </div>

    <div class="flex gap-2 h-6 items-center">
      <MediaPlayer :media-output="output.media" @toggle="toggleMedia" v-show="!isMediaHidden" />

      <SystemStats :cpu="output.cpu" :memory="output.memory" :battery="output.battery" />

      <Clock :date="output.date" :state="clockState" @toggle="toggleClock" />
    </div>
  </div>
</template>
