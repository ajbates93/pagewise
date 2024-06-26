<template>
  <div id="hero" class="flex justify-between items-center">
    <LayoutDeviceHeightContainer>
      <span class="text-3xl">&nbsp;</span>
      <div class="content space-y-6">
        <h1 class="text-5xl">
          <!-- Intuitive &amp; Elegant.  -->
          Designed with a love of books.
        </h1>
        <p class="text-xl"></p>
        <UForm
          :schema="schema"
          :state="state"
          class="space-y-4 max-w-xs"
          @submit="onSubmit"
        >
          <div class="flex space-x-2 items-start">
            <UFormGroup name="email">
              <UInput
                size="xl"
                class="w-full min-w-[300px]"
                placeholder="e.g. jrrtolkein@gmail.com"
                v-model="state.email"
              />
            </UFormGroup>
            <UButton
              class="text-xl"
              :class="success ? 'bg-success' : 'bg-primary'"
              type="submit"
            >
              <span v-if="!success">Join Waitlist</span>
              <span v-else>Joined!</span>
            </UButton>
          </div>
        </UForm>
      </div>
      <div class="text-gray-400">
        By signing up you agree to our
        <a class="underline" href="/">Terms of Service</a> and
        <a class="underline" href="/">Privacy Policy</a>.
      </div>
    </LayoutDeviceHeightContainer>
    <picture>
      <img src="../../public/book1.svg" class="w-[800px]" />
    </picture>
  </div>
</template>

<script setup lang="ts">
import type { FormSubmitEvent } from "#ui/types";
import { z } from "zod";
import { useAPIFetch } from "~/composables/useAPIFetch";
import type { BaseAPIResponse } from "~/types/API.ts";

const schema = z.object({
  email: z.string().email("Invalid email"),
});

type Schema = z.output<typeof schema>;

const state = reactive({
  email: "",
});
const success = ref(false);
const errorMessage = ref<string>("");

const onSubmit = async (event: FormSubmitEvent<Schema>) => {
  const { data, pending } = await useAPIFetch<BaseAPIResponse>("/waitlist", {
    method: "POST",
    body: JSON.stringify(event.data.email),
  });
  if (data.value?.success) {
    success.value = true;
  } else {
    success.value = false;
    errorMessage.value = data.value?.message || "";
  }
};
</script>

<style></style>
