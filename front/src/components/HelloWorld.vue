<template>
  <v-container>
    <v-layout
      text-center
      wrap
    >
      <v-flex xs12>
        <v-img
          :src="require('../assets/logo.svg')"
          class="my-3"
          contain
          height="200"
        ></v-img>
      </v-flex>

      <v-flex xs12>
        <h1 class="display-2 font-weight-bold mb-3">
          Welcome to our beautiful "Go chat"
        </h1>
      </v-flex>
      <v-flex xs2>
        <v-row align="center">
          <v-form>
            <v-text-field
                    ref="form"
                    v-model="username"
                    label="Name"
                    required
            ></v-text-field>
            <v-btn
                    color="success"
                    class="mr-4"
                    @click="setUsername"
            >
              Set username
            </v-btn>
          </v-form>
        </v-row>
      </v-flex>
      <v-flex xs6>
        <v-card width="600">
          <ChatArea></ChatArea>
        </v-card>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import { mapActions } from 'vuex'
import ChatArea from "./ChatArea";

export default {
  name: 'HelloWorld',
  components: {
    ChatArea
  },
  data: () => ({
    username: null,
    ecosystem: [
      {
        text: 'vuetify-loader',
        href: 'https://github.com/vuetifyjs/vuetify-loader',
      },
      {
        text: 'github',
        href: 'https://github.com/vuetifyjs/vuetify',
      },
      {
        text: 'awesome-vuetify',
        href: 'https://github.com/vuetifyjs/awesome-vuetify',
      },
    ],
    importantLinks: [
      {
        text: 'Documentation',
        href: 'https://vuetifyjs.com',
      },
      {
        text: 'Chat',
        href: 'https://community.vuetifyjs.com',
      },
      {
        text: 'Made with Vuetify',
        href: 'https://madewithvuejs.com/vuetify',
      },
      {
        text: 'Twitter',
        href: 'https://twitter.com/vuetifyjs',
      },
      {
        text: 'Articles',
        href: 'https://medium.com/vuetify',
      },
    ],
    whatsNext: [
      {
        text: 'Explore components',
        href: 'https://vuetifyjs.com/components/api-explorer',
      },
      {
        text: 'Select a layout',
        href: 'https://vuetifyjs.com/layout/pre-defined',
      },
      {
        text: 'Frequently Asked Questions',
        href: 'https://vuetifyjs.com/getting-started/frequently-asked-questions',
      },
    ],
  }),
  methods: {
    ...mapActions({
      storeUsername: 'setUsername'
    }),
    setUsername(){

      this.storeUsername(this.username);
    }
  }
};
</script>
