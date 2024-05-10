<script setup lang="ts">
const config = useRuntimeConfig();

interface Design {
    id: string,
    name: string,
    tag: string
}

const {data} = await useFetch(config.public.apiUrl + "designs");
const designs = data.value as Design[];
</script>

<template>
    <div class="designs">
        <div class="design" v-for="(design, i) in designs" :key="i">
            <img
                :src="config.public.apiUrl + 'design/' + design.id"
                :alt="i.toString()"
                draggable="false"
            >

            <div class="info">
                <p class="name">{{design.name}}</p>
                <p class="tag">{{design.tag}}</p>
            </div>
        </div>
    </div>
</template>

<style lang="scss">
body {
    color: #fff;

    background-color: #000;

    font-family: Inter, sans-serif;
}

* {
    margin: 0;
    padding: 0;
}

.designs {
    max-width: 1128px;

    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(264px, 1fr));
    grid-gap: 24px;

    margin: 0 auto;
    padding: 24px;

    .design {
        display: flex;
        flex-direction: column;
        gap: 12px;

        img {
            width: 100%;

            aspect-ratio: 3 / 4;
            object-fit: cover;

            background-color: #0a0b0a;
        }

        .info {
            display: flex;
            justify-content: space-between;

            font-weight: 600;

            .tag {
                color: #000;

                background-color: #fff;

                border-radius: 12px;

                user-select: none;

                padding: 0 8px;
            }
        }
    }
}
</style>
