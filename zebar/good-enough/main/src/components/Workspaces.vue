<script lang="ts" setup>
import Island from "./Island.vue";

defineProps<{
  glazewm: {
    currentWorkspaces: { name: string; displayName?: string; hasFocus: boolean }[];
    runCommand: (cmd: string) => void;
  } | null;
}>();
</script>

<template>
  <Island v-if="glazewm" class="px-1 rounded-sm">
    <template v-for="(ws, index) in glazewm.currentWorkspaces" :key="ws.name">
      <button @click="glazewm.runCommand(`focus --workspace ${ws.name}`)" :disabled="ws.hasFocus"
        class="border-b flex h-full items-center"
        :class="ws.hasFocus ? 'border-ctp-blue text-ctp-blue' : 'cursor-pointer border-transparent hover:border-ctp-subtext0 hover:text-ctp-text'">
        <span class="font-bold px-2 tracking-tighter uppercase">{{ ws.displayName ?? ws.name }}</span>
      </button>
      <span v-if="index < glazewm.currentWorkspaces.length - 1" class="border-l border-white h-2 mx-0.5 opacity-25" />
    </template>
  </Island>
</template>
