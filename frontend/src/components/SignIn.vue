<script setup>
import { ref } from 'vue';
import AppCard from './AppCard.vue';

const props = defineProps({
    setUser: {
        type: Function,
        required: true
    }
})

const name = ref('');
const valid = ref(false);
const form = ref(null);

const nameRules = [
    (v) => !!v || 'Name is required',
    (v) => (v && v.length >= 3) || 'Name must be at least 3 characters',
];

const submit = () => {
    form.value.validate();
    if (valid.value) {
        props.setUser(name.value);
    }
}
</script>

<template>
<AppCard title="Sign In">
    <template #body>
        <v-form ref="form" @submit.prevent="submit" v-model="valid">
            <v-text-field
                v-model="name"
                :rules="nameRules"
                name="name"
                label="Name"
                type="text"
            ></v-text-field>
        </v-form>
    </template>
    <template #footer>
        <v-spacer></v-spacer>
        <v-btn color="teal-darken-4" @click="submit">SignIn</v-btn>
    </template>
</AppCard>
</template>

<style scoped>
</style>