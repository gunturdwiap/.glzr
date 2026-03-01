<script lang="ts" setup>
import { computed } from "vue";

const props = withDefaults(defineProps<{
  position?: 'left' | 'right' | 'center';
}>(), {
  position: 'center'
});

const positionClasses = computed(() => {
  switch (props.position) {
    case 'left': return 'rounded-r-sm border-l-0';
    case 'right': return 'rounded-l-sm border-r-0';
    case 'center': return 'rounded-sm';
  }
});
</script>

<template>
  <div class="bg-ctp-mantle border border-ctp-surface0 flex h-6 items-center" :class="positionClasses">
    <slot />
  </div>
</template>
