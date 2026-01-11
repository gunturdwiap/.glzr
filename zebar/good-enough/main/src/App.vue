<script lang="ts" setup>
import { ref, computed, onMounted, shallowRef } from "vue";
import * as zebar from "zebar";

const CLOCK_STATES = { HIDDEN: 0, SHORT: 1, FULL: 2 };

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
const clockState = ref(CLOCK_STATES.HIDDEN);
const isMediaHidden = ref(false);

const formattedFullDate = computed(() => {
  const d = output.value.date?.new;
  if (!d) return "";
  const day = String(d.getDate()).padStart(2, '0');
  const month = String(d.getMonth() + 1).padStart(2, '0');
  return `${day}/${month} — ${output.value.date?.formatted}`;
});

const activeMedia = computed(() => output.value.media?.currentSession);

const getBatteryIcon = (percent: number) => {
  if (percent >= 95) return "nf-md-battery";
  if (percent >= 85) return "nf-md-battery_90";
  if (percent >= 75) return "nf-md-battery_80";
  if (percent >= 65) return "nf-md-battery_70";
  if (percent >= 55) return "nf-md-battery_60";
  if (percent >= 45) return "nf-md-battery_50";
  if (percent >= 35) return "nf-md-battery_40";
  if (percent >= 25) return "nf-md-battery_30";
  if (percent >= 15) return "nf-md-battery_20";
  if (percent >= 5) return "nf-md-battery_10";
  return "nf-md-battery_outline";
};

const pad = (val?: number) => Math.round(val ?? 0).toString().padStart(2, '0');

const truncate = (text: string | null, max: number) =>
  text && text.length > max ? text.substring(0, max) + "..." : text;

onMounted(() => {
  providers.onOutput(() => {
    output.value = { ...providers.outputMap };
  });
});

const toggleClock = () => { clockState.value = (clockState.value + 1) % 3; };
const toggleMedia = () => { isMediaHidden.value = !isMediaHidden.value; };
</script>

<template>
  <div
    class="bg-transparent flex font-mono items-center justify-between py-0.5 select-none text-[10px] text-ctp-subtext0 w-full">

    <div class="flex gap-2 items-center">
      <div class="bg-ctp-mantle border border-ctp-surface0 border-l-0 flex h-6 items-center px-3 rounded-r-sm">
        <i class="mt-px nf nf-fa-circle_notch text-[12px] text-ctp-blue"></i>
      </div>

      <div v-if="output.glazewm" class="bg-ctp-mantle border border-ctp-surface0 flex h-6 items-center px-1 rounded-sm">
        <template v-for="(ws, index) in output.glazewm.currentWorkspaces" :key="ws.name">
          <button @click="output.glazewm.runCommand(`focus --workspace ${ws.name}`)" :disabled="ws.hasFocus"
            class="border-b flex h-full items-center"
            :class="ws.hasFocus ? 'border-ctp-blue text-ctp-blue' : 'cursor-pointer border-transparent hover:border-ctp-subtext0 hover:text-ctp-text'">
            <span class="font-bold px-2 tracking-tighter uppercase">{{ ws.displayName ?? ws.name }}</span>
          </button>
          <span v-if="index < output.glazewm.currentWorkspaces.length - 1"
            class="border-l border-white h-2 mx-0.5 opacity-25" />
        </template>
      </div>
    </div>

    <div class="flex gap-2 h-6 items-center">

      <div v-if="activeMedia?.isPlaying" @click="toggleMedia"
        class="bg-ctp-mantle border border-ctp-surface0 cursor-pointer flex group h-full items-center px-3 rounded-sm hover:text-ctp-text">
        <div class="flex gap-2 items-center whitespace-nowrap"
          :class="isMediaHidden ? 'opacity-40 group-hover:opacity-100' : 'opacity-100'">
          <i class="mt-px nf nf-md-music_note text-ctp-lavender"></i>
          <span v-if="!isMediaHidden" class="italic leading-none">
            {{ truncate(activeMedia.title, 40) }}
            <span class="mx-1 opacity-40">•</span>
            <span class="not-italic opacity-60">{{ truncate(activeMedia.artist, 30) }}</span>
          </span>
          <span v-else class="font-bold opacity-50 px-1 tracking-widest">•••</span>
        </div>
      </div>

      <div
        class="bg-ctp-mantle border border-ctp-surface0 flex gap-3 h-full items-center px-3 rounded-sm cursor-default">
        <div class="flex gap-1.5 items-center tracking-tighter uppercase">
          <i class="mt-px nf nf-oct-cpu"></i>
          <span class="font-bold tabular-nums">{{ pad(output.cpu?.usage) }}</span>
        </div>

        <span class="border-l border-white h-2 mx-0.5 opacity-25" />

        <div class="flex gap-1.5 items-center tracking-tighter uppercase">
          <i class="mt-px nf nf-md-chip"></i>
          <span class="font-bold tabular-nums">{{ pad(output.memory?.usage) }}</span>
        </div>

        <template v-if="output.battery">
          <span class="border-l border-white h-2 mx-0.5 opacity-25" />
          <div class="flex gap-1.5 items-center tracking-tighter uppercase"
            :class="output.battery.isCharging ? 'text-ctp-green' : output.battery.chargePercent < 20 ? 'text-ctp-red' : ''">
            <i
              :class="['nf', output.battery.isCharging ? 'nf-md-battery_charging' : getBatteryIcon(output.battery.chargePercent)]"></i>
            <span class="font-bold tabular-nums">{{ pad(output.battery.chargePercent) }}</span>
          </div>
        </template>
      </div>

      <div @click="toggleClock"
        class="bg-ctp-mantle border border-r-0 border-ctp-surface0 cursor-pointer flex group h-full items-center px-4 rounded-l-sm hover:text-ctp-text">
        <div class="flex items-center whitespace-nowrap"
          :class="clockState === CLOCK_STATES.HIDDEN ? 'opacity-40 group-hover:opacity-100' : 'opacity-100'">
          <i class="mt-px nf nf-md-clock"></i>
          <span v-if="clockState !== CLOCK_STATES.HIDDEN" class="font-bold leading-none ml-2 tracking-tight">
            {{ clockState === CLOCK_STATES.SHORT ? output.date?.formatted : formattedFullDate }}
          </span>
        </div>
      </div>

    </div>
  </div>
</template>