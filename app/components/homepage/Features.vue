<template>
  <div id="features" class="">
    <LayoutDeviceHeightContainer>
      <div class="flex flex-col">
        <div class="text-center mb-20">
          <h2 class="text-4xl font-bold">Features</h2>
        </div>
        <div id="featuresContainer" class="space-y-12 relative">
          <div
            v-for="(feature, idx) in features"
            class="bg-white rounded border-primary-500 p-10 flex w-2/3"
            :class="`translate-x-${idx}/4`"
          >
            <div class="mr-5 w-full h-full max-w-40 block">
              <img :src="`/${feature.icon}`" />
            </div>
            <div>
              <div class="text-2xl text-primary-500">{{ feature.title }}</div>
              <div>{{ feature.description }}</div>
            </div>
          </div>
        </div>
      </div>
    </LayoutDeviceHeightContainer>
  </div>
</template>

<script lang="ts" setup>
type Feature = {
  title: string;
  description: string;
  icon: string;
};
const features = ref<Feature[]>([
  {
    title: "Recommendation System",
    description:
      "Either you like a book or you don't. We use a recommendation system that learns from your reading habits and suggests books that you are likely to enjoy.",
    icon: "recommendation.svg",
  },
  {
    title: "Library History",
    description:
      "Keep track of everything you've read, are reading, or want to read. We keep a history of all the books you've interacted with.",
    icon: "library.svg",
  },
  {
    title: "Meet People Like You",
    description:
      "PageWise is a community of book lovers. Meet people who share your choice in books and discuss your favorites together.",
    icon: "people.svg",
  },
]);
</script>

<style></style>
