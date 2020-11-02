<template>
    <div class="subtitle-1 text-center mt-5 pa-5">
        <p>どちらの選択肢が以下の質問を満たしていますか？</p>
        <p>{{current}} / {{total}}</p>
        <p>「{{question.japanese}}」</p>
    </div>
</template>

<script>
export default {
    props: {
        'current': {
            type: Number,
        },
        'total': {
            type: Number,
        },
        'question': {
            type: Object,
        }
    },
}
</script>