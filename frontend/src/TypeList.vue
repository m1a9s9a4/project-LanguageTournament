<template>
    <v-main>
        <h2>{{ title }}一覧</h2>
        <v-row>
            <v-col md="3" cols="6" v-for="(p, i) in players" v-bind:key="i">
                <Language 
                    :player=p
                />
            </v-col>
        </v-row>
    </v-main>
</template>

<script>
import Language from './components/TypeList/Language';
import axios from 'axios';

export default {
    components: {
        Language,
    },
    data () {
        return {
            title: null,
            playerType: {},
            players: [{}],
        }
    },
    methods: {
        getType() {
            axios
            .get("/api/v1/player_type/" + this.$route.params["tid"])
            .then(res => {
                console.log(res.data);
                this.title = res.data.japanese;
                this.playerType = res.data;
            })
            .catch(e => {
                console.log(e);
            })
        },
        getPlayerByTypeId() {
            axios
            .get("/api/v1/players/type/" + this.$route.params["tid"])
            .then(res => {
                console.log(res);
                this.players = res.data;
            })
            .catch(e => {
                console.log(e);
            });
        }
    },
    mounted: function () {
        this.getType();
        this.getPlayerByTypeId();
    },
}
</script>