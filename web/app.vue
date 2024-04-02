<script setup lang="ts">
interface Design {
    id: string,
    name: string,
    tag: string
}

const {data} = await useFetch("http://localhost:8080/designs");
const designs = data.value as Design[];

const config = useRuntimeConfig();
</script>

<template>
    <div class="designs">
        <div class="design" v-for="(design, i) in designs" :key="i">
            <img
                :src="config.apiUrl + 'design/' + design.id"
                :alt="i"
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

    display: flex;
    flex-wrap: wrap;
    gap: 24px;

    margin: 0 auto;
    padding: 24px;

    .design {
        max-width: 264px;

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

                padding: 0 8px;
            }
        }
    }
}
</style>
