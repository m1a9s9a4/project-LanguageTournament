<template>
  <v-main>
		<template v-if="player && opponent">
			<h3 class="text-center">{{ player.japanese }} vs {{opponent.japanese}} アンケート結果</h3>
			<p class="text-center">※バー中央の%が<strong>{{player.japanese}}</strong>を選んだ割合です。</p>
			<v-row>
				<v-col cols="6"><v-img :src="player.img"></v-img></v-col>
				<v-col cols="6"><v-img :src="opponent.img"></v-img></v-col>
			</v-row>
			<v-row>
				<v-col>
					<v-btn block color="primary" :href="questionUrl">回答する</v-btn>
				</v-col>
			</v-row>
			<v-row>
				<v-col cols="6">
					<Twitter :title="player.english+' vs '+opponent.english" :url="questionUrl" text="回答を求める" />
				</v-col>
				<v-col cols="6">
					<Facebook :title="player.english+' vs '+opponent.english" :url="questionUrl" />
				</v-col>
			</v-row>
			<template >
				<v-row>
					<template v-if="questions.length > 0">
						<v-col cols="12" v-for="(q, i) in questions" v-bind:key="i">
							<p>{{i+1}}. {{q.japanese}}</p>
							<template v-if="q.data">
								<v-progress-linear
									v-model="q.data.percentage"
									height="25"
									color="info"
								>
									<template v-slot="{ value }">
										<strong>{{ Math.ceil(value) }}% / 回答数{{q.data.sum}}</strong>
									</template>
								</v-progress-linear>
							</template>
							<template v-else>
								<v-alert dense border="left" type="warning">
										回答がありません
								</v-alert>
							</template>
						</v-col>
					</template>
					<template v-else>
						<v-col class="text-center">
							<Loading />
						</v-col>
					</template>
				</v-row>
			</template>
			<v-row>
				<v-col>
					<v-btn block color="primary" :href="questionUrl">回答する</v-btn>
				</v-col>
				<v-col>
					<v-btn block color="primary" outlined href="/">トップに戻る</v-btn>
				</v-col>
			</v-row>
		</template>
		<template v-else>
			<Loading />
		</template>
  </v-main>
</template>

<script>
import Axios from "axios";
import Twitter from './components/Share/Twitter';
import Facebook from './components/Share/Facebook';
import Loading from './components/Common/Loading';

export default {
  data: () => ({
    player: {},
		opponent: {},
		battleId: null,
		questions: [],
		answers: [],
		summaries: [],
	}),
	components: {
		Twitter,
		Facebook,
		Loading,
	},
	computed: {
		questionUrl: function() {
			return '/' + this.player.english + '/vs/' + this.opponent.english;
		}
	},
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
					if (res.data == null) return;

					let sum = 0;
					let p_count = 0;
					res.data.map(d => {
						if (d.target_id == this.player.id) {
							p_count = d.count;
						}
						sum += d.count;
					});
					const percentage = p_count / sum * 100;
					this.questions = this.questions.map(q => {
						if (q.id == qid) {
							q['data'] = {sum, percentage};
						}
						return q;
					})
					this.summaries.push({[qid]: {sum, percentage}});
				});
		},
    getBattleId: function () {
      Axios
        .get("/api/v1/battle/" +this.player.id+"/"+this.opponent.id)
        .then(res => {
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