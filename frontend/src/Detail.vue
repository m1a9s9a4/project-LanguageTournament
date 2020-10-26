<template>
    <v-main>
        <h2>{{ language }} is Compared with...</h2>
        <v-row>
            <v-col cols="12" md="3">
                <Language />
            </v-col>
        </v-row>
    </v-main>
</template>

<script>
import Language from './components/Detail/Language';

export default {
    computed: {
        language: function () {
        return this.$route.params["lang"];
        },
    },
    components: {
        Language,
    },
}
</script>