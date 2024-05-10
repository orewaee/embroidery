<script setup lang="ts">

</script>

<template>
    <header>
        <NuxtLink class="logo" to="/">
            <img src="public/logo.jpg" alt="" />
        </NuxtLink>
    </header>
</template>

<style scoped lang="scss">
header {
    position: sticky;
    top: 0;

    width: 100vw;

    display: flex;

    background-color: rgba(0, 0, 0, .6);
    backdrop-filter: blur(24px);
}

.logo {
    width: 48px;
    height: 48px;

    display: block;

    border-radius: 24px;

    overflow: hidden;

    margin: 16px auto;

    img {
        width: 48px;
        height: 48px;
    }
}
</style>