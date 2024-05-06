<script setup>
import { ref } from 'vue';
import SignIn from './components/SignIn.vue';
import Main from './components/Main.vue';

const user = ref('');

const setUser = (u) => {
  user.value = u;
}

</script>

<template>
  <Main v-if="user" :user="user" />
  <SignIn v-else="user" :setUser="setUser" />
</template>

<style scoped>
</style>
