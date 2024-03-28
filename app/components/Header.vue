<template>
  <header
    class="fixed top-0 py-5 xl:px-20 px-10 transition-all w-full z-50"
    :class="
      isScrolled ? 'border-gray-100/20 border-b backdrop-blur-sm' : 'pt-[80px]'
    "
  >
    <UHorizontalNavigation class="z-60" :links="links" />
  </header>
</template>

<script lang="ts" setup>
import { useWindowScroll } from "@vueuse/core";

const { y } = useWindowScroll();
const isScrolled = computed(() => y.value > 10);

const links = [
  [
    {
      label: "PageWise",
      to: "/",
      labelClass: "text-3xl text-black",
    },
  ],
  [
    {
      label: "Search",
      to: "/search",
      icon: "i-heroicons-magnifying-glass",
      labelClass: "text-black text-xl",
      iconClass: "text-primary text-xl",
    },
  ],
];
</script>

<style></style>
