<script lang="ts" setup>
import { onMounted, ref, shallowRef } from "vue";
import * as zebar from "zebar";
import Clock from "./components/Clock.vue";
import Island from "./components/Island.vue";
import MediaPlayer from "./components/MediaPlayer.vue";
import SystemStats from "./components/SystemStats.vue";
import Workspaces from "./components/Workspaces.vue";

const providers = zebar.createProviderGroup({
  network: { type: "network" },
  glazewm: { type: "glazewm" },
  date: { type: "date", formatting: "HH:mm" },
  battery: { type: "battery" },
  media: { type: "media" },
  cpu: { type: "cpu" },
  memory: { type: "memory" },
});

const output = shallowRef(providers.outputMap);
const clockState = ref(0);
const isMediaHidden = ref(false);

onMounted(() => {
  providers.onOutput(() => {
    output.value = { ...providers.outputMap };
  });
});

const toggleClock = () => { clockState.value = (clockState.value + 1) % 3; };
const toggleMedia = () => { isMediaHidden.value = !isMediaHidden.value; };

// const mediaPlayerRef = ref<InstanceType<typeof MediaPlayer> | null>(null);
</script>

<template>
  <div class="bg-transparent flex items-center justify-between py-0.5 select-none text-[10px] text-ctp-subtext0 w-full">

    <div class="flex gap-2 items-center">
      <Island position="left" class="px-3">
        <i class="mt-px nf nf-fa-circle_notch text-[12px] text-ctp-blue"></i>
      </Island>

      <Workspaces :glazewm="output.glazewm ?? null" />
    </div>

    <div class="flex gap-2 h-6 items-center">
      <MediaPlayer ref="mediaPlayerRef" :media-output="output.media ?? undefined" @toggle="toggleMedia"
        v-show="!isMediaHidden" />

      <SystemStats :cpu="output.cpu ?? undefined" :memory="output.memory ?? undefined"
        :battery="output.battery ?? undefined" />

      <Clock :date="output.date ?? undefined" :state="clockState" @toggle="toggleClock" />
    </div>
  </div>
</template>
