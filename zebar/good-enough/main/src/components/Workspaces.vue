<script lang="ts" setup>
import { type GlazeWmOutput } from "zebar";
import Island from "./Island.vue";
import Separator from "./Separator.vue";

defineProps<{
  glazewm: GlazeWmOutput | null;
}>();
</script>

<template>
  <Island v-if="glazewm" position="center" class="px-0.5">
    <template v-for="(ws, index) in glazewm.currentWorkspaces" :key="ws.name">
      <button @click="glazewm.runCommand(`focus --workspace ${ws.name}`)" :disabled="ws.hasFocus"
        class="border-b flex h-full items-center px-2"
        :class="ws.hasFocus ? 'border-ctp-blue text-ctp-blue' : 'cursor-pointer border-transparent hover:border-ctp-subtext0 hover:text-ctp-text'">
        <span class="font-bold tracking-tighter uppercase">{{ ws.displayName ?? ws.name }}</span>
      </button>
      <Separator v-if="index < glazewm.currentWorkspaces.length - 1" />
    </template>
  </Island>
</template>
