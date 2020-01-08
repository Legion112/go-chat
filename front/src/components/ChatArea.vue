<template>
    <div align="center">
        <h1>Hi I am a chat area</h1>
        <v-form>
        <v-text-field
            v-model="message"
            label="Type your message here"
        >
            <div slot="append-outer">
                <v-btn
                    @click="sendMessage"
                    class="primary">
                    Sent message
                </v-btn>
            </div>
        </v-text-field>
        </v-form>
        <v-list>
            <v-list-item
                    v-for="message in messages"
                    :key="message.id"
            >
                <v-list-item-content>
                    <v-list-item-title v-text="message.username"></v-list-item-title>
                </v-list-item-content>
                <v-list-item-content>
                    <v-list-item-title v-text="message.text"></v-list-item-title>
                </v-list-item-content>
            </v-list-item>
        </v-list>
    </div>
</template>

<script>
    import SocketService from "../services/SocketService";
    export default {
        name: "ChatArea",
        data(){
            return {
                message: null,
            }
        },
        methods: {
            sendMessage(){
                SocketService.sendDate({
                    text: this.message,
                    username: this.$store.state.username
                });
                this.message = null;
            }
        },
        computed:{
            messages(){
                return this.$store.state.messages
            }
        }
    }
</script>

<style scoped>

</style>