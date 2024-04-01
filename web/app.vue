<script setup lang="ts">
interface Design {
    id: string,
    path: string,
    extension: string
}

const {data} = await useFetch("http://localhost:8080/designs");
const designs = data.value as Design[];

const config = useRuntimeConfig();
</script>

<template>
    <div class="designs">
        <img
            v-for="(design, i) in designs"
            :src="config.apiUrl + 'design/' + design.id"
            :alt="i"
            draggable="false"
            :key="i"
        >
    </div>
</template>

<style lang="scss">
body {
    color: #fff;

    background-color: #000;
}

* {
    margin: 0;
    padding: 0;
}

.designs {
    max-width: 1128px;

    display: flex;
    flex-wrap: wrap;
    gap: 24px;

    margin: 0 auto;
    padding: 24px;

    img {
        max-width: 264px;

        aspect-ratio: 3 / 4;
        object-fit: cover;

        background-color: #0a0b0a;
    }
}
</style>
