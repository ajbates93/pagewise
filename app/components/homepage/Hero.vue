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
            <UButton class="text-xl bg-primary" type="submit"
              >Join Waitlist</UButton
            >
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

const schema = z.object({
  email: z.string().email("Invalid email"),
});

type Schema = z.output<typeof schema>;

const state = reactive({
  email: undefined,
});

const onSubmit = (event: FormSubmitEvent<Schema>) => {
  console.log(`Email captured: ${event.data.email}`);
};
</script>

<style></style>
