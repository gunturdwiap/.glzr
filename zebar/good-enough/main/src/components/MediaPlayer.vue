<script lang="ts" setup>
import { computed } from "vue";
import type { MediaOutput } from "zebar";
import { useMediaPlayer } from "../composables/useMediaPlayer";
import Island from "./Island.vue";

const props = defineProps<{
  mediaOutput: MediaOutput | undefined;
}>();

const emit = defineEmits<{
  toggle: [];
}>();

const activeMedia = computed(() => props.mediaOutput?.currentSession ?? null);
const { progress, isHidden, toggle, onAction, truncate } = useMediaPlayer(
  activeMedia,
  computed(() => props.mediaOutput)
);

defineExpose({ toggle });
</script>

<template>
  <Island v-if="activeMedia" @click="emit('toggle')"
    class="relative cursor-pointer group px-3 rounded-sm hover:text-ctp-text overflow-hidden">
    <div class="flex gap-3 items-center whitespace-nowrap"
      :class="isHidden ? 'opacity-40 group-hover:opacity-100' : 'opacity-100'">

      <i class="mt-px nf nf-md-music_note text-ctp-lavender"></i>
      <span class="border-l border-white h-2 mx-0.5 opacity-25" />

      <span v-if="!isHidden" class="italic leading-none">
        {{ truncate(activeMedia.title, 40) }}
        <span class="mx-1 opacity-40">•</span>
        <span class="not-italic opacity-60">{{ truncate(activeMedia.artist, 30) }}</span>
      </span>
      <span v-else class="font-bold opacity-50 px-1 tracking-widest">•••</span>

      <span class="border-l border-white h-2 mx-0.5 opacity-25" />

      <div class="flex gap-2 text-ctp-subtext0">
        <i @click.stop="onAction('previous')"
          class="nf nf-md-skip_previous text-[8px] hover:text-ctp-lavender transition-colors"></i>
        <i @click.stop="onAction('togglePlayPause')"
          :class="['nf text-[8px] hover:text-ctp-lavender transition-colors', activeMedia.isPlaying ? 'nf-md-pause' : 'nf-md-play']"></i>
        <i @click.stop="onAction('next')"
          class="nf nf-md-skip_next text-[8px] hover:text-ctp-lavender transition-colors"></i>
      </div>
    </div>

    <div class="absolute bottom-0 left-0 border-b border-ctp-blue transition-all duration-100" :style="{
      width: activeMedia?.endTime
        ? `${(progress / activeMedia.endTime) * 100}%`
        : '0'
    }"></div>
  </Island>
</template>
