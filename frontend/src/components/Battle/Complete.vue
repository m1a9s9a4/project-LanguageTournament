<template>
    <v-main class="text-center">
        <h3>全ての問題に回答しました！</h3>
        <p class="subtitle-1">ご協力ありがとうございました！</p>
        <v-row>
            <v-col cols="6">
                <v-img :src="p1.img"></v-img>
            </v-col>
            <v-col cols="6">
                <v-img :src="p2.img"></v-img>
            </v-col>
        </v-row>
        <v-row>
            <v-col>
                <v-btn href="/" block color="secondary">
                    トップに戻る
                </v-btn>
            </v-col>
        </v-row>
    </v-main>
</template>

<script>
export default {
    props: {
        p1: {
            type: Object,
        },
        p2: {
            type: Object,
        },
    }
}
</script>