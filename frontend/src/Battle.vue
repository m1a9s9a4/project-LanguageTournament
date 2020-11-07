<template>
    <v-main v-cloak>
      <template v-if="questionFinished">
        <Complete :src1="player1.img" :src2="player2.img" :english1="english1" :english2="english2" />
      </template>
      <template v-else>
        <Question v-if="currentQuestion" :question="currentQuestion" :current="currentNumber" :total="numberOfQuestions" />
        <v-row class="align-center">
          <template v-if="isAlreadyAnswered">
            <p>すでに回答済みの質問です。</p>
          </template>
          <template v-else>
            <v-col md="5" cols="12">
              <Language 
                :id="player1.id"
                :name="player1.english"
                :imgsrc="player1.img"
                @selected-id="saveAnswer"
              />
            </v-col>
            <v-col class="display-3" md="2" cols="12">
              <div class="text-center">
                or
              </div>
            </v-col>
            <v-col md="5" cols="12">
              <Language
                :id="player2.id"
                :name="player2.english"
                :imgsrc="player2.img"
                @selected-id="saveAnswer"
              />
            </v-col>
          </template>
        </v-row>
      </template>
    </v-main>
</template>

<script>
import Question from './components/Battle/Question';
import Language from './components/Battle/Language';
import Complete from './components/Battle/Complete';
import Axios from 'axios';

export default {
  name: 'App',

  components: {
    Question,
    Language,
    Complete,
  },

  data: () => ({
    english1: null,
    english2: null,
    player1: [],
    player2: [],
    questions: {},
    currentQuestion: null,
    currentNumber: 1,
    battleId: null,
    uid: null,
    isAlreadyAnswered: false,
    numberOfQuestions: 0,
  }),


  created: function () {
    this.english1 = this.$route.params["p1"];
    this.english2 = this.$route.params["p2"];
    this.uid = localStorage.getItem("uid");
    this.getPlayer(this.english1);
    this.getPlayer(this.english2);
  },

  computed: {
    currentIndex: function () {
      return this.currentNumber - 1;
    },
    questionFinished: function () {
      return this.numberOfQuestions <= this.currentNumber;
    },
  },

  methods: {
    getBattleId: function () {
      Axios
        .get("/api/v1/battle/" +this.player1.id+"/"+this.player2.id)
        .then(res => {
          this.battleId = res.data.id;
        })
        .catch(e => {
          console.error(e);
        })
        .finally(() => {
          this._checkAnsweredQuestions();
        });
    },
    getQuestion: function () {
      Axios
        .get("/api/v1/questions/type/"+this.player1.type_id)
        .then(res => {
          this.questions = res.data;
          this.numberOfQuestions = this.questions.length;
        })
        .catch(e => {
          console.error(e);
        })
        .finally(() => {
          
        })
    },
    getPlayer: function (english) {
      Axios
      .get("/api/v1/player/english/" + english)
      .then(res => {
        if (this.player1.length < 1) {
          this.player1 = res.data;
          this.getQuestion();
          return;
        }

        this.player2 = res.data;
        this.getBattleId();
      })
      .catch(e => {
        console.error(e);
      });
    },
    saveAnswer: function (id) {
      const params = {
          question_id: this.currentQuestion.id,
          selected_player_id: id,
          battle_id: this.battleId,
      };
      Axios
        .post("/api/v1/answer", params, {
          headers: {
            'Content-Type': 'application/json',
          }
        })
        .then(() => {
          this._saveUser(this.currentQuestion.id);
        })
        .catch(e => {
          console.error(e);
        })
        .finally(() => {
          this._toNextQuestion();
        });
    },
    _checkAnsweredQuestions: function() {
      Axios
        .get("/api/v1/answer/users", {
          params: {
            uid: this.uid,
            bid: this.battleId,
          }
        })
        .then(res => {
          const qids = res.data.map(data => {
            return data.question_id;
          });
          this.questions = this.questions.filter((q) => {
            return !qids.includes(q.id);
          });
        })
        .catch(e => {
          console.error(e);
        })
        .finally(() => {
          this._resetTotalQuestionNumber();
          this._resetCurrentQuestion();
        });
    },
    _resetCurrentQuestion: function () {
      if (this.numberOfQuestions > 0) {
        this.currentQuestion = this.questions[this.currentIndex];
      }
    },
    _resetTotalQuestionNumber: function () {
      this.numberOfQuestions = this.questions.length;
    },
    _saveUser: function (qid) {
      if (this.uid == null) {
        return;
      }
      const params = {
        question_id: qid,
        battle_id: this.battleId,
        uid: this.uid,
      };
      Axios
        .post("/api/v1/answer/user", params, {
          headers: {
            'Content-Type': 'application/json',
          }
        })
        .catch(e => {
          console.error(e);
        });
    },
    _setCurrentNumberFromIndex: function (index) {
      this.currentNumber = index++;
    },
    _toNextQuestion: function () {
      if (this.currentNumber <= this.numberOfQuestions) {
        this.currentNumber++;
        this._resetCurrentQuestion();
      }
    },
  }
};
</script>
