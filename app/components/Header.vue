<template>
  <header
    class="fixed top-0 py-5 xl:px-20 px-10 transition-all w-full z-50 backdrop-blur-sm"
    :class="isScrolled ? 'border-gray-100/20 border-b' : 'pt-[80px]'"
  >
    <!-- <div
      class="w-full h-full absolute left-0 top-0 transition-all z-10"
      :class="{ 'blur-md': isScrolled }"
    ></div> -->
    <div class="text-3xl z-60">PageWise</div>
  </header>
</template>

<script lang="ts" setup>
import { useWindowScroll } from "@vueuse/core";

const { y } = useWindowScroll();

const isScrolled = computed(() => y.value > 10);
</script>

<style></style>
