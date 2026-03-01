<script lang="ts" setup>
import { computed } from "vue";
import Island from "./Island.vue";

const CLOCK_STATES = { HIDDEN: 0, SHORT: 1, FULL: 2 } as const;

const props = defineProps<{
  date: { formatted: string; new: Date } | undefined;
  state: number;
}>();

const emit = defineEmits<{
  toggle: [];
}>();

const formattedFullDate = computed(() => {
  const d = props.date?.new;
  if (!d) return "";
  const day = String(d.getDate()).padStart(2, '0');
  const month = String(d.getMonth() + 1).padStart(2, '0');
  return `${day}/${month} — ${props.date?.formatted}`;
});
</script>

<template>
  <Island position="right" class="cursor-pointer group px-3 hover:text-ctp-text" @click="emit('toggle')">
    <div class="flex items-center whitespace-nowrap"
      :class="state === CLOCK_STATES.HIDDEN ? 'opacity-40 group-hover:opacity-100' : 'opacity-100'">
      <i class="nf nf-fa-clock font-bold text-[12px]"></i>
      <span v-if="state !== CLOCK_STATES.HIDDEN" class="font-bold leading-none ml-2 tracking-tight">
        {{ state === CLOCK_STATES.SHORT ? date?.formatted : formattedFullDate }}
      </span>
    </div>
  </Island>
</template>
