<template>
    <v-main>
        <h2>{{ title }}一覧</h2>
        <template v-if="players.length > 0">
            <v-row>
                <v-col md="3" cols="6" v-for="(p, i) in players" v-bind:key="i">
                    <Language 
                        :name=p.japanese
                        :english=p.english
                        :imgsrc=p.img
                    />
                </v-col>
            </v-row>
        </template>
        <template v-else>
            <Loading />
        </template>
    </v-main>
</template>

<script>
import Language from './components/Home/Language';
import Loading from './components/Common/Loading';
import axios from 'axios';

export default {
    components: {
        Language,
        Loading,
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
            .get("/api/v1/player_type/1")
            .then(res => {
                this.title = res.data.japanese;
                this.playerType = res.data;
            })
            .catch(e => {
                console.error(e);
            })
        },
        getPlayerByTypeId() {
            axios
            .get("/api/v1/players/type/1")
            .then(res => {
                this.players = res.data;
            })
            .catch(e => {
                console.error(e);
            });
        }
    },
    mounted: function () {
        this.getType();
        this.getPlayerByTypeId();
    },
}
</script>