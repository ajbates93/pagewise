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
            @input="handleSearchInput"
          />
        </div>
      </div>
      <div></div>
    </LayoutDeviceHeightContainer>
    <div id="results" v-for="(book, idx) in searchResults">
      {{ book.title }}
    </div>
  </div>
</template>

<script lang="ts" setup>
import type { BaseAPIResponse } from "~/types/API";

type BookAPIResponse = {
  title: string;
};

const searchInput = ref<string>("");
const searchResults = ref<BookAPIResponse[]>();

const handleSearchInput = async () => {
  if (searchInput.value.length > 3) {
    // Call search API
    const { data, pending } = await useAPIFetch<BaseAPIResponse>(
      `/search?q=${searchInput.value}`
    );

    if (data.value && data.value.success) {
      searchResults.value = data.value.data as BookAPIResponse[];
    }
  }
};
</script>

<style></style>
