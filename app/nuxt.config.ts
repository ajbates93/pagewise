// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  modules: ["@nuxt/ui", "@vueuse/nuxt"],
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

  tailwindcss: {
    config: {
      theme: {
        extend: {
          colors: {
            customPrimary: {
              50: "#fef2f3",
              100: "#fde5e7",
              200: "#fbc9ce",
              300: "#f9adb5",
              400: "#f6909d",
              500: "#f49097",
              600: "#dc828a",
              700: "#b66b70",
              800: "#905457",
              900: "#763f43",
            },
          },
        },
      },
    },
  },

  css: ["@/assets/css/main.css"],
});
