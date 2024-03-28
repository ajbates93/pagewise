<template>
  <div id="search" class="flex justify-between items-center">
    <LayoutDeviceHeightContainer class="w-full">
      <div></div>
      <div class="flex items-center gap-10 w-full">
        <div class="space-y-2">
          <h1 class="text-5xl">Search all books</h1>
          <p class="text-xl text-gray-400">
            Search for a title, ISBN, author or genre
          </p>
        </div>
        <div>
          <UInput
            size="xl"
            class="w-full min-w-[500px]"
            placeholder="e.g. The Lord of the Rings"
            v-model="searchInput"
          />
        </div>
      </div>
      <div></div>
    </LayoutDeviceHeightContainer>
  </div>
</template>

<script lang="ts" setup>
const searchInput = ref<string>("");
</script>

<style></style>
