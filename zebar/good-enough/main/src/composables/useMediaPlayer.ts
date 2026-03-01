import { onUnmounted, ref, watch, type ComputedRef } from "vue";
import type { MediaOutput, MediaSession } from "zebar";

export function useMediaPlayer(
  activeMedia: ComputedRef<MediaSession | null>,
  media: ComputedRef<MediaOutput | null>
) {
  const progress = ref(0);
  const isHidden = ref(false);
  let timer: ReturnType<typeof setInterval> | null = null;
  let isResetting = false;
  let lastTitle = "";
  let lastSessionId = "";

  const toggle = () => { isHidden.value = !isHidden.value; };

  const onAction = (action: 'previous' | 'togglePlayPause' | 'next') => {
    media.value?.[action]?.();
  };

  const truncate = (text: string | null, max: number) =>
    text && text.length > max ? text.substring(0, max) + "..." : text;

  watch(
    () => [activeMedia.value?.title, activeMedia.value?.sessionId],
    ([newTitle, newSessionId]) => {
      if (!newTitle) return;
      if (newSessionId !== lastSessionId) {
        lastSessionId = newSessionId as string;
        lastTitle = newTitle as string;
        progress.value = activeMedia.value?.position ?? 0;
        isResetting = false;
        return;
      }
      if (newTitle !== lastTitle) {
        lastTitle = newTitle as string;
        progress.value = 0;
        isResetting = true;
        setTimeout(() => { isResetting = false; }, 3000);
      }
    }
  );

  watch(
    () => activeMedia.value?.position,
    (newPos) => {
      if (newPos === undefined) return;
      if (isResetting) {
        if (newPos < 5) {
          progress.value = newPos;
          isResetting = false;
        }
        return;
      }
      if (Math.abs(progress.value - newPos) > 2) {
        progress.value = newPos;
      }
    }
  );

  const startTimer = () => {
    if (timer) return;
    timer = setInterval(() => {
      const session = activeMedia.value;
      if (session?.isPlaying && session.endTime > 0 && progress.value < session.endTime) {
        progress.value++;
      }
    }, 1000);
  };

  const stopTimer = () => {
    if (timer) { clearInterval(timer); timer = null; }
  };

  watch(() => activeMedia.value?.isPlaying, (playing) => {
    playing ? startTimer() : stopTimer();
  }, { immediate: true });

  onUnmounted(stopTimer);

  return { progress, isHidden, toggle, onAction, truncate };
}
