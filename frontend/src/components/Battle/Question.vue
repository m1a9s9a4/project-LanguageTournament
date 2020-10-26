<template>
    <p class="subtitle-1 text-center mt-5 pa-5" v-text="question"></p>
</template>

<script>
export default {
    props: {
        'question': {
            type: String,
            default: 'hogehgoe',
        }
    },
}
</script>