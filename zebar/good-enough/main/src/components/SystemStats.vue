<script lang="ts" setup>
import Island from "./Island.vue";
import Separator from "./Separator.vue";

defineProps<{
  cpu?: { usage: number };
  memory?: { usage: number };
  battery?: { isCharging: boolean; chargePercent: number };
}>();

const pad = (val?: number) => Math.round(val ?? 0).toString().padStart(2, '0');

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
</script>

<template>
  <Island position="center" class="gap-2 px-2.5">
    <div class="flex gap-1.5 items-center tracking-tighter uppercase">
      <i class="mt-px nf nf-oct-cpu"></i>
      <span class="font-bold tabular-nums">{{ pad(cpu?.usage) }}</span>
    </div>

    <Separator />

    <div class="flex gap-1.5 items-center tracking-tighter uppercase">
      <i class="mt-px nf nf-md-chip"></i>
      <span class="font-bold tabular-nums">{{ pad(memory?.usage) }}</span>
    </div>

    <template v-if="battery">
      <Separator />
      <div class="flex gap-1.5 items-center tracking-tighter uppercase"
        :class="battery.isCharging ? 'text-ctp-green' : battery.chargePercent < 20 ? 'text-ctp-red' : ''">
        <i :class="['nf', battery.isCharging ? 'nf-md-battery_charging' : getBatteryIcon(battery.chargePercent)]"></i>
        <span class="font-bold tabular-nums">{{ pad(battery.chargePercent) }}</span>
      </div>
    </template>
  </Island>
</template>
