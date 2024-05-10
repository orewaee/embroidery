// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    devtools: {enabled: true},
    modules: [
        [
            "@nuxtjs/google-fonts",
            {
                families: {
                    Inter: [400, 500, 600]
                }
            }
        ]
    ],
    runtimeConfig: {
        public: {
            apiUrl: process.env.NUXT_PUBLIC_API_URL
        }
    },
});
