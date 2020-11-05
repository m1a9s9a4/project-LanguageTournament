<template>
    <div class="subtitle-1 text-center">
        <p>{{current}} / {{total}} どちらが該当しますか？</p>
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