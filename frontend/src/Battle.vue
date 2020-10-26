<template>
    <v-main>
      <Question :question="question" />
      <v-row class="align-center">
        <v-col md="5" cols="12">
          <Language 
            :name="language_1"
            imgsrc="https://www.php.net//images/logos/new-php-logo.svg"
          />
        </v-col>
        <v-col class="display-3" cols="12">
          <div class="text-center">
            or
          </div>
        </v-col>
        <v-col md="5" cols="12">
          <Language
            :name="language_2"
            imgsrc="https://golang.org/lib/godoc/images/go-logo-blue.svg"
          />
        </v-col>
      </v-row>
    </v-main>
</template>

<script>
import Question from './components/Battle/Question';
import Language from './components/Battle/Language';

export default {
  name: 'App',

  components: {
    Question,
    Language,
  },

  data: () => ({
    question: '',
  }),

  computed: {
    language_1: function () {
      return this.$route.params["l1"];
    },
    language_2: function () {
      return this.$route.params["l2"];
    }
  },

  created: function () {
    this.getQuestion();
  },

  methods: {
    getQuestion: function () {
      this.question = '質問が入ります';
    },

  }
};
</script>
