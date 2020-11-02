<template>
    <v-card class="mx-auto" max-width="300" elevation="10" @click="selectedPlayer">
        <v-img contain :alt="name" height="200" :src="imgsrc"></v-img>
    </v-card>
</template>

<script>
export default {
  props: {
    id: {
      type: Number,
    },
    name: {
      type: String,
      default: "vuetify",
    },
    imgsrc: {
      type: String,
      default: "https://cdn.vuetifyjs.com/images/logos/vuetify-logo-dark.png",
    },
  },
  methods: {
    selectedPlayer() {
      this.$emit('selected-id', this.id);
    }
  }
};
</script>