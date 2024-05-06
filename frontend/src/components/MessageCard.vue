<script setup>
const props = defineProps({
    user: {
        type: String,
        required: true
    },
    message: {
        type: String,
        required: true
    }
})
</script>

<template>
<v-card class="msg-card">
    <v-card-text>
        <p>{{ props.user }}:</p>
        <p>{{ props.message }}</p>
    </v-card-text>
</v-card>
</template>

<style scoped>
.msg-card {
    width: 100%;
    margin-bottom: 7px;
}
</style>