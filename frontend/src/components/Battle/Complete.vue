<template>
    <v-main class="text-center">
        <h3>全ての問題に回答しました！</h3>
        <p class="subtitle-1">ご協力ありがとうございました！</p>
        <v-row>
            <v-col cols="6">
                <v-img :src="src1"></v-img>
            </v-col>
            <v-col cols="6">
                <v-img :src="src2"></v-img>
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
        src1: {
            type: String,
        },
        src2: {
            type: String,
        },
    }
}
</script>