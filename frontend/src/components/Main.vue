<script setup>
import { onMounted, ref } from 'vue';
import AppCard from './AppCard.vue';
import MessageCard from './MessageCard.vue';

const props = defineProps({
    user: {
        type: String,
        required: true
    }
})

const user = props.user;
const message = ref('');
const messages = ref([]);

const websocket = new WebSocket("ws://localhost:8080/ws");

const sendMessage = () => {
    websocket.send(JSON.stringify({
        user: user,
        message: message.value
    }));
    message.value = '';
}

websocket.onmessage = (event) => {
    const data = JSON.parse(event.data);
    messages.value.push(data);
}
</script>

<template>
    <AppCard title="Chatapp">
        <template #body>
            <v-container class="msg-container">
                <v-row v-for="m in messages">
                    <MessageCard :user="m.user" :message="m.message"></MessageCard>
                </v-row>
            </v-container>
        </template>

        <template #footer>
            <v-container class="input-container">
                <v-row>
                    <v-text-field class="message-input" v-model="message" label="Message" variant="solo"></v-text-field>
                    <v-btn class="input-btn" color="teal-darken-4" @click="sendMessage">Send</v-btn>
                </v-row>
            </v-container>
        </template>
    </AppCard>
</template>

<style scoped>
.input-container {
    height: 66px;
}
.input-btn {
    height: 50px;
}
.msg-container {
    height: calc(100vh - 300px);
    overflow-y: auto;
}
</style>