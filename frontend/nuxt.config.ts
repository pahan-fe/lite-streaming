import tailwindcss from '@tailwindcss/vite'

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',
  devtools: { enabled: true },

  modules: ['@nuxt/eslint', '@vueuse/nuxt', '@nuxt/fonts'],

  css: ['~/assets/css/main.css', 'vue-sonner/style.css'],

  fonts: {
    experimental: {
      processCSSVariables: true,
    },
    families: [
      { name: 'Fraunces', provider: 'google', weights: [400, 500, 600], styles: ['normal'] },
      {
        name: 'Hanken Grotesk',
        provider: 'google',
        weights: [400, 500, 600, 700],
        styles: ['normal'],
      },
    ],
  },

  vite: {
    plugins: [tailwindcss()],
    server: {
      proxy: {
        '/api': process.env.NUXT_API_BASE || 'http://localhost:8080',
      },
    },
    optimizeDeps: {
      include: ['class-variance-authority', 'clsx', 'reka-ui', 'tailwind-merge', 'zod'],
    },
  },

  runtimeConfig: {
    apiBase: 'http://localhost:8080',
    public: {
      apiBase: '',
    },
  },

  typescript: {
    strict: true,
    typeCheck: true,
  },

  eslint: {
    config: {
      stylistic: false,
    },
  },

  components: [
    { path: '~/features', pattern: '**/*.vue', pathPrefix: false },
    { path: '~/shared/components', pattern: '**/*.vue', pathPrefix: false },
  ],
  imports: {
    dirs: ['features/**/composables', 'features/**/api'],
  },
})
