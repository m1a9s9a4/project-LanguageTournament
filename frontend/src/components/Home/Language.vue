<template>
    <v-card class="mx-auto" max-width="300" elevation="10">
        <a :href="'/detail/'+english">
          <v-img contain :alt="name" height="200" :src="imgsrc"></v-img>
        </a>
    </v-card>
</template>

<script>
export default {
  props: {
    english: {
      type: String,
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
};
</script>