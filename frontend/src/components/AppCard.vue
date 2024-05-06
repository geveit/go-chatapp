<script setup>
const props = defineProps({
    title: {
        type: String,
        required: true
    }
})
</script>

<template>
<v-app>
    <v-main>
        <v-container class="fill-height">
            <v-row class="fill-height" align="center" justify="center">
                <v-card class="app-card">
                    <v-toolbar color="teal-darken-4">
                        <v-toolbar-title>{{ props.title }}</v-toolbar-title>
                    </v-toolbar>
                    
                    <v-card-text>
                        <slot name="body"></slot>
                    </v-card-text>

                    <v-card-actions>
                        <slot name="footer"></slot>
                    </v-card-actions>
                </v-card>
            </v-row>
        </v-container>
    </v-main>
</v-app>
</template>

<style scoped>  
.app-card {
    width: 400px;
}
</style>