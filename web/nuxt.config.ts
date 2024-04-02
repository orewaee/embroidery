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
        apiUrl: process.env.NUXT_API_URL
    }
});
