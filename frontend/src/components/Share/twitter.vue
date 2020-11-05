<template>
    <v-btn block @click="twitterShare" color="info"><v-icon>mdi-twitter</v-icon> {{ text }}</v-btn>
</template>
<script>
export default {
    data: () => ({
        baseUrl: 'https://twitter.com/intent/tweet',
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
        twitterShare() {
            //シェアする画面を設定
            const shareURL = this.baseUrl+'?text='+this.title+"%0a最強言語決定戦！！%0a%20%23languagetournament"+'&url='+this.url;  
           //シェア用の画面へ移行
            location.href = shareURL 
        },
    }
}
</script>