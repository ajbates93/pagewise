// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  modules: ["@nuxt/ui"],
  devtools: { enabled: true },

  app: {
    head: {
      link: [
        {
          rel: "preconnect",
          href: "https://cdn.fonts.net",
        },
        {
          rel: "stylesheet",
          href: "https://cdn.fonts.net/kit/b7c9ac61-d608-474b-a267-8e7869901d30/b7c9ac61-d608-474b-a267-8e7869901d30.css",
        },
      ],
    },
  },

  css: ["@/assets/css/main.css"],
});
