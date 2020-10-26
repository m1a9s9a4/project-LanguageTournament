/* eslint-disable */
<template>
  <v-main>
    <h2>home</h2>
    <v-row>
    <v-col md="3" cols="6" v-for="n in 10" v-bind:key="n">
        <Language 
            name="php"
        />
    </v-col>
    </v-row>
  </v-main>
</template>

<script>
import Language from './components/Home/Language';

export default {
    components: {
        Language,
    },
}
</script>