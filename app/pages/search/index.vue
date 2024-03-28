<template>
  <div id="search">
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
            :loading="isLoading"
            icon="i-heroicons-magnifying-glass-20-solid"
            class="w-full min-w-[500px]"
            placeholder="e.g. The Lord of the Rings"
            v-model="searchInput"
          />
        </div>
      </div>
      <div></div>
    </LayoutDeviceHeightContainer>
    <LayoutDeviceHeightContainer>
      <div class="p-20">
        <h2 class="text-3xl">Results</h2>
        <div id="results" v-for="(book, idx) in searchResults">
          <div>
            <h3>{{ book.volumeInfo.title }}</h3>
            <p class="text-gray-400">
              {{ book.volumeInfo.authors.join(", ") }}
            </p>
            <p>{{ book.volumeInfo.description }}</p>
          </div>
        </div>
      </div>
    </LayoutDeviceHeightContainer>
  </div>
</template>

<script lang="ts" setup>
import type { BaseAPIResponse, APIBookResponse } from "~/types/API";
import { refDebounced } from "@vueuse/core";

const isLoading = ref<boolean>(false);
const searchInput = ref<string>("");
const searchInputDebounced = refDebounced(searchInput, 500);
const searchResults = ref<APIBookResponse[]>();

const handleSearchInput = async () => {
  try {
    if (searchInput.value.length > 3) {
      // Call search API
      const { data, pending } = await useAPIFetch<BaseAPIResponse>(
        `/search?q=${searchInput.value}`
      );

      if (data.value && data.value.success) {
        searchResults.value = data.value.data as APIBookResponse[];
      }
    } else if (searchInput.value.length === 0) {
      searchResults.value = [];
    }
  } catch (e: any) {
    console.error(e);
  } finally {
    isLoading.value = false;
  }
};

watch(searchInputDebounced, handleSearchInput);
watch(searchInput, () => {
  if (searchInput.value !== searchInputDebounced.value) isLoading.value = true;
});
</script>

<style></style>
