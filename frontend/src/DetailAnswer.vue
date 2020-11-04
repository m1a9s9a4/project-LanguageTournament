<template>
  <v-main>
    <h3 class="text-center">{{ player.japanese }} vs {{opponent.japanese}} アンケート結果</h3>
		<v-row>
			<v-col cols="12" v-for="(q, i) in questions" v-bind:key="i">
				<p>{{i+1}}. {{q.japanese}}</p>
			</v-col>
		</v-row>
  </v-main>
</template>

<script>
import Axios from "axios";

export default {
  data: () => ({
    player: {},
		opponent: {},
		battleId: null,
		questions: [],
		answers: [],
	}),
	created: function() {
		this._getPlayerByEnglish(this.$route.params["eng1"], 'player');
		this._getPlayerByEnglish(this.$route.params["eng2"], 'opponent')
	},
	mounted: function() {
	},
  methods: {
		_getPlayerByEnglish: function (eng, target) {
			Axios
				.get("/api/v1/player/english/"+eng)
				.then(res => {
					if (target == 'player') {
						this.player = res.data;
					} else if (target == 'opponent') {
						this.opponent = res.data;
					}
				})
				.catch(e => {
					console.error(e);
				})
				.finally(() => {
					if (this.player.id && this.opponent.id) {
						this.getBattleId();
					}
				})
		},
    _getAnswersByQid: function (qid) {
			Axios
				.get("/api/v1/answer/battle/"+this.battleId+"/"+qid+"/count")
				.then(res => {
					res.data.map(d => {
						// いい感じにq => [total]とかにした
						this.answers.push({[qid]: {d}});
					});
				})
				.finally(() => {

				});
		},
    getBattleId: function () {
      Axios
        .get("/api/v1/battle/" +this.player.id+"/"+this.opponent.id)
        .then(res => {
					console.log(res.data);
					this.battleId = res.data.id;
        })
        .catch(e => {
          console.error(e);
				})
				.finally(() => {
					if (this.questions.length <= 0) {
						this.getQuestions();
					}
				})
		},
		getQuestions: function() {
			Axios
				.get("/api/v1/questions/type/"+this.player.type_id)
				.then(res => {
					this.questions = res.data;
				})
				.finally(() => {
					this.questions.map(q => {
						this._getAnswersByQid(q.id);
					})
				})
		}
	},
};
</script>