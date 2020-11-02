<template>
    <v-main>
        <Question v-if="currentQuestion" :question="currentQuestion" :current="currentNumber" :total="totalQuestions" />
        <v-row class="align-center">
          <v-col md="5" cols="12">
            <Language 
              :id="player1.id"
              :name="player1.english"
              :imgsrc="player1.img"
              @selected-id="execute"
            />
          </v-col>
          <v-col class="display-3" cols="12">
            <div class="text-center">
              or
            </div>
          </v-col>
          <v-col md="5" cols="12">
            <Language
              :id="player2.id"
              :name="player2.english"
              :imgsrc="player2.img"
              @selected-id="execute"
            />
          </v-col>
        </v-row>        
    </v-main>
</template>

<script>
import Question from './components/Battle/Question';
import Language from './components/Battle/Language';
import Axios from 'axios';

export default {
  name: 'App',

  components: {
    Question,
    Language,
  },

  data: () => ({
    player1: [],
    player2: [],
    questions: {},
    currentQuestion: null,
    currentNumber: 1,
    totalQuestions: 0,
    error: [],
  }),


  created: function () {
    this.getQuestion();
  },

  mounted: function () {
    this.getPlayer(this.$route.params["p1"]);
    this.getPlayer(this.$route.params["p2"]);
  },

  methods: {
    getQuestion: function () {
      Axios
        .get("/api/v1/questions/type/1")
        .then(res => {
          this.questions = res.data;
          this.totalQuestions = this.questions.length;
          if (this.totalQuestions > 0) {
            this.currentQuestion = res.data[this.currentNumber];
          }
        })
        .catch(e => {
          console.log(e);
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
      })
      .catch(e => {
        console.log(e);
        this.error.push(english +": not found");
        return null;
      });
    },
    execute: function (id) {
      this._toNextQuestion();
      console.log(id);
    },
    _setCurrentNumberFromIndex: function (index) {
      this.currentNumber = index++;
    },
    _toNextQuestion: function () {
      if (this.currentNumber < this.totalQuestions) {
        this.currentNumber++;
        this.currentQuestion = this.questions[this.currentNumber];
      }
    }
  }
};
</script>
