<template>
    <v-btn block @click="share" color="primary">Facebook</v-btn>
</template>
<script>
export default {
    data: () => ({
        baseUrl: 'https://www.facebook.com/sharer/sharer.php',
    }),
    props: {
        title: {
            Type: String,
            Default: '',
        },
        url: {
            Type: String,
            Default: '/',
        },
        text: {
            Type: String,
            Default: 'シェアする',
        },
    },
    methods: {
        share: function() {
            location.href = this.baseUrl+'u='+this.url;
        }
    }
}
</script>